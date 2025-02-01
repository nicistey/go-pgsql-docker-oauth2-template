package api

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
)

func TestHandle(t *testing.T) {
    router := mux.NewRouter()
    api := New(router, nil, nil)
    api.Handle(nil)

    req, _ := http.NewRequest("GET", "/health", nil)
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
    }
}