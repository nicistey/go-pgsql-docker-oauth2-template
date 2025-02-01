package main

import (
    "os"
    "Server/config"
    "testing"
)

func TestLoadConfig(t *testing.T) {
    os.Setenv("CONN_STR_DB", "test_db")
    os.Setenv("GOOGLE_CLIENT_ID", "test_client_id")
    os.Setenv("GOOGLE_CLIENT_SECRET", "test_client_secret")
    os.Setenv("GOOGLE_REDIRECT_URL", "test_redirect_url")
    os.Setenv("JWT_SECRET_KEY", "test_jwt")

    cfg, err := config.LoadConfig()
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if cfg.DBConnString != "test_db" {
        t.Errorf("Expected %v, got %v", "test_db", cfg.DBConnString)
    }
}