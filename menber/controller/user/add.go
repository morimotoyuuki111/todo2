package user

import (
	"gin_test/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddDisplayAction(c *gin.Context){
    c.HTML(200, "user-add.html", gin.H{})
}

// 登録処理
func AddCompleteAction(c *gin.Context){
    name := c.PostForm("name")

    // モデルでテーブルへの登録
    model.AddUserData(name)

    // 会員一覧にリダイレクト
    c.Redirect(http.StatusSeeOther, "/user")
}