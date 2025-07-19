package testutil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"

	"github.com/mccutchen/go-httpbin/v2/httpbin"
)

// TestServer creates a test server with go-httpbin and custom endpoints
func NewTestServer() *httptest.Server {
	// Create go-httpbin app
	httpbinApp := httpbin.New()
	
	// State for users
	var (
		users = make(map[int]User)
		usersMutex sync.RWMutex
		nextUserID = 1
	)
	
	// Create a handler that combines go-httpbin and custom endpoints
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle custom endpoints first
		switch {
		case r.URL.Path == "/users" && r.Method == "POST":
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
			
		case strings.HasPrefix(r.URL.Path, "/users/") && r.Method == "GET":
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
			
		case r.URL.Path == "/auth" && r.Method == "POST":
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
			
		case r.URL.Path == "/login" && r.Method == "POST":
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
			
		case r.URL.Path == "/profile" && r.Method == "GET":
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
			
		default:
			// Delegate to go-httpbin for all other requests
			httpbinApp.ServeHTTP(w, r)
		}
	})
	
	return httptest.NewServer(handler)
}

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}