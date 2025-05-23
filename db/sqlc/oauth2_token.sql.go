// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: oauth2_token.sql

package db

import (
	"context"
	"time"
)

const getOauth2Token = `-- name: GetOauth2Token :one
select id, user_id, provider,api,access_token,token_type,refresh_token,expiry
from oauth2_tokens
where user_id = $1
and provider = $2
and api = $3
and delete_at is null
order by id desc
`

type GetOauth2TokenParams struct {
	UserID   int64  `json:"user_id"`
	Provider string `json:"provider"`
	Api      string `json:"api"`
}

type GetOauth2TokenRow struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	Provider     string    `json:"provider"`
	Api          string    `json:"api"`
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
}

func (q *Queries) GetOauth2Token(ctx context.Context, arg GetOauth2TokenParams) (GetOauth2TokenRow, error) {
	row := q.db.QueryRow(ctx, getOauth2Token, arg.UserID, arg.Provider, arg.Api)
	var i GetOauth2TokenRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Provider,
		&i.Api,
		&i.AccessToken,
		&i.TokenType,
		&i.RefreshToken,
		&i.Expiry,
	)
	return i, err
}

const insertOauth2Token = `-- name: InsertOauth2Token :one
insert into oauth2_tokens(
                          user_id,
                          provider,
                          api,
                          access_token,
                          token_type,
                          refresh_token,
                          expiry
) values (
          $1,$2,$3,$4,$5,$6,$7
         ) RETURNING id, user_id, provider, api, access_token, token_type, refresh_token, expiry, created_at, delete_at
`

type InsertOauth2TokenParams struct {
	UserID       int64     `json:"user_id"`
	Provider     string    `json:"provider"`
	Api          string    `json:"api"`
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
}

func (q *Queries) InsertOauth2Token(ctx context.Context, arg InsertOauth2TokenParams) (Oauth2Token, error) {
	row := q.db.QueryRow(ctx, insertOauth2Token,
		arg.UserID,
		arg.Provider,
		arg.Api,
		arg.AccessToken,
		arg.TokenType,
		arg.RefreshToken,
		arg.Expiry,
	)
	var i Oauth2Token
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Provider,
		&i.Api,
		&i.AccessToken,
		&i.TokenType,
		&i.RefreshToken,
		&i.Expiry,
		&i.CreatedAt,
		&i.DeleteAt,
	)
	return i, err
}

const updateAccessToken = `-- name: UpdateAccessToken :one
update oauth2_tokens
set access_token = $1, token_type = $2, expiry = $3
where user_id = $4 and provider = $5 and refresh_token = $6 and token_type = $7
RETURNING id, user_id, provider, api, access_token, token_type, refresh_token, expiry, created_at, delete_at
`

type UpdateAccessTokenParams struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	Expiry       time.Time `json:"expiry"`
	UserID       int64     `json:"user_id"`
	Provider     string    `json:"provider"`
	RefreshToken string    `json:"refresh_token"`
	TokenType_2  string    `json:"token_type_2"`
}

func (q *Queries) UpdateAccessToken(ctx context.Context, arg UpdateAccessTokenParams) (Oauth2Token, error) {
	row := q.db.QueryRow(ctx, updateAccessToken,
		arg.AccessToken,
		arg.TokenType,
		arg.Expiry,
		arg.UserID,
		arg.Provider,
		arg.RefreshToken,
		arg.TokenType_2,
	)
	var i Oauth2Token
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Provider,
		&i.Api,
		&i.AccessToken,
		&i.TokenType,
		&i.RefreshToken,
		&i.Expiry,
		&i.CreatedAt,
		&i.DeleteAt,
	)
	return i, err
}
