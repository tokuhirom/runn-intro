package main

import (
	"context"
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/k1LoW/runn"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func TestUserAPI(t *testing.T) {
	// テスト用サーバーを起動
	db := setupTestDB(t)
	defer db.Close()
	
	srv := httptest.NewServer(NewApp(db))
	defer srv.Close()

	// runnでテストを実行
	opts := []runn.Option{
		runn.T(t),
		runn.Runner("req", srv.URL),
	}
	
	o, err := runn.Load("user-api-test.yml", opts...)
	if err != nil {
		t.Fatal(err)
	}

	if err := o.RunN(context.Background()); err != nil {
		t.Fatal(err)
	}
}