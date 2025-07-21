package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type App struct {
	db *sql.DB
}

func NewApp(db *sql.DB) http.Handler {
	app := &App{db: db}
	mux := http.NewServeMux()
	
	// ユーザー作成
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var user User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			
			result, err := app.db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			
			id, _ := result.LastInsertId()
			user.ID = int(id)
			
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(user)
		}
	})
	
	// ユーザー取得
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			id := r.URL.Path[len("/users/"):]
			
			var user User
			err := app.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
			if err != nil {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		}
	})
	
	return mux
}

func setupDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	
	_, err = db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	app := NewApp(db)
	
	fmt.Println("Server starting on :8081...")
	log.Fatal(http.ListenAndServe(":8081", app))
}