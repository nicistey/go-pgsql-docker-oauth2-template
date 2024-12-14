package api

import (
	"Server/pkg/models"
	"Server/config"
	"context"       
	"encoding/json" 
	"fmt"           
	"log"
	"net/http" 
	"golang.org/x/oauth2"        
	"golang.org/x/oauth2/google" 
)

func getOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
	 ClientID:     cfg.GoogleClientID,
	 ClientSecret: cfg.GoogleClientSecret,
	 RedirectURL:  cfg.GoogleRedirectURL,
	 Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	 Endpoint:     google.Endpoint,
	}
   }

type GoogleUser struct {
	ID    string `json:"sub"`
	Name  string `json:"name"`
	Email string `json:"email"`
   }
//перенаправляет на страницу атентификации
func (api *api)handleGoogleAuth(w http.ResponseWriter, r *http.Request) { 
	oGoogleAuthConfig := getOAuthConfig(api.cfg) 
	url := oGoogleAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
//сюда приходит ответ
func (api *api) handleGoogleCallback(w http.ResponseWriter, r *http.Request) { 
	oGoogleAuthConfig := getOAuthConfig(api.cfg)
	code := r.FormValue("code") 
	if code == "" {            
		http.Error(w, "Error1!", http.StatusBadRequest) 
		return                                        
	}

	token, err := oGoogleAuthConfig.Exchange(context.Background(), code) 
	if err != nil {                                             
		http.Error(w, fmt.Sprintf("Error2!: %v", err), http.StatusInternalServerError) 
		return                                                             
	}

	client := oGoogleAuthConfig.Client(context.Background(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo") 
	if err != nil {                                                              
		http.Error(w, fmt.Sprintf("Не смог получить инфу юзера: %v", err), http.StatusInternalServerError) 
		return                                                                     
	}
	defer userInfo.Body.Close() 

	var googleUser GoogleUser 
	err = json.NewDecoder(userInfo.Body).Decode(&googleUser) 
	if err != nil {                                   
		http.Error(w, fmt.Sprintf("Не смог декодировать данные в декоде: %v", err), http.StatusInternalServerError) 
		return                                                                          
	}
	existsUserByGoogleID, err :=api.db.CheckGoogleIDExists(googleUser.ID)
	if err != nil {                                   
		http.Error(w, fmt.Sprintf("Ошибка в проверке пользователя по GoogleID: %v", err), http.StatusInternalServerError) 
		return                                                                         
	}
	log.Println(existsUserByGoogleID)
	if (!existsUserByGoogleID){
		var user models.User
		user.GoogleID = googleUser.ID
		user.Name = googleUser.Name
		user.Email=googleUser.Email
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
	}else{
		log.Println("Существует пользователь  с GoogleID: " +googleUser.ID )
	}

	
}

