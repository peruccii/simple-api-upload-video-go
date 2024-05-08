package controllers

import (
	"go_2/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"go_2/pkg/utils"
)

var ControlsVideo models.Video

func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	videoDetails := models.GetAllVideo()
	res, _ := json.Marshal(videoDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetVideoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // <-- return routes
	videoId := vars["videoId"] // <-- search for params that in my case is {videoId}
	ID, err := strconv.ParseInt(videoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	videoDetails, _:= models.GetVideoById(ID)
	res, _ := json.Marshal(videoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateVideo(w http.ResponseWriter, r *http.Request){
	CreateVideo := &models.Video{}
	utils.ParseBody(r, CreateVideo)
	v:= CreateVideo.CreateVideo()
	res, _ := json.Marshal(v)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	videoId := vars["videoId"]
	ID, err := strconv.ParseInt(videoId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	videoDetails := models.DeleteVideo(ID)
	res, _ := json.Marshal(videoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateVideo(w http.ResponseWriter, r *http.Request){
	var updateVideo = &models.Video{}
	utils.ParseBody(r, updateVideo)
	vars := mux.Vars(r)
	videoId := vars["videoId"]
	ID, err := strconv.ParseInt(videoId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	videoDetails, db:=models.GetVideoById(ID)
	if updateVideo.Name != ""{
		videoDetails.Name = updateVideo.Name
	}
	if updateVideo.Url != ""{
		videoDetails.Url = updateVideo.Url
	}
	if updateVideo.Uploated_at != ""{
		videoDetails.Uploated_at = updateVideo.Uploated_at
	}
	db.Save(&videoDetails)
	res, _ := json.Marshal(videoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}