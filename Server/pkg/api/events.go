package api

import (
// "context"
 "encoding/json"
 "net/http"
 "strconv"
 "log"
 "time"
 "github.com/gorilla/mux"
 "Server/pkg/models"
 "Server/pkg/cache"

)
func (api *api) getAllEvents(w http.ResponseWriter, r *http.Request) {
    cacheKey := "events:all"

    // попытка получить данные из редис
    cached, err := api.redis.Get(cache.Ctx, cacheKey).Result()
    if err == nil && cached != "" {
		log.Println("cached data events:all is used")
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(cached))
        return
    }
	//если кэша нет, то берем из базы
	log.Println("cached data events:all is not used")
    events, err := api.db.GetEvents()
    if err != nil {
		log.Println("error in getting data from db")
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    //записываем в джейсон, для кэширования в редис
    response, err := json.Marshal(events)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	//кэшируем на 5 минут
    err = api.redis.Set(cache.Ctx, cacheKey, response, 5*time.Minute).Err()
    if err != nil {
		log.Println("error in setting cache")
		log.Println(err.Error())
    }
	//отправляем данные клиенту
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}


func (api *api) getEventsByID(w http.ResponseWriter, r *http.Request){
	userID, ok := r.Context().Value("userID").(int) //приводим к нужному типу
	if !ok {
		http.Error(w, "Missing or invalid userID in context", http.StatusInternalServerError)
		return
	}

	cacheKey := "events:"+strconv.Itoa(userID)
	cached, err := api.redis.Get(cache.Ctx, cacheKey).Result()
	if err == nil && cached != "" {
		log.Println("cached data events:byID is used")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cached))
		return
	}

	log.Println("cached data events:byID is not used")
	data, err := api.db.GetEventsByID(userID)
  	if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
  	}	

	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//кэшируем на 5 минут
    err = api.redis.Set(cache.Ctx, cacheKey, response, 5*time.Minute).Err()
    if err != nil {
		log.Println("error in setting cache")
		log.Println(err.Error())
    }
	//отправляем данные клиенту
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}


func (api *api) newEvent(w http.ResponseWriter, r *http.Request){
	userID, ok := r.Context().Value("userID").(int) //приводим к нужному типу
	if !ok {
		http.Error(w, "Missing or invalid userID in context", http.StatusInternalServerError)
		return
	}
	var events models.Event
	//принимает указатель на структуру 
	err:=json.NewDecoder(r.Body).Decode(&events)
	if err!=nil {
		log.Println("error in parsing params")
		http.Error(w, "error in parsing params", http.StatusInternalServerError)
		return
	}
	id, err := api.db.NewEvent(events, userID)
	if err!=nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// удаление кеша для events:all
	err = api.redis.Del(cache.Ctx, "events:all").Err();
    if  err != nil {
        log.Println("error deleting events:all cache: " + err.Error())
    }
    // удаление кеша для events:byID
    userCacheKey := "events:" + strconv.Itoa(userID)
	err = api.redis.Del(cache.Ctx, userCacheKey).Err()
    if  err != nil {
        log.Println("error deleting " + userCacheKey + " cache: " + err.Error())
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

	// удаление кеша для events:all
	err = api.redis.Del(cache.Ctx, "events:all").Err();
    if  err != nil {
        log.Println("error deleting events:all cache: " + err.Error())
    }
    // удаление кеша для events:byID
    userCacheKey := "events:" + strconv.Itoa(events.IDus)
	err = api.redis.Del(cache.Ctx, userCacheKey).Err()
    if  err != nil {
        log.Println("error deleting " + userCacheKey + " cache: " + err.Error())
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
	userID,err := api.db.DeleteEvent(id)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// удаление кеша для events:all
	err = api.redis.Del(cache.Ctx, "events:all").Err();
    if  err != nil {
        log.Println("error deleting events:all cache: " + err.Error())
    }
    // удаление кеша для events:byID
    userCacheKey := "events:" + strconv.Itoa(userID)
	err = api.redis.Del(cache.Ctx, userCacheKey).Err()
    if  err != nil {
        log.Println("error deleting " + userCacheKey + " cache: " + err.Error())
    }
}





