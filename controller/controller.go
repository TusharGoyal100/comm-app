package controller

import (
	"communication-app/connector"
	"communication-app/models"
	"communication-app/service"
	"communication-app/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func SentMesaage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside SentMessage")
	w.Header().Set("Content-Type", "application/json")

	msg := models.Message{
		MessageBody: "Hi everyone",
		PhoneNumber: "8607009442",
		TimeStamp:   time.Now().String(),
	}
	json.NewEncoder(w).Encode(msg)
}

func RecieveMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside ReciveMessage")
	w.Header().Set("Content-Type", "application/json")

	var msg models.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		json.NewEncoder(w).Encode("Body format is wrong")
		return
	}

	if msg.MessageBody == "" {
		json.NewEncoder(w).Encode("Message is empty")
		return
	}

	if msg.PhoneNumber == "" {
		json.NewEncoder(w).Encode("Message is empty")
		return
	}

	msg.TimeStamp = time.Now().Local().String()
	fmt.Println("Message Body is:-", msg)

	if err := connector.InsertOne(msg); err != nil {
		json.NewEncoder(w).Encode("Internal server error")
		return
	}

	json.NewEncoder(w).Encode("Succesfully Added Msg")
}

func GetMessageById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside GetMessageById")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var msgID string

	if params["id"] == "" {
		json.NewEncoder(w).Encode("please give an id")
		return
	} else {
		msgID = params["id"]
	}

	msg, err := connector.GetOne(msgID)
	if err != nil {
		json.NewEncoder(w).Encode("Internal server error")
		return
	}

	json.NewEncoder(w).Encode(msg)

}

func GetAllMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside GetAllMessage")
	w.Header().Set("Content-Type", "application/json")

	messages, err := connector.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode("Internal server error")
		return
	}

	json.NewEncoder(w).Encode(messages)

}

func DeleteOneMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside DeleteOneMessage")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var msgID string

	if params["id"] == "" {
		json.NewEncoder(w).Encode("please give an id")
		return
	} else {
		msgID = params["id"]
	}

	if err := connector.DeleteOne(msgID); err != nil {
		json.NewEncoder(w).Encode("Internal server error")
		return
	}

	json.NewEncoder(w).Encode("Succesfully deleted message")

}

func DeleteAllMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside DeleteAllMessage")
	w.Header().Set("Content-Type", "application/json")

	if err := connector.DeleteAll(); err != nil {
		json.NewEncoder(w).Encode("Internal server error")
		return
	}

	json.NewEncoder(w).Encode("Succesfully deleted all message")

}

func UpdateOneMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside UpdateOneMessage")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var msgID string

	var msg models.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		json.NewEncoder(w).Encode("Body format is wrong")
		return
	}

	if params["id"] == "" {
		json.NewEncoder(w).Encode("please give an id")
		return
	} else {
		msgID = params["id"]
	}

	if err := connector.UpdateOne(msgID, msg); err != nil {
		fmt.Println("Error;-", err)
		json.NewEncoder(w).Encode("internal server error")
		return
	}

	json.NewEncoder(w).Encode("Succesfully updated message")

}

func ServeAllMsgFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside ServeAllMsgFile")
	w.Header().Set("Content-Type", "application/text")

	content, err := service.CreateFile()
	if err != nil {
		json.NewEncoder(w).Encode("Internal server error")
		return
	}

	fileName := fmt.Sprintf("attachment; filename=AllMessages_%s.json", utils.CreateTimeStamp())

	w.Header().Set("Content-Disposition", fileName)

	http.ServeContent(w, r, "", time.Now(), strings.NewReader(content))

}

// func ServeVideoFile(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Inside ServeAllMsgFile")
// 	w.Header().Set("Content-Type", "video/mp4")

// 	first, err := moviego.Load("/Users/tushargoyal/Downloads/WhatsApp Video 2022-11-27 at 14.19.52.mp4")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// A simple screenshot from the video.
// 	_, err = first.Screenshot(2, "./simple-screen.png")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	http.ServeFile(w, r, "/Users/tushargoyal/Downloads/WhatsApp Video 2022-11-27 at 14.19.52.mp4")
// }
