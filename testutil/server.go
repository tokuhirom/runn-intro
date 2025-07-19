package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
)

// TestServer creates a test server with go-httpbin compatible endpoints and custom endpoints
func NewTestServer() *httptest.Server {
	mux := http.NewServeMux()
	
	// State for users
	var (
		users = make(map[int]User)
		usersMutex sync.RWMutex
		nextUserID = 1
	)
	
	// go-httpbin compatible endpoints
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		response := map[string]interface{}{
			"args":    r.URL.Query(),
			"headers": getHeaders(r),
			"origin":  r.RemoteAddr,
			"url":     r.URL.String(),
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	
	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		
		var jsonData interface{}
		var data interface{}
		
		if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			json.Unmarshal(body, &jsonData)
			data = jsonData
		} else {
			data = string(body)
		}
		
		response := map[string]interface{}{
			"args":    r.URL.Query(),
			"data":    data,
			"headers": getHeaders(r),
			"origin":  r.RemoteAddr,
			"url":     r.URL.String(),
		}
		
		if jsonData != nil {
			response["json"] = jsonData
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		response := map[string]interface{}{
			"slideshow": map[string]interface{}{
				"author": "Yours Truly",
				"date":   "date of publication",
				"slides": []map[string]interface{}{
					{
						"title": "Wake up to WonderWidgets!",
						"type":  "all",
					},
					{
						"items": []string{
							"Why <em>WonderWidgets</em> are great",
							"Who <em>buys</em> WonderWidgets",
						},
						"title": "Overview",
						"type":  "all",
					},
				},
				"title": "Sample Slide Show",
			},
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	
	mux.HandleFunc("/bearer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Unauthorized",
			})
			return
		}
		
		token := strings.TrimPrefix(authHeader, "Bearer ")
		
		response := map[string]interface{}{
			"authenticated": true,
			"token":        token,
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	
	// Custom endpoints for runn examples
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			var user User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			
			usersMutex.Lock()
			user.ID = nextUserID
			nextUserID++
			users[user.ID] = user
			usersMutex.Unlock()
			
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(user)
			
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		// Extract user ID from path
		path := strings.TrimPrefix(r.URL.Path, "/users/")
		userID, err := strconv.Atoi(path)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		
		usersMutex.RLock()
		user, exists := users[userID]
		usersMutex.RUnlock()
		
		if !exists {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})
	
	mux.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		var authReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		
		if err := json.NewDecoder(r.Body).Decode(&authReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		// Simple auth check - accept any non-empty username/password
		if authReq.Username != "" && authReq.Password != "" {
			resp := map[string]string{
				"token": "test-token-123",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}
	})
	
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		var loginReq struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		
		if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		// Simple login check - accept any non-empty email/password
		if loginReq.Email != "" && loginReq.Password != "" {
			resp := map[string]string{
				"token": "mock-jwt-token",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}
	})
	
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		// Check authorization header
		authHeader := r.Header.Get("Authorization")
		validTokens := map[string]bool{
			"Bearer test-token-123": true,
			"Bearer mock-jwt-token": true,
		}
		
		if _, valid := validTokens[authHeader]; !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		
		// Return a mock profile
		profile := map[string]interface{}{
			"id":    1,
			"name":  "テストユーザー",
			"email": "test@example.com",
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(profile)
	})
	
	return httptest.NewServer(mux)
}

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// getHeaders extracts headers from request
func getHeaders(r *http.Request) map[string]string {
	headers := make(map[string]string)
	for k, v := range r.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}
	return headers
}