# 第7章：Goテストヘルパー編 - runnの最終兵器！

**ついに来た、本書のクライマックス！** これまでCLIツールとしてrunnを使ってきたが、それは**ほんの入り口**に過ぎない。**runnの真の力**は、Goテストヘルパーとして使用することで**爆発的に解放**される！Goの強力なテストフレームワークと**完璧に融合**し、**最強のテスト環境**を構築しよう！

## 🤔 なぜGoテストヘルパーなのか？ - 従来の方法の限界！

### 😩 従来のAPIテストの悲惨な現実

```go
// 従来のGoでのAPIテスト例
func TestUserAPI(t *testing.T) {
    // 1. テストサーバーの起動
    server := httptest.NewServer(handler)
    defer server.Close()
    
    // 2. HTTPクライアントの設定
    client := &http.Client{Timeout: 10 * time.Second}
    
    // 3. リクエストの作成
    reqBody := `{"name":"Alice","email":"alice@example.com"}`
    req, _ := http.NewRequest("POST", server.URL+"/users", strings.NewReader(reqBody))
    req.Header.Set("Content-Type", "application/json")
    
    // 4. リクエストの実行
    resp, err := client.Do(req)
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()
    
    // 5. レスポンスの検証
    if resp.StatusCode != 201 {
        t.Errorf("Expected 201, got %d", resp.StatusCode)
    }
    
    // 6. ボディの解析と検証
    var user User
    if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
        t.Fatal(err)
    }
    
    if user.Name != "Alice" {
        t.Errorf("Expected Alice, got %s", user.Name)
    }
    
    // さらに複雑な検証が続く...
}
```

### 🎆 runnを使った場合 - 革命的なシンプルさ！

```go
func TestUserAPI(t *testing.T) {
    // 1. テストサーバーの起動
    server := httptest.NewServer(handler)
    defer server.Close()
    
    // 2. runnでテストを実行
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", server.URL),
    }
    
    o, err := runn.Load("testdata/user_test.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}
```

```yaml
{{ includex("examples/test-helper/user_test.yml") }}
```

## 🚀 基本的な統合方法 - runnとGoの幸せな結婚！

### 📁 プロジェクト構造 - 理想的なディレクトリ構成

```
myproject/
├── main.go
├── handler.go
├── model.go
├── go.mod
├── go.sum
└── testdata/
    ├── scenarios/
    │   ├── user/
    │   │   ├── create_user.yml
    │   │   ├── update_user.yml
    │   │   └── delete_user.yml
    │   ├── auth/
    │   │   ├── login.yml
    │   │   └── logout.yml
    │   └── integration/
    │       └── full_workflow.yml
    └── fixtures/
        ├── users.json
        └── products.json
```

### 🔧 基本的なテストセットアップ - これがrunnテストの基本形！

```go
package main

import (
    "context"
    "database/sql"
    "net/http/httptest"
    "testing"
    
    "github.com/k1LoW/runn"
    _ "github.com/lib/pq"
)

func TestAPI(t *testing.T) {
    // データベースのセットアップ
    db := setupTestDB(t)
    defer db.Close()
    
    // テストサーバーの起動
    app := NewApp(db)
    server := httptest.NewServer(app.Handler())
    defer server.Close()
    
    // runnの設定
    opts := []runn.Option{
        runn.T(t),                          // テストコンテキストを渡す
        runn.Runner("api", server.URL),     // APIサーバーのURL
        runn.DBRunner("db", db),            // データベース接続
        runn.Var("test_user_email", "test@example.com"),
    }
    
    // シナリオの実行
    o, err := runn.Load("testdata/scenarios/**/*.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}

func setupTestDB(t *testing.T) *sql.DB {
    // テスト用データベースの設定
    db, err := sql.Open("postgres", "postgres://test:test@localhost/testdb?sslmode=disable")
    if err != nil {
        t.Fatal(err)
    }
    
    // マイグレーションの実行
    if err := runMigrations(db); err != nil {
        t.Fatal(err)
    }
    
    // テストデータの投入
    if err := seedTestData(db); err != nil {
        t.Fatal(err)
    }
    
    return db
}
```

## 🎯 高度な統合パターン - プロフェッショナルの極意！

### 🗄️ テストごとの独立したデータベース - 完全なアイソレーション！

```go
func TestUserCRUD(t *testing.T) {
    // 各テストで独立したデータベースを使用
    dbName := fmt.Sprintf("test_%d", time.Now().UnixNano())
    db := createTestDatabase(t, dbName)
    defer dropTestDatabase(t, dbName)
    defer db.Close()
    
    // アプリケーションの起動
    app := NewApp(db)
    server := httptest.NewServer(app.Handler())
    defer server.Close()
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", server.URL),
        runn.DBRunner("db", db),
        runn.Var("test_db_name", dbName),
    }
    
    o, err := runn.Load("testdata/user_crud.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}

func createTestDatabase(t *testing.T, dbName string) *sql.DB {
    // 管理用データベースに接続
    adminDB, err := sql.Open("postgres", "postgres://admin:admin@localhost/postgres?sslmode=disable")
    if err != nil {
        t.Fatal(err)
    }
    defer adminDB.Close()
    
    // テスト用データベースを作成
    _, err = adminDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
    if err != nil {
        t.Fatal(err)
    }
    
    // 新しいデータベースに接続
    testDB, err := sql.Open("postgres", fmt.Sprintf("postgres://admin:admin@localhost/%s?sslmode=disable", dbName))
    if err != nil {
        t.Fatal(err)
    }
    
    // スキーマの作成
    if err := createSchema(testDB); err != nil {
        t.Fatal(err)
    }
    
    return testDB
}
```

### 🎭 モックサーバーとの統合 - 外部APIを完璧にシミュレート！

```go
func TestExternalAPIIntegration(t *testing.T) {
    // 外部APIのモックサーバー
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/external/users":
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(200)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "users": []map[string]interface{}{
                    {"id": 1, "name": "External User 1"},
                    {"id": 2, "name": "External User 2"},
                },
            })
        case "/external/auth":
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(200)
            json.NewEncoder(w).Encode(map[string]string{
                "token": "mock-jwt-token",
                "expires_in": "3600",
            })
        default:
            w.WriteHeader(404)
        }
    }))
    defer mockServer.Close()
    
    // メインアプリケーション
    db := setupTestDB(t)
    defer db.Close()
    
    app := NewApp(db)
    app.SetExternalAPIURL(mockServer.URL) // 外部APIのURLを設定
    server := httptest.NewServer(app.Handler())
    defer server.Close()
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", server.URL),
        runn.Runner("external", mockServer.URL),
        runn.DBRunner("db", db),
    }
    
    o, err := runn.Load("testdata/external_integration.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}
```

### 🎆 複雑なテストデータの準備 - リアルなデータで真のテスト！

```go
func TestComplexScenario(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()
    
    // 複雑なテストデータの準備
    testData := prepareComplexTestData(t, db)
    
    server := httptest.NewServer(NewApp(db).Handler())
    defer server.Close()
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", server.URL),
        runn.DBRunner("db", db),
        
        // 準備したテストデータを変数として渡す
        runn.Var("admin_user_id", testData.AdminUser.ID),
        runn.Var("regular_user_id", testData.RegularUser.ID),
        runn.Var("test_products", testData.Products),
        runn.Var("test_categories", testData.Categories),
        
        // 動的に生成されたデータ
        runn.Var("test_orders", generateTestOrders(testData)),
    }
    
    o, err := runn.Load("testdata/complex_scenario.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}

type TestData struct {
    AdminUser   User
    RegularUser User
    Products    []Product
    Categories  []Category
}

func prepareComplexTestData(t *testing.T, db *sql.DB) *TestData {
    data := &TestData{}
    
    // 管理者ユーザーの作成
    adminUser := User{
        Name:     "Admin User",
        Email:    "admin@example.com",
        Role:     "admin",
        Active:   true,
    }
    if err := createUser(db, &adminUser); err != nil {
        t.Fatal(err)
    }
    data.AdminUser = adminUser
    
    // 一般ユーザーの作成
    regularUser := User{
        Name:     "Regular User",
        Email:    "user@example.com",
        Role:     "user",
        Active:   true,
    }
    if err := createUser(db, &regularUser); err != nil {
        t.Fatal(err)
    }
    data.RegularUser = regularUser
    
    // カテゴリの作成
    categories := []Category{
        {Name: "Electronics", Description: "Electronic products"},
        {Name: "Books", Description: "Books and publications"},
        {Name: "Clothing", Description: "Clothing and accessories"},
    }
    for i := range categories {
        if err := createCategory(db, &categories[i]); err != nil {
            t.Fatal(err)
        }
    }
    data.Categories = categories
    
    // 商品の作成
    products := []Product{
        {Name: "Laptop", Price: 999.99, CategoryID: categories[0].ID, Stock: 10},
        {Name: "Smartphone", Price: 599.99, CategoryID: categories[0].ID, Stock: 20},
        {Name: "Programming Book", Price: 49.99, CategoryID: categories[1].ID, Stock: 50},
        {Name: "T-Shirt", Price: 19.99, CategoryID: categories[2].ID, Stock: 100},
    }
    for i := range products {
        if err := createProduct(db, &products[i]); err != nil {
            t.Fatal(err)
        }
    }
    data.Products = products
    
    return data
}
```

## 💼 実践的なテストパターン - 現場で使える最強テクニック！

### 🔐 認証フローのテスト - セキュリティを完璧に検証！

```go
func TestAuthenticationFlow(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()
    
    // JWTシークレットキーを設定
    jwtSecret := "test-secret-key"
    
    app := NewApp(db)
    app.SetJWTSecret(jwtSecret)
    server := httptest.NewServer(app.Handler())
    defer server.Close()
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", server.URL),
        runn.DBRunner("db", db),
        runn.Var("jwt_secret", jwtSecret),
        runn.Var("test_username", "testuser"),
        runn.Var("test_password", "testpass123"),
    }
    
    o, err := runn.Load("testdata/auth_flow.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}
```

```yaml
{{ includex("examples/test-helper/auth_flow.yml") }}
```

### 🌐 E2Eワークフローテスト - システム全体を完全テスト！

```go
func TestE2EWorkflow(t *testing.T) {
    // 複数のサービスを起動
    db := setupTestDB(t)
    defer db.Close()
    
    // メインAPIサーバー
    mainApp := NewMainApp(db)
    mainServer := httptest.NewServer(mainApp.Handler())
    defer mainServer.Close()
    
    // 通知サービス
    notificationApp := NewNotificationApp()
    notificationServer := httptest.NewServer(notificationApp.Handler())
    defer notificationServer.Close()
    
    // 決済サービス（モック）
    paymentServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 決済処理のモック
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "transaction_id": "txn_" + generateRandomID(),
            "status": "completed",
            "amount": 1000,
        })
    }))
    defer paymentServer.Close()
    
    // 外部APIの設定
    mainApp.SetPaymentServiceURL(paymentServer.URL)
    mainApp.SetNotificationServiceURL(notificationServer.URL)
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", mainServer.URL),
        runn.Runner("notification", notificationServer.URL),
        runn.Runner("payment", paymentServer.URL),
        runn.DBRunner("db", db),
        
        // テストデータ
        runn.Var("test_customer", map[string]interface{}{
            "name":  "Test Customer",
            "email": "customer@example.com",
            "phone": "+81-90-1234-5678",
        }),
        runn.Var("test_product_id", 1),
        runn.Var("test_quantity", 2),
    }
    
    o, err := runn.Load("testdata/e2e_workflow.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}
```

```yaml
{{ includex("examples/test-helper/e2e_workflow.yml") }}
```

## ⚡ パフォーマンステスト - 速度の限界に挑戦！

### 🚀 負荷テストの実装 - システムの耐久力を測れ！

```go
func TestAPIPerformance(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping performance test in short mode")
    }
    
    db := setupTestDB(t)
    defer db.Close()
    
    // パフォーマンステスト用のデータを大量に準備
    preparePerformanceTestData(t, db, 10000) // 10,000件のテストデータ
    
    app := NewApp(db)
    server := httptest.NewServer(app.Handler())
    defer server.Close()
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", server.URL),
        runn.DBRunner("db", db),
        runn.Var("concurrent_users", 50),
        runn.Var("requests_per_user", 100),
    }
    
    o, err := runn.Load("testdata/performance_test.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    start := time.Now()
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
    duration := time.Since(start)
    
    t.Logf("Performance test completed in %v", duration)
    
    // パフォーマンス要件の検証
    totalRequests := 50 * 100 // concurrent_users * requests_per_user
    rps := float64(totalRequests) / duration.Seconds()
    
    if rps < 100 { // 最低100 RPS要求
        t.Errorf("Performance requirement not met: %.2f RPS (required: 100 RPS)", rps)
    }
    
    t.Logf("Achieved %.2f requests per second", rps)
}
```

```yaml
{{ includex("examples/test-helper/performance_test.yml") }}
```

```yaml
{{ includex("examples/test-helper/performance/user_simulation.yml") }}
```

## 🔍 デバッグとトラブルシューティング - 問題を瞬時に解決！

### 📝 デバッグ情報の出力 - すべてを可視化せよ！

```go
func TestWithDebug(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()
    
    server := httptest.NewServer(NewApp(db).Handler())
    defer server.Close()
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", server.URL),
        runn.DBRunner("db", db),
        
        // デバッグ情報を有効化
        runn.Debug(true),
        runn.Verbose(true),
        
        // カスタムログ出力
        runn.Logger(log.New(os.Stdout, "[RUNN] ", log.LstdFlags)),
    }
    
    o, err := runn.Load("testdata/debug_test.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}
```

### 💥 テスト失敗時の詳細情報 - 失敗から学べ！

```go
func TestWithDetailedErrorInfo(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()
    
    server := httptest.NewServer(NewApp(db).Handler())
    defer server.Close()
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", server.URL),
        runn.DBRunner("db", db),
        
        // 失敗時のスクリーンショット保存
        runn.ScreenshotDir("./test_screenshots"),
        
        // 失敗時のHTTPダンプ保存
        runn.HTTPDumpDir("./test_http_dumps"),
    }
    
    o, err := runn.Load("testdata/scenarios/**/*.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        // 詳細なエラー情報を出力
        if runnErr, ok := err.(*runn.Error); ok {
            t.Logf("Failed step: %s", runnErr.StepName)
            t.Logf("Error details: %s", runnErr.Details)
            t.Logf("Request: %s", runnErr.Request)
            t.Logf("Response: %s", runnErr.Response)
        }
        t.Fatal(err)
    }
}
```

## 🔄 CI/CDとの統合 - 自動化の極み！

### 💙 GitHub Actionsでの実行 - クラウドで最強テスト！

```yaml
{{ includex("examples/test-helper/github_actions.yml") }}
```

### 🐳 Dockerを使った統合テスト - コンテナで完璧なテスト環境！

```go
func TestWithDocker(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping Docker integration test in short mode")
    }
    
    // Docker Composeでテスト環境を起動
    cmd := exec.Command("docker-compose", "-f", "docker-compose.test.yml", "up", "-d")
    if err := cmd.Run(); err != nil {
        t.Fatal(err)
    }
    
    // テスト終了時にクリーンアップ
    defer func() {
        cmd := exec.Command("docker-compose", "-f", "docker-compose.test.yml", "down", "-v")
        cmd.Run()
    }()
    
    // サービスの起動を待機
    time.Sleep(10 * time.Second)
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", "http://localhost:8080"),
        runn.Runner("db", "postgres://test:test@localhost:5433/testdb?sslmode=disable"),
    }
    
    o, err := runn.Load("testdata/docker_integration.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}
```

## 🎆 まとめ - Goテストヘルパーの達人誕生！

**素晴らしい！** あなたは今、**runnの最強の使い方をマスター**した！

### 🏆 この章で習得した7つの極意：

1. **🔗 基本的な統合**: `runn.T(t)`で**Goテストと完璧な融合**！
2. **🎯 高度な統合パターン**: 独立データベース、モックサーバーで**理想的なテスト環境**！
3. **🎆 複雑なテストデータ準備**: 動的生成で**リアルなテストシナリオ**！
4. **💼 実践的なテストパターン**: 認証フロー、E2Eワークフローを**完璧にカバー**！
5. **⚡ パフォーマンステスト**: **負荷の限界を見極めろ**！
6. **🔍 デバッグとトラブルシューティング**: エラーを**瞬時に特定**！
7. **🔄 CI/CDとの統合**: GitHub Actions、Dockerで**完全自動化**！

**runnの真の価値**は、Goテストヘルパーとして使用することで**爆発的に発揮**される。YAMLの**宣言的な美しさ**とGoの**圧倒的なパワー**が融合し、**史上最強のテストスイート**が誕生する！

あなたはもう、**単なるテスターではない**。**テストの芸術家**だ！
