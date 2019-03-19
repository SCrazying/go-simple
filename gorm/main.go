package main

import (
	"github.com/gin-gonic/gin"
	"go-simple/gorm/control"
)

func main() {

	router := gin.Default()
	router.POST("readerlogin", control.ReaderLogin)
	router.GET("/reader/:id/", control.ReaderInfo)
	router.GET("/reader", control.ReadersInfo)
	router.PUT("/reader/:id/password/:password", control.UpdatePasswordById)
	router.PUT("/reader/:id/name/:name", control.UpdateNameById)
	router.POST("/reader/", control.RegisterReader)
	router.Run(":2333")
}
