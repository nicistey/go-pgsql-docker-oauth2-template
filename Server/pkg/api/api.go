package api

import (
	"Server/pkg/repository"
    "Server/config"
    "net/http"
    "github.com/gorilla/mux"
	"github.com/go-redis/redis/v8"
)

type api struct {
	r   *mux.Router
	db  *repository.PGRepo
	cfg *config.Config 
    redis *redis.Client
}

func New(router *mux.Router, db *repository.PGRepo, cfg *config.Config, redisClient *redis.Client) *api {
	return &api{r: router, db: db, cfg: cfg,redis: redisClient}
}



func (api *api) Handle(cfg *config.Config) {
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
	api.r.HandleFunc("/health", api.health).Methods(http.MethodGet)
	api.r.Use(api.middleware)
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}

func (api *api) health(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}