package config

import (
 "fmt"
 "os"

 "github.com/joho/godotenv"
)

type Config struct {
 DBConnString      	string
 GoogleClientID    	string
 GoogleClientSecret string
 GoogleRedirectURL 	string
 SecretJWTKey 		string
ReddisAddr 		    string
RedisPassword 	    string
}
//подключение (чтение) файла .env
func LoadConfig() (*Config, error) {
 err := godotenv.Load()
 if err != nil {
  return nil, fmt.Errorf("failed to load .env file: %w", err)
 }

 config := &Config{
  DBConnString:      os.Getenv("CONN_STR_DB"),
  GoogleClientID:    os.Getenv("GOOGLE_CLIENT_ID"),
  GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
  GoogleRedirectURL: os.Getenv("GOOGLE_REDIRECT_URL"),
  SecretJWTKey: os.Getenv("JWT_SECRET_KEY"),
  ReddisAddr: os.Getenv("REDIS_ADDR"),
  RedisPassword: os.Getenv("REDIS_PASSWORD"),
 }
//проверка на пустоту
 if config.DBConnString == "" || config.GoogleClientID == "" || config.GoogleClientSecret == "" || config.GoogleRedirectURL == ""|| config.SecretJWTKey == ""  || config.ReddisAddr == "" {
  return nil, fmt.Errorf("missing required environment variables")
 }

 return config, nil
}