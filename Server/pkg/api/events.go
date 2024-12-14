package api

import (
	"Server/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"
	"log"
	"github.com/gorilla/mux"
)
func (api *api) getAllEvents(w http.ResponseWriter, r *http.Request){

	data, err := api.db.GetEvents()
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
}


func (api *api) getEventByID(w http.ResponseWriter, r *http.Request){
	// TODO: вынести в middleware
    w.Header().Set("Access-Control-Allow-Origin", "*")
    if r.Method == http.MethodOptions {
        return
    }
	vars := mux.Vars(r)
	stringID := vars["IDev"]
	id, err := strconv.Atoi(stringID)
	if err != nil {
	  http.Error(w, "Invalid Events ID", http.StatusBadRequest) // Более информативное сообщение об ошибке
	  return
	}
	data, err := api.db.GetEventByID(id)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err = json.NewEncoder(w).Encode(data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}


func (api *api) newEvent(w http.ResponseWriter, r *http.Request){
	var events models.Event
	//принимает указатель на структуру 
	err:=json.NewDecoder(r.Body).Decode(&events)
	if err!=nil {
		http.Error(w, "error in parsing params", http.StatusInternalServerError)
		return
	}

	id, err := api.db.NewEvent(events)
	if err!=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err=json.NewEncoder(w).Encode(id)
	if err!=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *api) updateEvent(w http.ResponseWriter, r *http.Request){
	var events models.Event
	//принимает указатель на структуру 
	vars:= mux.Vars(r)
	stringID:= vars["IDev"]
	id,err :=strconv.Atoi(stringID)
	log.Println(id)
	if err!=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err=json.NewDecoder(r.Body).Decode(&events)
	if err!=nil {
		http.Error(w, "error in parsing params", http.StatusInternalServerError)
		return
	}
	id, err = api.db.UpdateEvent(id,events)
	if err!=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err=json.NewEncoder(w).Encode(id)
	if err!=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *api) deleteEvent(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	stringID, ok := vars["IDev"]
	if !ok {//если был указан id, то возвращаем по id
		return
	}
	id,err :=strconv.Atoi(stringID)
	if err!=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.db.DeleteEvent(id)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}





