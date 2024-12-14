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
	"time"
 	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID       int    `json:"id"` 
	GoogleID string `json:"google_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
   }

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

   func generateJWT(userID int, googleID string, name string, email string, secretKey string) (string, error) {
	claims := &Claims{
	 ID:       userID,
	 GoogleID: googleID,
	 Name:     name,
	 Email:    email,
	 RegisteredClaims: jwt.RegisteredClaims{
	  ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), 
	  IssuedAt:  jwt.NewNumericDate(time.Now()),
	 },
	}
   
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) 
   
	tokenString, err := token.SignedString([]byte(secretKey)) 
	if err != nil {
	 return "", err
	}
   
	return tokenString, nil
   }

func (api *api)handleGoogleAuth(w http.ResponseWriter, r *http.Request) { 
	oGoogleAuthConfig := getOAuthConfig(api.cfg) // получаем config из api
	url := oGoogleAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
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
	var user models.User
	if (!existsUserByGoogleID){
		
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
		user, err = api.db.GetUserByGoogleID(googleUser.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	secretKey := api.cfg.SecretJWTKey
	tokenString, err := generateJWT(user.IDus, googleUser.ID, googleUser.Name, googleUser.Email, secretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// отправка JWT клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

	
}

