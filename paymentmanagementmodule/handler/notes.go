package handler

import (
	pb "com.example.paymentmanagement/proto"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
	"net/http"
	"fmt"
	"net/url"
)

var notestableName = "notes"

func AddNote(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var note *pb.NotesRequest
	_ = json.NewDecoder(request.Body).Decode(&note)
	fmt.Print(note);
	e := dbClient.Table(notestableName).Create(&note)
	if e.Error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + e.Error.Error() + `" }`))
		return
	}
	logger.Infof("Inserted data")
	json.NewEncoder(response).Encode(note)
}

func ReadNoteById(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var note *pb.NotesResponse
	e := dbClient.Table(notestableName).First(&note, id)
	if e.Error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + e.Error.Error() + `" }`))
		return
	}
	logger.Infof("Fetched Note With Id:" + id)
	json.NewEncoder(response).Encode(note)
}

func GetNotes(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("content-type", "application/json")
	var notes []*pb.NotesResponse
	e := dbClient.Table(notestableName).Find(&notes)
	if e.Error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + e.Error.Error() + `" }`))
		return
	}
	logger.Infof("Fetched all Notes")
	json.NewEncoder(response).Encode(notes)
}

func UpdateNote(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var note *pb.NotesResponse
	var ev *pb.NotesResponse
	_ = json.NewDecoder(request.Body).Decode(&note)
	dbClient.Table(notestableName).First(&ev, note.Id)
	dbClient.Table(notestableName).Model(&ev).Updates(note)
	logger.Infof("Updated Event with Id:" + note.Id)
	json.NewEncoder(response).Encode("Updated Event with Id:" + note.Id)
}

func DeleteNote(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var note *pb.NotesResponse
	queryParams, err := url.ParseQuery(request.URL.RawQuery)
	if err != nil {
		http.Error(response, "Error parsing query parameters", http.StatusBadRequest)
		return
	}
	idValues, ok := queryParams["id"]
	if !ok || len(idValues) == 0 {
		http.Error(response, "Missing or empty 'id' parameter", http.StatusBadRequest)
		return
	}
	id := idValues[0]
	result := dbClient.Table(notestableName).Where("id = ?", id).Delete(&note)
	if result.RowsAffected == 0 {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`Unable to delete Please Verify`))
		return
	}
	logger.Infof("Deleted Note with Id:" + id)
	json.NewEncoder(response).Encode(result.RowsAffected)
}
