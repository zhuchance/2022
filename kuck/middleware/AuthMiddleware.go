package middleware

import (
	"github.com/EDDYCJY/go-gin-example/common"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": "权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//通过验证后获取claim中的userID
		userID := claims.UserID
		DB := common.InitDB()
		var user models.User
		DB.First(&user,userID)

		// 用户
		if user.ID == 0{
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":401,"msg": "权限不足"})
			ctx.Abort()
			return
		}
		//用户存在 将user 的信息写入上下文
		ctx.Set("user",user)
		ctx.Next()
	}
}
