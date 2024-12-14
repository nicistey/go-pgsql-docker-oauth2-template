package config

import (
 "fmt"
 "os"

 "github.com/joho/godotenv"
)

type Config struct {
 DBConnString      string
 GoogleClientID    string
 GoogleClientSecret string
 GoogleRedirectURL string
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
 }
//проверка на пустоту
 if config.DBConnString == "" || config.GoogleClientID == "" || config.GoogleClientSecret == "" || config.GoogleRedirectURL == "" {
  return nil, fmt.Errorf("missing required environment variables")
 }

 return config, nil
}