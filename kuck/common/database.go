package common

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func InitDB()  *gorm.DB{
	//dsn:=viper.GetString("datasource.username")+viper.GetString("dataspurce.password")+viper.GetString("datasource.")
	//host := viper.GetString("datasource.host")
	//port:=viper.GetString("datasource.port")
	//
	dsn := "root:root123@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//fmt.Println("error: ")
		panic("fail to connection database, err: " + err.Error())
	}
	db.AutoMigrate(&models.User{})
	return db



}
