package api

import (
	"Server/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
func (api *api) getAllUsers(w http.ResponseWriter, r *http.Request){
	data, err := api.db.GetUsers()
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


func (api *api) getUserByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	stringID := vars["IDus"]
	id, err := strconv.Atoi(stringID)
	if err != nil {
	  http.Error(w, "Invalid user ID", http.StatusBadRequest) // Более информативное сообщение об ошибке
	  return
	}
	data, err := api.db.GetUserByID(id)
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


func (api *api) newUser(w http.ResponseWriter, r *http.Request){
	var user models.User
	//принимает указатель на структуру 
	err:=json.NewDecoder(r.Body).Decode(&user)
	if err!=nil {
		http.Error(w, "error in parsing params", http.StatusInternalServerError)
		return
	}

	id, err := api.db.NewUser(user)
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

  func (api *api) updateUser(w http.ResponseWriter, r *http.Request){
	var users models.User
	//принимает указатель на структуру 
	vars:= mux.Vars(r)
	stringID:= vars["IDus"]
	id,err :=strconv.Atoi(stringID)
	if err!=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err=json.NewDecoder(r.Body).Decode(&users)
	if err!=nil {
		http.Error(w, "error in parsing params", http.StatusInternalServerError)
		return
	}

	id, err = api.db.UpdateUser(id,users)
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

  func (api *api) deleteUser(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	stringID, ok := vars["IDus"]
	if !ok {//если был указан id, то возвращаем по id
		return
	}
	id,err :=strconv.Atoi(stringID)
	if err!=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.db.DeleteUser(id)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}





