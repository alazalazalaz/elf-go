package handler

import (
	"elf-go/app"
	"elf-go/components/logs"
	"elf-go/example/app/dao/entity"
	"github.com/gin-gonic/gin"
)

//
//type UserAndArticle struct{
//	Id int `json:"id"`
//	Username string `json:"username"`
//	CreatedAt int `json:"created_at"`
//	ArticleInfo entity.Article `gorm:"foreignkey:Id"`
//	//ArticleInfo entity.Article
//}

type Article struct{
	Id int `json:"id" gorm:"primary_key"`
	Uid int `json:"uid"`
	Title string `json:"title"`
	CreatedAt int `json:"created_at"`
	UserInfo entity.User `gorm:"foreignkey:articles_ibfk_1"`
}

func GetUserInfo(ctx *gin.Context){
	logs.Info("get user info")
	article := Article{}
	app.Mysql().Find(&article)
	//app.Mysql().Table("article as a").Joins("left join user as u on u.id=a.uid").
	//	Where("a.id>?", 0).Find(&article)
	logs.Info("result: ", logs.Content{"article:": article})
}

func Version(ctx *gin.Context){
	logs.Info("version")
}

func Update(ctx *gin.Context){
	logs.Warning("update")
}