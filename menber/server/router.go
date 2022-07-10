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
    router.GET("/user", user.UserListDisplayAction)
    router.GET("/user/add", user.UserAddDisplayAction)
    router.GET("/user/edit/:id", user.UserEditDisplayAction)

    return router
}