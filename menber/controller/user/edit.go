package user

import (
	"gin_test/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func UserEditDisplayAction(c *gin.Context) {
	id := c.Param("id") //URLからIDを取得して表示を分ける
	c.HTML(200,"user-edit.htnl",gin.H{
		"user": model.GetUserData(id),
	})
}

func EditCompleteAction(c *gin.Context){
	id := c.PostForm("id")
	name := c.PostForm("name")

	// モデルでテーブルへの変更
	model.EditUserData(id, name)

	// 会員一覧にリダイレクト
	c.Redirect(http.StatusSeeOther, "/user")
}