package template

import "fmt"

var ModelApiUserTmp = `package api

type LoginReq struct {
	Username string ` + fmt.Sprintf("`%s`", `json:"username" binding:"required"`) + `
	Password string ` + fmt.Sprintf("`%s`", `json:"password" binding:"required"`) + `
}`

var ModelEntityUserTmp = `package entity

type User struct {
	Id       int    ` + fmt.Sprintf("`%s`", `gorm:"column:id"`) + `
	Username string ` + fmt.Sprintf("`%s`", `gorm:"column:username"`) + `
	Password string ` + fmt.Sprintf("`%s`", `gorm:"column:password"`) + `
}

func (u *User) TableName() string {
	return "user"
}`

var ModelUserTmp = `package model

type UserOutput struct {
	Username string
	Password string
}
`
