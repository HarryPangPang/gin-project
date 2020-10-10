package controller

import "github.com/gin-gonic/gin"

// import (
// 	model "gmt-go/model"
// 	entitys "gmt-go/model/entitys"

// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// Test 测试
func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": "hello",
	})
}

// // GetUser 获取所有数据
// func GetUser(c *gin.Context) {
// 	users := []entitys.User{}
// 	model.DB.Find(&users)
// 	if err := model.DB.Find(&users).Error; err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code": 1,
// 			"msg":  "发生错误",
// 			"data": nil,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 0,
// 		"msg":  "获取成功",
// 		"data": &users,
// 	})
// }

// //AddUser 新增数据
// func AddUser(c *gin.Context) {
// 	// 参数校验
// 	var user entitys.User
// 	err := c.ShouldBind(&user)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"code": 1,
// 			"msg":  err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	result := model.DB.Create(&user)

// 	if result.Error != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code": 1,
// 			"msg":  result.Error.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 0,
// 		"msg":  "创建成功",
// 		"data": user.ID,
// 	})
// }
