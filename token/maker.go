package token

import "time"

type Maker interface {
	//创建token
	CreateToken(userId int64, username string, duration time.Duration) (string, *Payload, error)

	//校验token是否有效
	VerifyToken(token string) (*Payload, error)
}
