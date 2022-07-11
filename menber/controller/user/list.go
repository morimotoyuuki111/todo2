package user

import (
	"github.com/gin-gonic/gin"
	"gin_test/model"
)

func ListDisplayAction(c *gin.Context){
	c.HTML(200, "user-list.html", gin.H{
			// モデルでテーブルから会員全てを取得
			"list": model.GetUserList(),
	})
}