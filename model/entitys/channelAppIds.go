package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	GameName      string `form:"gameName" json:"gameName"`
	alias         string `form:"alias" json:"alias"`
	gmId          string `form:"email" json:"email" binding:"required,email"`
	bbxRegion     string `form:"gender" json:"gender"`
	url           string `form:"gameName" json:"gameName" binding:"required"`
	imUrl         string `form:"alias" json:"alias" binding:"required"`
	accessType    int    `form:"email" json:"email" binding:"required,email"`
	apiSecret     string `form:"gender" json:"gender"`
	isAutoGetData int    `form:"gender" json:"gender"`
	sdk           string `form:"gender" json:"gender"`
	unisdk        string `form:"gender" json:"gender"`
}
