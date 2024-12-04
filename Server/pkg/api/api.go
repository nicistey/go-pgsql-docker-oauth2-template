package api

import (
	"Server/pkg/repository"
	"net/http"

	//"github.com/rs/cors"

	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type api struct {
	r  *mux.Router
	db *repository.PGRepo
}

func New(router *mux.Router, db *repository.PGRepo) *api {
	return &api{r: router, db: db}
}

func (api *api) Hadle() {
	//прописали, что функция принимает только гет метод прямо в инициализации, можно перечислять.Queries объясняем что возможно будем педавать
	api.r.HandleFunc("/api/users", api.getAllUsers).Methods(http.MethodGet,http.MethodOptions)        // получение всех юзеров
	api.r.HandleFunc("/api/users/{IDus}", api.getUserByID).Methods(http.MethodGet,http.MethodOptions) // получение юзера по ID
	api.r.HandleFunc("/api/users", api.newUser).Methods(http.MethodPost,http.MethodOptions)           // создание нового юзера
	api.r.HandleFunc("/api/users/{IDus}", api.updateUser).Methods(http.MethodPost,http.MethodOptions)
	api.r.HandleFunc("/api/users/{IDus}", api.deleteUser).Methods(http.MethodDelete,http.MethodOptions)              // удаление юзера
	api.r.HandleFunc("/api/events", api.getAllEvents).Methods(http.MethodGet, http.MethodOptions) // получение всех ивентов
	api.r.HandleFunc("/api/events/{IDev}", api.getEventByID).Methods(http.MethodGet, http.MethodOptions)
	api.r.HandleFunc("/api/events", api.newEvent).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/events/{IDev}", api.updateEvent).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/events/{IDev}", api.deleteEvent).Methods(http.MethodDelete, http.MethodOptions)

	api.r.Use(api.middleware)
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
