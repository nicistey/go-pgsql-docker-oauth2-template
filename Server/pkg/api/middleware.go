package api

import (
	"log"
	"net/http"
)


// /api
// /events/


func (api *api) middleware(next http.Handler) http.Handler { // Определяет middleware функцию как метод структуры api, принимающую следующий обработчик (next) и возвращающую новый обработчик.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // Создает новый обработчик с помощью http.HandlerFunc.
	 log.Println(r.URL.Path, r.Method) // Логирует путь и метод запроса.
	 w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	 //Access-Control-Allow-Origin
	 w.Header().Set("Access-Control-Allow-Origin", "*") // Разрешаем все origins (в продакшене нужно указать конкретные) 
	 w.Header().Set("Content-Type", "application/json") // Устанавливает заголовок Content-Type для ответа в формате JSON.
	 w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	 //log.Println(r.Method)
	 //из за этой падлы на работало ничего если кому то нужно будет объяснить, то напишите
	 if r.Method == "OPTIONS" {
		log.Println("!OPTIONS!")
		return
	}  
	 next.ServeHTTP(w, r) // Вызывает следующий обработчик в цепочке.
	})
}
   