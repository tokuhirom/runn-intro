package testutil

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"time"
)

// NewTestBlogServer creates a test server with go-httpbin and custom endpoints
func NewTestBlogServer() *httptest.Server {
	// State for users
	var (
		users      = make(map[int]User)
		usersMutex sync.RWMutex
		nextUserID = 1

		// State for posts
		posts      = make(map[int]Post)
		postsMutex sync.RWMutex
		nextPostID = 1

		// State for auth
		registeredUsers = make(map[string]RegisteredUser)
		authMutex       sync.RWMutex

		// Valid tokens
		validTokens = map[string]string{
			"test-token-123":   "test@example.com",
			"mock-jwt-token":   "user@example.com",
			"new-access-token": "test@example.com",
		}
	)

	// Create a handler function that we can reference
	var handler http.Handler
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle custom endpoints first
		switch {
		// ユーザー登録エンドポイント
		case r.URL.Path == "/register" && r.Method == "POST":
			var regReq struct {
				Email    string `json:"email"`
				Password string `json:"password"`
				Name     string `json:"name"`
			}

			if err := json.NewDecoder(r.Body).Decode(&regReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			authMutex.Lock()
			registeredUsers[regReq.Email] = RegisteredUser{
				Email:    regReq.Email,
				Password: regReq.Password,
				Name:     regReq.Name,
			}
			authMutex.Unlock()

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(map[string]interface{}{
				"id":    len(registeredUsers),
				"email": regReq.Email,
				"name":  regReq.Name,
			}); err != nil {
				slog.Error("failed to encode response", "error", err)
			}

		// トークンリフレッシュエンドポイント
		case r.URL.Path == "/refresh" && r.Method == "POST":
			var refreshReq struct {
				RefreshToken string `json:"refreshToken"`
			}

			if err := json.NewDecoder(r.Body).Decode(&refreshReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Simple refresh - accept any refresh token
			if refreshReq.RefreshToken != "" {
				resp := map[string]string{
					"accessToken": "new-access-token",
				}
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					slog.Error("failed to encode response", "error", err)
				}
			} else {
				http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
			}

		// ユーザー一覧取得エンドポイント
		case r.URL.Path == "/users" && r.Method == "GET":
			usersMutex.RLock()
			userList := make([]User, 0, len(users))
			for _, u := range users {
				userList = append(userList, u)
			}
			usersMutex.RUnlock()

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(userList); err != nil {
				slog.Error("failed to encode response", "error", err)
			}

		// ユーザー作成エンドポイント
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
			if err := json.NewEncoder(w).Encode(user); err != nil {
				slog.Error("failed to encode response", "error", err)
			}

		// ユーザー取得エンドポイント（ID指定）
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
			if err := json.NewEncoder(w).Encode(user); err != nil {
				slog.Error("failed to encode response", "error", err)
			}

		// ユーザー更新エンドポイント（ID指定）
		case strings.HasPrefix(r.URL.Path, "/users/") && r.Method == "PUT":
			// Extract user ID from path
			path := strings.TrimPrefix(r.URL.Path, "/users/")
			userID, err := strconv.Atoi(path)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusBadRequest)
				return
			}

			var updateReq User
			if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			usersMutex.Lock()
			if user, exists := users[userID]; exists {
				user.Name = updateReq.Name
				user.Email = updateReq.Email
				users[userID] = user
				usersMutex.Unlock()

				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(user); err != nil {
					slog.Error("failed to encode response", "error", err)
				}
			} else {
				usersMutex.Unlock()
				http.Error(w, "User not found", http.StatusNotFound)
			}

		// 認証エンドポイント
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
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					slog.Error("failed to encode response", "error", err)
				}
			} else {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			}

		// ログインエンドポイント
		case r.URL.Path == "/login" && r.Method == "POST":
			var loginReq struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}

			if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Check if user is registered
			authMutex.RLock()
			regUser, registered := registeredUsers[loginReq.Email]
			authMutex.RUnlock()

			if registered && regUser.Password == loginReq.Password {
				resp := map[string]string{
					"accessToken":  "test-token-123",
					"refreshToken": "refresh-token-456",
					"token":        "mock-jwt-token", // for backward compatibility
				}
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					slog.Error("failed to encode response", "error", err)
				}
			} else if loginReq.Email != "" && loginReq.Password != "" {
				// Accept any non-empty email/password for backward compatibility
				resp := map[string]string{
					"token": "mock-jwt-token",
				}
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					slog.Error("failed to encode response", "error", err)
				}
			} else {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			}

		// プロフィール取得エンドポイント
		case r.URL.Path == "/profile" && r.Method == "GET":
			// Check authorization header
			authHeader := r.Header.Get("Authorization")
			token := strings.TrimPrefix(authHeader, "Bearer ")

			if email, valid := validTokens[token]; valid {
				// Return profile based on token
				profile := map[string]interface{}{
					"id":    1,
					"name":  "テストユーザー",
					"email": email,
				}
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(profile); err != nil {
					slog.Error("failed to encode response", "error", err)
				}
			} else {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}

		// 投稿作成エンドポイント
		case strings.HasPrefix(r.URL.Path, "/posts") && r.Method == "POST":
			var post Post
			if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			postsMutex.Lock()
			post.ID = nextPostID
			post.CreatedAt = time.Now()
			nextPostID++
			posts[post.ID] = post
			postsMutex.Unlock()

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(post); err != nil {
				slog.Error("failed to encode response", "error", err)
			}

		// 投稿取得エンドポイント（ID指定）
		case strings.HasPrefix(r.URL.Path, "/posts/") && r.Method == "GET":
			// Extract post ID from path
			path := strings.TrimPrefix(r.URL.Path, "/posts/")
			postID, err := strconv.Atoi(path)
			if err != nil {
				http.Error(w, "Invalid post ID", http.StatusBadRequest)
				return
			}

			postsMutex.RLock()
			post, exists := posts[postID]
			postsMutex.RUnlock()

			if !exists {
				http.Error(w, "Post not found", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(post); err != nil {
				slog.Error("failed to encode response", "error", err)
			}

		// 投稿更新エンドポイント（ID指定）
		case strings.HasPrefix(r.URL.Path, "/posts/") && r.Method == "PUT":
			// Extract post ID from path
			path := strings.TrimPrefix(r.URL.Path, "/posts/")
			postID, err := strconv.Atoi(path)
			if err != nil {
				http.Error(w, "Invalid post ID", http.StatusBadRequest)
				return
			}

			var updateReq Post
			if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			postsMutex.Lock()
			if post, exists := posts[postID]; exists {
				post.Title = updateReq.Title
				post.Content = updateReq.Content
				post.UpdatedAt = time.Now()
				posts[postID] = post
				postsMutex.Unlock()

				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(post); err != nil {
					slog.Error("failed to encode response", "error", err)
				}
			} else {
				postsMutex.Unlock()
				http.Error(w, "Post not found", http.StatusNotFound)
			}

		// 投稿削除エンドポイント（ID指定）
		case strings.HasPrefix(r.URL.Path, "/posts/") && r.Method == "DELETE":
			// Extract post ID from path
			path := strings.TrimPrefix(r.URL.Path, "/posts/")
			postID, err := strconv.Atoi(path)
			if err != nil {
				http.Error(w, "Invalid post ID", http.StatusBadRequest)
				return
			}

			postsMutex.Lock()
			if _, exists := posts[postID]; exists {
				delete(posts, postID)
				postsMutex.Unlock()
				w.WriteHeader(http.StatusNoContent)
			} else {
				postsMutex.Unlock()
				http.Error(w, "Post not found", http.StatusNotFound)
			}

		// /api/posts配下の再帰的ハンドリング
		case strings.HasPrefix(r.URL.Path, "/api/posts") && r.Method != "":
			// Handle /api/posts/* paths by recursively calling this handler
			r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
			handler.ServeHTTP(w, r)
			return

		// テスト用エンドポイント
		case r.URL.Path == "/test" && r.Method == "GET":
			// Simple test endpoint
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(map[string]string{
				"message": "Test successful",
				"status":  "ok",
			}); err != nil {
				slog.Error("failed to encode response", "error", err)
			}

		// バージョン付きテスト用エンドポイント
		case strings.Contains(r.URL.Path, "/test") && r.Method == "GET":
			// Handle versioned test endpoints like /v1/test
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(map[string]string{
				"message": "Test successful",
				"version": strings.Split(r.URL.Path, "/")[1],
				"status":  "ok",
			}); err != nil {
				slog.Error("failed to encode response", "error", err)
			}

		// バージョン付きユーザー一覧取得エンドポイント
		case strings.Contains(r.URL.Path, "/users") && strings.Contains(r.URL.Path, "/v"):
			// Handle versioned user endpoints like /v1/users
			usersMutex.RLock()
			userList := make([]User, 0, len(users))
			for _, u := range users {
				userList = append(userList, u)
			}
			usersMutex.RUnlock()

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(userList); err != nil {
				slog.Error("failed to encode response", "error", err)
			}

		// その他（未定義パス）
		default:
			http.NotFound(w, r)
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

// Post represents a blog post
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// RegisteredUser represents a registered user with auth info
type RegisteredUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
