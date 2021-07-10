package model

import (
	"github.com/tal-tech/go-zero/core/stores/sqlc"
)

type UserModel struct {
	sqlc.CachedConn
	table string
}

type UserModelTable struct {
	Id int64	`db:"id"`
	Version string `db:"version"`
	Content string	`db:"content"`
	ContentDescription string `db:"content_description"`
	UpdateTime string	`db:"update_time"`
	CreateTime string `db:"createtime"`
}
