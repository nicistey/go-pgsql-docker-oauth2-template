package repository

import (
    "context"
    "testing"

    "github.com/jackc/pgx/v4/pgxpool"
)

const connStrTest = "postgres://nicis:123@localhost:5433/kudago" // порт стоит 5433 так как в докере переброс портов
//и бд которая работает в докере по порту 5432, для локального запуска тестов нужно запустить по порту 5432

func TestSelectOne(t *testing.T) {
    pool, err := pgxpool.Connect(context.Background(), connStrTest)
    if err != nil {
        t.Fatalf("Unable to connect to database: %v", err)
    }
    defer pool.Close()

    var result int
    err = pool.QueryRow(context.Background(), "SELECT 1").Scan(&result)
    if err != nil {
        t.Fatalf("QueryRow failed: %v", err)
    }

    if result != 1 {
        t.Errorf("Expected 1, got %d", result)
    }
}