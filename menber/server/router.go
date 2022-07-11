package server

import (
    "github.com/gin-gonic/gin"
    "gin_test/controller/top"
    "gin_test/controller/user"
)

func GetRouter() *gin.Engine {    // *gin.Engineの表記は返り値の型
    router := gin.Default()
    router.LoadHTMLGlob("view/*.html")

    router.GET("/",top.IndexDisplayAction)
    router.GET("/user/", user.ListDisplayAction)
    router.GET("/user/add", user.AddDisplayAction)
    router.GET("/user/edit/:id", user.UserEditDisplayAction)

    router.POST("/user/edit/complete", user.EditCompleteAction)
    router.POST("/user/add/complete", user.AddCompleteAction)
    return router
}