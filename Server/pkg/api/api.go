package api

import (
	"Server/pkg/repository"
	"Server/config"
	"net/http"
	"github.com/gorilla/mux"
)

type api struct {
	r   *mux.Router
	db  *repository.PGRepo
	cfg *config.Config 
	
   }

func New(router *mux.Router, db *repository.PGRepo, cfg *config.Config) *api {
	return &api{r: router, db: db, cfg: cfg}
}



func (api *api) Hadle(cfg *config.Config) {
	//прописали, что функция принимает только гет метод прямо в инициализации, можно перечислять.Queries объясняем что возможно будем педавать
	api.r.HandleFunc("/auth", api.handleGoogleAuth).Methods(http.MethodGet,http.MethodOptions)
	api.r.HandleFunc("/auth/callback", api.handleGoogleCallback).Methods(http.MethodGet,http.MethodOptions)
	
	
	api.r.HandleFunc("/api/users", api.getAllUsers).Methods(http.MethodGet,http.MethodOptions)        // получение всех юзеров
	api.r.HandleFunc("/api/users/{IDus}", api.getUserByID).Methods(http.MethodGet,http.MethodOptions) // получение юзера по ID
	api.r.HandleFunc("/api/users", api.newUser).Methods(http.MethodPost,http.MethodOptions)           // создание нового юзера
	api.r.HandleFunc("/api/users/{IDus}", api.updateUser).Methods(http.MethodPost,http.MethodOptions)
	api.r.HandleFunc("/api/users/{IDus}", api.deleteUser).Methods(http.MethodDelete,http.MethodOptions)              // удаление юзера
	api.r.HandleFunc("/api/events", api.getAllEvents).Methods(http.MethodGet, http.MethodOptions) // получение всех ивентов
	api.r.HandleFunc("/api/eventsByID", api.getEventsByID).Methods(http.MethodGet, http.MethodOptions)
	api.r.HandleFunc("/api/events", api.newEvent).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/events/{IDev}", api.updateEvent).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/events/{IDev}", api.deleteEvent).Methods(http.MethodDelete, http.MethodOptions)

	api.r.Use(api.middleware)
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
