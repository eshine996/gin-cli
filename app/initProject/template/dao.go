package template

var DaoTmp = `package dao

import (
	"{{projectName}}/dao/internal"
)

var User = internal.User`

var DaoInternalBaseTmp = `package internal

import (
	g "{{projectName}}/global"
	"gorm.io/gorm"
)

type baseDao struct{}

func (b *baseDao) Mysql(dsn ...string) (db *gorm.DB) {
	if len(dsn) < 1 {
		db, _ = g.GetMysqlDB()
	} else {
		db, _ = g.GetMysqlDBByDSN(dsn[0])
	}
	return
}`

var DaoInternalUserTmp = `package internal

import (
	"context"
	"{{projectName}}/model/entity"
	"gorm.io/gorm"
)

var User = doUser{}

type doUser struct {
	baseDao
	entity.User
}

func (d *doUser) Ctx(ctx context.Context, dsn ...string) *gorm.DB {
	return d.Mysql(dsn...).Model(&d.User).WithContext(ctx)
}`
