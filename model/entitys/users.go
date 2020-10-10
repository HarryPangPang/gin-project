package entitys

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Gender   string `form:"gender" json:"gender"`
}
