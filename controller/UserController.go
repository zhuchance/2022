package controller

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/common"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/response"
	"github.com/EDDYCJY/go-gin-example/tdo"
	"github.com/EDDYCJY/go-gin-example/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.InitDB()
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		fmt.Println(telephone)
		return
	}
	if len(password) < 6 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "密码不能少于6位")
		return
	}
	// 如果名称没有传，就给一个10位的字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExist(common.InitDB(), telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	//创建用户
	//加密用户密码
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, gin.H{"code": 422, "msg": "加密错误"})
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "密码错误")
		return
	}
	newUser := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasePassword),
	}
	//返回结果
	//ctx.JSON(200, gin.H{
	//	"message": "注册成功",
	//})
	response.Success(ctx, nil, "注册成功")
	DB.Create(&newUser)
	//fmt.Println(&newUser)
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

func Loing(ctx *gin.Context) {

	DB := common.InitDB()
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//判断手机号是否存在
	var user models.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	//ctx.JSON(200, gin.H{
	//	"code": 200,
	//	"data": gin.H{"token": token},
	//	"msg":  "登录成功",
	//})
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": tdo.ToUserDto(user.(models.User))}})
}
