package router

import (
	"communication-app/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/msg", controller.GetAllMessage).Methods("GET")
	router.HandleFunc("/all_msg_file", controller.ServeAllMsgFile).Methods("GET")
	//router.HandleFunc("/sendVideo", controller.ServeVideoFile).Methods("GET")
	router.HandleFunc("/msg", controller.DeleteAllMessage).Methods("DELETE")
	router.HandleFunc("/msg/{id}", controller.GetMessageById).Methods("GET")
	router.HandleFunc("/msg/{id}", controller.DeleteOneMessage).Methods("DELETE")
	router.HandleFunc("/msg/{id}", controller.UpdateOneMessage).Methods("PUT")
	router.HandleFunc("/msg", controller.RecieveMessage).Methods("POST")

	return router
}
