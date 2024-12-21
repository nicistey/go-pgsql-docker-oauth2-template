package api

import (
	"Server/config"
	"Server/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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

    code := r.URL.Query().Get("code")
    //state := r.URL.Query().Get("state")

    if code == "" {
        http.Error(w, "Error: Missing code parameter", http.StatusBadRequest)
        return
    }

    token, err := oGoogleAuthConfig.Exchange(context.Background(), code)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error exchanging code for token: %v, Code: %s", err, code), http.StatusInternalServerError)
        log.Printf("Google OAuth2 error: %v, Code: %s", err, code)
        return
    }

    client := oGoogleAuthConfig.Client(context.Background(), token)
    userInfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
    if err != nil {
        http.Error(w, fmt.Sprintf("Error getting user info: %v", err), http.StatusInternalServerError)
        return
    }
    defer userInfo.Body.Close()

    var googleUser GoogleUser
    err = json.NewDecoder(userInfo.Body).Decode(&googleUser)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error decoding user info: %v", err), http.StatusInternalServerError)
        return
    }
	existsUserByGoogleID, err := api.db.CheckGoogleIDExists(googleUser.ID)
    if err != nil {
        http.Error(w, fmt.Sprintf("Ошибка в проверке пользователя по GoogleID: %v", err), http.StatusInternalServerError)
        return
    }

    var user models.User
    if !existsUserByGoogleID {
        user.GoogleID = googleUser.ID
        user.Name = googleUser.Name
        user.Email = googleUser.Email
        _, err := api.db.NewUser(user)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        //  Этот код уже был - сохраним его!
        // err = json.NewEncoder(w).Encode(id) // удаляем эту строку!

    } else {
        log.Println("Существует пользователь  с GoogleID: " + googleUser.ID)
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

    // Redirect to frontend with token as a query parameter
    redirectURL := fmt.Sprintf("http://localhost:8080/callback?token=%s", url.QueryEscape(tokenString))
    http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}