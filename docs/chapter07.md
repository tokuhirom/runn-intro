---
layout: default
title: 第7章：Goテストヘルパー編
---

# 第7章：Goテストヘルパー編

**この章は本書の核心です。** runnの真の力は、Goテストヘルパーとして使用することで発揮されます。CLIツールとしての使用も便利ですが、Goのテストフレームワークと統合することで、より強力で柔軟なテスト環境を構築できます。

## なぜGoテストヘルパーなのか？

### 従来のAPIテストの課題

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

### runnを使った場合

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
# testdata/user_test.yml
desc: ユーザーAPI テスト
steps:
  create_user:
    req:
      api:///users:
        post:
          body:
            application/json:
              name: "Alice"
              email: "alice@example.com"
    test: |
      current.res.status == 201 &&
      current.res.body.name == "Alice" &&
      current.res.body.id > 0

  get_user:
    req:
      api:///users/{{ steps.create_user.res.body.id }}:
        get:
    test: |
      current.res.status == 200 &&
      current.res.body.name == "Alice"
```

## 基本的な統合方法

### プロジェクト構造

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

### 基本的なテストセットアップ

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

## 高度な統合パターン

### テストごとの独立したデータベース

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

### モックサーバーとの統合

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

### 複雑なテストデータの準備

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

## 実践的なテストパターン

### 認証フローのテスト

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
# testdata/auth_flow.yml
desc: 認証フローの完全テスト
steps:
  # ユーザー登録
  register_user:
    req:
      api:///auth/register:
        post:
          body:
            application/json:
              username: "{{ vars.test_username }}"
              password: "{{ vars.test_password }}"
              email: "test@example.com"
    test: current.res.status == 201

  # ログイン
  login:
    req:
      api:///auth/login:
        post:
          body:
            application/json:
              username: "{{ vars.test_username }}"
              password: "{{ vars.test_password }}"
    test: |
      current.res.status == 200 &&
      current.res.body.token != null &&
      current.res.body.expires_in > 0

  # 認証が必要なエンドポイントへのアクセス
  access_protected:
    req:
      api:///profile:
        get:
          headers:
            Authorization: "Bearer {{ steps.login.res.body.token }}"
    test: |
      current.res.status == 200 &&
      current.res.body.username == vars.test_username

  # トークンの検証
  verify_token:
    dump:
      # JWTトークンをデコード（実際の実装では適切なライブラリを使用）
      token_payload: |
        fromBase64(split(steps.login.res.body.token, ".")[1])
    test: |
      current.token_payload.username == vars.test_username

  # 無効なトークンでのアクセス
  invalid_token_access:
    req:
      api:///profile:
        get:
          headers:
            Authorization: "Bearer invalid-token"
    test: current.res.status == 401

  # ログアウト
  logout:
    req:
      api:///auth/logout:
        post:
          headers:
            Authorization: "Bearer {{ steps.login.res.body.token }}"
    test: current.res.status == 200

  # ログアウト後のアクセス
  access_after_logout:
    req:
      api:///profile:
        get:
          headers:
            Authorization: "Bearer {{ steps.login.res.body.token }}"
    test: current.res.status == 401
```

### E2Eワークフローテスト

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
# testdata/e2e_workflow.yml
desc: E2Eワークフローテスト - 顧客登録から注文完了まで
steps:
  # 1. 顧客登録
  register_customer:
    req:
      api:///customers:
        post:
          body:
            application/json: "{{ vars.test_customer }}"
    test: |
      current.res.status == 201 &&
      current.res.body.id > 0

  # 2. 商品情報の取得
  get_product:
    req:
      api:///products/{{ vars.test_product_id }}:
        get:
    test: |
      current.res.status == 200 &&
      current.res.body.stock >= vars.test_quantity

  # 3. カートに商品を追加
  add_to_cart:
    req:
      api:///customers/{{ steps.register_customer.res.body.id }}/cart:
        post:
          body:
            application/json:
              product_id: "{{ vars.test_product_id }}"
              quantity: "{{ vars.test_quantity }}"
    test: current.res.status == 200

  # 4. 注文の作成
  create_order:
    req:
      api:///orders:
        post:
          body:
            application/json:
              customer_id: "{{ steps.register_customer.res.body.id }}"
              items:
                - product_id: "{{ vars.test_product_id }}"
                  quantity: "{{ vars.test_quantity }}"
                  price: "{{ steps.get_product.res.body.price }}"
    test: |
      current.res.status == 201 &&
      current.res.body.order_id != null &&
      current.res.body.total_amount == steps.get_product.res.body.price * vars.test_quantity

  # 5. 決済処理
  process_payment:
    req:
      payment:///payments:
        post:
          body:
            application/json:
              order_id: "{{ steps.create_order.res.body.order_id }}"
              amount: "{{ steps.create_order.res.body.total_amount }}"
              customer_id: "{{ steps.register_customer.res.body.id }}"
    test: |
      current.res.status == 200 &&
      current.res.body.status == "completed"

  # 6. 注文ステータスの更新
  update_order_status:
    req:
      api:///orders/{{ steps.create_order.res.body.order_id }}/payment:
        put:
          body:
            application/json:
              transaction_id: "{{ steps.process_payment.res.body.transaction_id }}"
              status: "paid"
    test: current.res.status == 200

  # 7. 在庫の確認
  verify_stock_update:
    req:
      api:///products/{{ vars.test_product_id }}:
        get:
    test: |
      current.res.body.stock == steps.get_product.res.body.stock - vars.test_quantity

  # 8. 通知の送信確認
  verify_notification:
    req:
      notification:///notifications:
        get:
          query:
            customer_id: "{{ steps.register_customer.res.body.id }}"
            type: "order_confirmation"
    test: |
      current.res.status == 200 &&
      len(current.res.body.notifications) > 0 &&
      current.res.body.notifications[0].order_id == steps.create_order.res.body.order_id

  # 9. データベースの整合性確認
  verify_database_consistency:
    db:
      db:///
        query: |
          SELECT 
            o.id as order_id,
            o.status,
            o.total_amount,
            c.name as customer_name,
            p.name as product_name,
            p.stock
          FROM orders o
          JOIN customers c ON o.customer_id = c.id
          JOIN order_items oi ON o.id = oi.order_id
          JOIN products p ON oi.product_id = p.id
          WHERE o.id = $1
        params:
          - "{{ steps.create_order.res.body.order_id }}"
    test: |
      len(current.rows) == 1 &&
      current.rows[0].status == "paid" &&
      current.rows[0].customer_name == vars.test_customer.name &&
      current.rows[0].stock == steps.get_product.res.body.stock - vars.test_quantity
```

## パフォーマンステスト

### 負荷テストの実装

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
# testdata/performance_test.yml
desc: APIパフォーマンステスト
steps:
  # 並行負荷テスト
  load_test:
    loop:
      count: "{{ vars.concurrent_users }}"
    include:
      path: ./performance/user_simulation.yml
      vars:
        user_id: "{{ i }}"
        requests_count: "{{ vars.requests_per_user }}"
```

```yaml
# testdata/performance/user_simulation.yml
desc: 単一ユーザーのシミュレーション
steps:
  user_requests:
    loop:
      count: "{{ vars.requests_count }}"
    req:
      api:///users:
        get:
          query:
            page: "{{ (i % 100) + 1 }}"
            limit: 10
    test: |
      current.res.status == 200 &&
      current.res.response_time < 1000  # 1秒以内のレスポンス
```

## デバッグとトラブルシューティング

### デバッグ情報の出力

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

### テスト失敗時の詳細情報

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

## CI/CDとの統合

### GitHub Actionsでの実行

```yaml
# .github/workflows/api_test.yml
name: API Tests

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      postgres:
        image: postgres:13
        env:
          POSTGRES_PASSWORD: postgres
          POSTGRES_DB: testdb
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21
    
    - name: Install dependencies
      run: go mod download
    
    - name: Run database migrations
      run: |
        go run ./cmd/migrate up
      env:
        DATABASE_URL: postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable
    
    - name: Run API tests
      run: |
        go test -v ./... -tags=integration
      env:
        DATABASE_URL: postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable
        TEST_ENV: ci
    
    - name: Upload test results
      uses: actions/upload-artifact@v3
      if: always()
      with:
        name: test-results
        path: |
          test_screenshots/
          test_http_dumps/
          coverage.out
```

### Dockerを使った統合テスト

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

## まとめ

この章では、runnをGoテストヘルパーとして活用する方法について詳しく学びました：

1. **基本的な統合**: `runn.T(t)`を使ったGoテストとの統合
2. **高度な統合パターン**: 独立したデータベース、モックサーバーとの連携
3. **複雑なテストデータ準備**: 動的なテストデータ生成と管理
4. **実践的なテストパターン**: 認証フロー、E2Eワークフローのテスト
5. **パフォーマンステスト**: 負荷テストの実装と評価
6. **デバッグとトラブルシューティング**: 詳細なエラー情報の取得
7. **CI/CDとの統合**: GitHub Actions、Dockerとの連携

**runnの真の価値は、Goテストヘルパーとして使用することで発揮されます。** YAMLの宣言的な記述とGoの強力なテストフレームワークを組み合わせることで、保守性が高く、理解しやすい、そして強力なテストスイートを構築できます。

次章では、これまで学んだ知識を活用した実践的なプロジェクト例について見ていきます。

[第8章：実践編へ →](chapter08.md)