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

//修改用户名
func UpdateNameById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户id错误"})
		return
	}
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户名不能为空"})
		return
	}
	reader := model.UpdateReaderName(id, name)
	if reader != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "更新用户名成功"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "更新用户名失败"})
	}
}

//注册用户
func RegisterReader(ctx *gin.Context) {

	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	sex := ctx.PostForm("sex")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户名不能为空"})
		return
	}
	if password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "密码不能为空"})
		return
	}
	if sex == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "性别不能为空"})
		return
	}
	reader := model.InsertReaderInfo(name, password, sex)
	if reader != nil {
		ctx.JSON(http.StatusCreated, gin.H{"message": "用户注册成功"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "注册用户失败"})
	}
}

//添加用户
func AddReader(ctx *gin.Context) {

	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	sex := ctx.PostForm("sex")
	reader := model.InsertReaderInfo(name, password, sex)
	if reader != nil {
		ctx.JSON(http.StatusCreated, gin.H{"message": "用户添加成功"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "添加用户失败"})
	}
}
