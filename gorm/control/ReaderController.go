package control

import (
	"github.com/gin-gonic/gin"
	"go-simple/gorm/model"
	"net/http"
	"strconv"
)

type ReaderController struct {
}

//读者登录
func ReaderLogin(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户id错误"})

		return
	}
	password := ctx.PostForm("password")
	if password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "密码不能为空"})
		return
	}
	user := model.FindReader(id, password)
	if user != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "用户登录成功"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户名或者密码错误"})
	}

}

//获取用户信息
func ReaderInfo(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户id错误"})
		return
	}
	user := model.FindReaderInfo(id)
	if user != nil {
		//json, err := json.Marshal(user)

		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户不存在"})
	}
}

//获取所有读者信息
func ReadersInfo(ctx *gin.Context) {
	users := model.FindReaders()
	if users != nil {
		ctx.JSON(http.StatusOK, users)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "获取读者信息列表失败"})
	}
}

//修改密码
func UpdatePasswordById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户id错误"})
		return
	}
	password := ctx.Param("password")
	if password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "密码不能为空"})
		return
	}
	reader := model.UpdateReaderPassword(id, password)
	if reader != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "更新用户成功"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "更新用户失败"})
	}
}
