package routes

import (
	"github.com/gorilla/mux"
	"go_2/pkg/controllers"
)

var RegisterVideoStore = func(router *mux.Router){
	router.HandleFunc("/video/", controllers.CreateVideo).Methods("POST")
	router.HandleFunc("/video/", controllers.GetAllVideos).Methods("GET")
	router.HandleFunc("/video/{videoId}", controllers.GetVideoById).Methods("GET")
	router.HandleFunc("/video/{videoId}", controllers.DeleteVideo).Methods("DELETE")
	router.HandleFunc("/video/{videoId}", controllers.UpdateVideo).Methods("PUT")
}