// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: video.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Video struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Url           string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Duration      int32                  `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Video) Reset() {
	*x = Video{}
	mi := &file_video_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Video) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Video) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type GenerateVideoRequest struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	VideoSubject        string                 `protobuf:"bytes,1,opt,name=video_subject,json=videoSubject,proto3" json:"video_subject,omitempty"`
	VideoScript         string                 `protobuf:"bytes,2,opt,name=video_script,json=videoScript,proto3" json:"video_script,omitempty"`
	VideoTerms          string                 `protobuf:"bytes,3,opt,name=video_terms,json=videoTerms,proto3" json:"video_terms,omitempty"`
	VideoAspect         string                 `protobuf:"bytes,4,opt,name=video_aspect,json=videoAspect,proto3" json:"video_aspect,omitempty"`
	VideoConCatMode     string                 `protobuf:"bytes,5,opt,name=video_con_cat_mode,json=videoConCatMode,proto3" json:"video_con_cat_mode,omitempty"`
	VideoTransitionMode string                 `protobuf:"bytes,6,opt,name=video_transition_mode,json=videoTransitionMode,proto3" json:"video_transition_mode,omitempty"`
	VideoClipDuration   int32                  `protobuf:"varint,7,opt,name=video_clip_duration,json=videoClipDuration,proto3" json:"video_clip_duration,omitempty"`
	VideoCount          int32                  `protobuf:"varint,8,opt,name=video_count,json=videoCount,proto3" json:"video_count,omitempty"`
	VideoSource         string                 `protobuf:"bytes,9,opt,name=video_source,json=videoSource,proto3" json:"video_source,omitempty"`
	VideoMaterals       string                 `protobuf:"bytes,10,opt,name=video_materals,json=videoMaterals,proto3" json:"video_materals,omitempty"` // 注意: 原拼写可能是materials
	VideoLanguage       string                 `protobuf:"bytes,11,opt,name=video_language,json=videoLanguage,proto3" json:"video_language,omitempty"`
	VoiceName           string                 `protobuf:"bytes,12,opt,name=voice_name,json=voiceName,proto3" json:"voice_name,omitempty"`
	VoiceVolume         float32                `protobuf:"fixed32,13,opt,name=voice_volume,json=voiceVolume,proto3" json:"voice_volume,omitempty"`
	VoiceRate           float32                `protobuf:"fixed32,14,opt,name=voice_rate,json=voiceRate,proto3" json:"voice_rate,omitempty"`
	BgmType             string                 `protobuf:"bytes,15,opt,name=bgm_type,json=bgmType,proto3" json:"bgm_type,omitempty"`
	BgmFile             string                 `protobuf:"bytes,16,opt,name=bgm_file,json=bgmFile,proto3" json:"bgm_file,omitempty"`
	BgmVolume           float32                `protobuf:"fixed32,17,opt,name=bgm_volume,json=bgmVolume,proto3" json:"bgm_volume,omitempty"`
	SubtitleEnabled     bool                   `protobuf:"varint,18,opt,name=subtitle_enabled,json=subtitleEnabled,proto3" json:"subtitle_enabled,omitempty"`
	SubtitlePosition    string                 `protobuf:"bytes,19,opt,name=subtitle_position,json=subtitlePosition,proto3" json:"subtitle_position,omitempty"`
	CustomPosition      float32                `protobuf:"fixed32,20,opt,name=custom_position,json=customPosition,proto3" json:"custom_position,omitempty"`
	FontName            string                 `protobuf:"bytes,21,opt,name=font_name,json=fontName,proto3" json:"font_name,omitempty"`
	TextForeColor       string                 `protobuf:"bytes,22,opt,name=text_fore_color,json=textForeColor,proto3" json:"text_fore_color,omitempty"`
	TextBackgroundColor bool                   `protobuf:"varint,23,opt,name=text_background_color,json=textBackgroundColor,proto3" json:"text_background_color,omitempty"`
	FontSize            int32                  `protobuf:"varint,24,opt,name=font_size,json=fontSize,proto3" json:"font_size,omitempty"`
	StrokeColor         string                 `protobuf:"bytes,25,opt,name=stroke_color,json=strokeColor,proto3" json:"stroke_color,omitempty"`
	StrokeWidth         float32                `protobuf:"fixed32,26,opt,name=stroke_width,json=strokeWidth,proto3" json:"stroke_width,omitempty"` // 修正: 原结构体中有重复的stroke_color
	NThreads            int32                  `protobuf:"varint,27,opt,name=n_threads,json=nThreads,proto3" json:"n_threads,omitempty"`
	ParagraphNumber     int32                  `protobuf:"varint,28,opt,name=paragraph_number,json=paragraphNumber,proto3" json:"paragraph_number,omitempty"`
	FileName            string                 `protobuf:"bytes,29,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *GenerateVideoRequest) Reset() {
	*x = GenerateVideoRequest{}
	mi := &file_video_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateVideoRequest) ProtoMessage() {}

func (x *GenerateVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateVideoRequest.ProtoReflect.Descriptor instead.
func (*GenerateVideoRequest) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateVideoRequest) GetVideoSubject() string {
	if x != nil {
		return x.VideoSubject
	}
	return ""
}

func (x *GenerateVideoRequest) GetVideoScript() string {
	if x != nil {
		return x.VideoScript
	}
	return ""
}

func (x *GenerateVideoRequest) GetVideoTerms() string {
	if x != nil {
		return x.VideoTerms
	}
	return ""
}

func (x *GenerateVideoRequest) GetVideoAspect() string {
	if x != nil {
		return x.VideoAspect
	}
	return ""
}

func (x *GenerateVideoRequest) GetVideoConCatMode() string {
	if x != nil {
		return x.VideoConCatMode
	}
	return ""
}

func (x *GenerateVideoRequest) GetVideoTransitionMode() string {
	if x != nil {
		return x.VideoTransitionMode
	}
	return ""
}

func (x *GenerateVideoRequest) GetVideoClipDuration() int32 {
	if x != nil {
		return x.VideoClipDuration
	}
	return 0
}

func (x *GenerateVideoRequest) GetVideoCount() int32 {
	if x != nil {
		return x.VideoCount
	}
	return 0
}

func (x *GenerateVideoRequest) GetVideoSource() string {
	if x != nil {
		return x.VideoSource
	}
	return ""
}

func (x *GenerateVideoRequest) GetVideoMaterals() string {
	if x != nil {
		return x.VideoMaterals
	}
	return ""
}

func (x *GenerateVideoRequest) GetVideoLanguage() string {
	if x != nil {
		return x.VideoLanguage
	}
	return ""
}

func (x *GenerateVideoRequest) GetVoiceName() string {
	if x != nil {
		return x.VoiceName
	}
	return ""
}

func (x *GenerateVideoRequest) GetVoiceVolume() float32 {
	if x != nil {
		return x.VoiceVolume
	}
	return 0
}

func (x *GenerateVideoRequest) GetVoiceRate() float32 {
	if x != nil {
		return x.VoiceRate
	}
	return 0
}

func (x *GenerateVideoRequest) GetBgmType() string {
	if x != nil {
		return x.BgmType
	}
	return ""
}

func (x *GenerateVideoRequest) GetBgmFile() string {
	if x != nil {
		return x.BgmFile
	}
	return ""
}

func (x *GenerateVideoRequest) GetBgmVolume() float32 {
	if x != nil {
		return x.BgmVolume
	}
	return 0
}

func (x *GenerateVideoRequest) GetSubtitleEnabled() bool {
	if x != nil {
		return x.SubtitleEnabled
	}
	return false
}

func (x *GenerateVideoRequest) GetSubtitlePosition() string {
	if x != nil {
		return x.SubtitlePosition
	}
	return ""
}

func (x *GenerateVideoRequest) GetCustomPosition() float32 {
	if x != nil {
		return x.CustomPosition
	}
	return 0
}

func (x *GenerateVideoRequest) GetFontName() string {
	if x != nil {
		return x.FontName
	}
	return ""
}

func (x *GenerateVideoRequest) GetTextForeColor() string {
	if x != nil {
		return x.TextForeColor
	}
	return ""
}

func (x *GenerateVideoRequest) GetTextBackgroundColor() bool {
	if x != nil {
		return x.TextBackgroundColor
	}
	return false
}

func (x *GenerateVideoRequest) GetFontSize() int32 {
	if x != nil {
		return x.FontSize
	}
	return 0
}

func (x *GenerateVideoRequest) GetStrokeColor() string {
	if x != nil {
		return x.StrokeColor
	}
	return ""
}

func (x *GenerateVideoRequest) GetStrokeWidth() float32 {
	if x != nil {
		return x.StrokeWidth
	}
	return 0
}

func (x *GenerateVideoRequest) GetNThreads() int32 {
	if x != nil {
		return x.NThreads
	}
	return 0
}

func (x *GenerateVideoRequest) GetParagraphNumber() int32 {
	if x != nil {
		return x.ParagraphNumber
	}
	return 0
}

func (x *GenerateVideoRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type GenerateVideoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateVideoResponse) Reset() {
	*x = GenerateVideoResponse{}
	mi := &file_video_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateVideoResponse) ProtoMessage() {}

func (x *GenerateVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateVideoResponse.ProtoReflect.Descriptor instead.
func (*GenerateVideoResponse) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateVideoResponse) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type GetVideosRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Num           int32                  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVideosRequest) Reset() {
	*x = GetVideosRequest{}
	mi := &file_video_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVideosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideosRequest) ProtoMessage() {}

func (x *GetVideosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideosRequest.ProtoReflect.Descriptor instead.
func (*GetVideosRequest) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{3}
}

func (x *GetVideosRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetVideosRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type GetVideosResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Videos        *Video                 `protobuf:"bytes,1,opt,name=videos,proto3" json:"videos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVideosResponse) Reset() {
	*x = GetVideosResponse{}
	mi := &file_video_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVideosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideosResponse) ProtoMessage() {}

func (x *GetVideosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideosResponse.ProtoReflect.Descriptor instead.
func (*GetVideosResponse) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{4}
}

func (x *GetVideosResponse) GetVideos() *Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

var File_video_proto protoreflect.FileDescriptor

var file_video_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x5b, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbd,
	0x08, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x65, 0x72, 0x6d, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a, 0x12, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6f, 0x6e,
	0x5f, 0x63, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x32, 0x0a, 0x15, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6c,
	0x69, 0x70, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x11, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6c, 0x69, 0x70, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x67, 0x6d, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x67, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x67, 0x6d, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x67, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x67, 0x6d, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x62, 0x67, 0x6d, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10,
	0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x6f, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x6f, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x78, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x74, 0x65, 0x78, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6e, 0x74, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x6f, 0x6e, 0x74, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x5f, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x6f, 0x6b,
	0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65,
	0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x6f, 0x6b, 0x65, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x5f, 0x74,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x30,
	0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x36, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x73, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x75, 0x6c, 0x65, 0x31, 0x32, 0x33, 0x34, 0x2f, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46,
	0x6f, 0x72, 0x67, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_video_proto_rawDescOnce sync.Once
	file_video_proto_rawDescData []byte
)

func file_video_proto_rawDescGZIP() []byte {
	file_video_proto_rawDescOnce.Do(func() {
		file_video_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_video_proto_rawDesc), len(file_video_proto_rawDesc)))
	})
	return file_video_proto_rawDescData
}

var file_video_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_video_proto_goTypes = []any{
	(*Video)(nil),                 // 0: pb.Video
	(*GenerateVideoRequest)(nil),  // 1: pb.GenerateVideoRequest
	(*GenerateVideoResponse)(nil), // 2: pb.GenerateVideoResponse
	(*GetVideosRequest)(nil),      // 3: pb.GetVideosRequest
	(*GetVideosResponse)(nil),     // 4: pb.GetVideosResponse
}
var file_video_proto_depIdxs = []int32{
	0, // 0: pb.GetVideosResponse.videos:type_name -> pb.Video
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_video_proto_init() }
func file_video_proto_init() {
	if File_video_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_video_proto_rawDesc), len(file_video_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_video_proto_goTypes,
		DependencyIndexes: file_video_proto_depIdxs,
		MessageInfos:      file_video_proto_msgTypes,
	}.Build()
	File_video_proto = out.File
	file_video_proto_goTypes = nil
	file_video_proto_depIdxs = nil
}
