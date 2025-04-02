// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateCopy(ctx context.Context, arg CreateCopyParams) (Copywriting, error)
	CreateMultipleCopy(ctx context.Context, arg CreateMultipleCopyParams) error
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCopy(ctx context.Context, arg DeleteCopyParams) error
	GetCopy(ctx context.Context, id int64) (Copywriting, error)
	GetOauth2Token(ctx context.Context, arg GetOauth2TokenParams) (GetOauth2TokenRow, error)
	GetPlatforms(ctx context.Context) ([]GetPlatformsRow, error)
	GetPlatformsByName(ctx context.Context, platform string) (GetPlatformsByNameRow, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetThirdKeyByName(ctx context.Context, name string) (ThirdKey, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserByName(ctx context.Context, username string) (User, error)
	InsertOauth2Token(ctx context.Context, arg InsertOauth2TokenParams) (Oauth2Token, error)
	InsertThirdKey(ctx context.Context, arg InsertThirdKeyParams) (ThirdKey, error)
	InsertVideo(ctx context.Context, arg InsertVideoParams) (Video, error)
	ListCopies(ctx context.Context, arg ListCopiesParams) ([]Copywriting, error)
	UpdateAccessToken(ctx context.Context, arg UpdateAccessTokenParams) (Oauth2Token, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
