package models

import (
	"github.com/jinzhu/gorm"
	"go_2/pkg/config"
)

var db * gorm.DB

type Video struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Url string `json:"url"`
	Uploated_at string `json:"uploated_at"`
}

func init(){
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Video{})
}

func (v *Video) CreateVideo() *Video {
	db.NewRecord(v)
	db.Create(v)
	return v
}

func GetAllVideo() []Video {
	var Videos []Video 
	db.Find(&Videos)
	return Videos
}

func GetVideoById(Id int64) (*Video, *gorm.DB) {
	var getVideo Video
	db := db.Where("ID=?", Id).Find(&getVideo)
	return &getVideo, db
}

func DeleteVideo(ID int64) Video{
	var video Video
	db.Where("ID=?", ID).Delete(video)
	return video
}