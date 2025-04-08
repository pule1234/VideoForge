package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pule1234/VideoForge/config"
	db "github.com/pule1234/VideoForge/db/sqlc"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/downloader"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/objects"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"log"
	"time"
)

type QiNiu struct {
	UploadManager   *uploader.UploadManager
	ObjectsManager  *objects.ObjectsManager
	DownloadManager *downloader.DownloadManager
}

func NewQiNiu(accessKey, secretKey string) *QiNiu {
	mac := credentials.NewCredentials(accessKey, secretKey)
	upload := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})

	object := objects.NewObjectsManager(&objects.ObjectsManagerOptions{
		Options: http_client.Options{Credentials: mac},
	})

	download := downloader.NewDownloadManager(&downloader.DownloadManagerOptions{
		Options: http_client.Options{Credentials: mac},
	})
	return &QiNiu{
		UploadManager:   upload,
		ObjectsManager:  object,
		DownloadManager: download,
	}
}

// 上传单个文件
func (q *QiNiu) UploadFile(ctx context.Context, localFile, bucketName, fileName, objectsName string) error {
	err := q.UploadManager.UploadFile(ctx, localFile, &uploader.ObjectOptions{
		BucketName: bucketName,
		ObjectName: &objectsName,
		FileName:   fileName,
		CustomVars: map[string]string{
			"name": fileName,
		},
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

// 获取用户的所有视频
func (q *QiNiu) GetAllFileByUser(ctx context.Context, bucketName string, user_id int64, store db.Store) ([]string, error) {
	bucket := q.ObjectsManager.Bucket(bucketName)

	// 回调函数
	onResponse := func(od *objects.ObjectDetails) {
		marshal, err := json.Marshal(od)
		if err != nil {
			return
		}
		fmt.Println(string(marshal))
		fmt.Printf("%s: %d bytes\n", od.Name, od.Size)
	}

	onError := func(err error) {
		fmt.Printf("Error: %v\n", err)
	}

	//获取文件列表
	fileNames := make([]string, 0)
	videos, _ := store.GetVideosByUid(ctx, user_id)
	for _, video := range videos {
		fileNames = append(fileNames, video.Title)
	}
	ops := make([]objects.Operation, len(fileNames))
	for i, name := range fileNames {
		ops[i] = bucket.Object(name).Stat().
			OnResponse(onResponse).
			OnError(onError)
	}
	if err := q.ObjectsManager.Batch(context.Background(), ops, nil); err != nil {
		return nil, err
	}
	return nil, nil
}

// 下载文件
func (q *QiNiu) DownloadFile(
	conf config.Config,
	fileName, userName, bucket string,
	subscribe int64,
	domains ...string,
) (string, error) {
	ds := make([]string, 0)
	for _, domain := range domains {
		ds = append(ds, domain)
	}
	timestamp := time.Now().Unix()
	urlsProvider := downloader.NewStaticDomainBasedURLsProvider(ds)
	//localFile := "/Users/a0000/111.mp4"
	localFile := conf.TempDir + fileName + "/" + userName + "_" + fmt.Sprintf("%d", timestamp) + ".mp4"
	//bucket := "videofore-videos"
	//objectName := "为什么要运动.mp4"
	objectName := fileName + userName + fmt.Sprintf("%d", subscribe) + ".mp4"
	downloadManager := q.DownloadManager
	downloaded, err := downloadManager.DownloadToFile(context.Background(), objectName, localFile, &downloader.ObjectOptions{
		GenerateOptions: downloader.GenerateOptions{
			BucketName:          bucket,
			UseInsecureProtocol: true,
		},
		DownloadURLsProvider: urlsProvider,
	})
	if err != nil && downloaded == 0 {
		log.Fatal(err)
		return "", err
	}
	return localFile, nil
}
