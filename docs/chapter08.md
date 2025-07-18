# 第8章：実践編

この章では、実際のプロジェクトでrunnを活用する具体的な例を通して、実践的な使い方を学んでいきます。よくあるユースケースから始まり、ベストプラクティス、デバッグ方法、パフォーマンスチューニングまでを網羅します。

## よくあるユースケース

### 1. RESTful APIの包括的テスト

実際のECサイトAPIを例に、包括的なテストスイートを構築してみましょう。

#### プロジェクト構造

```
ecommerce-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   └── repository/
├── testdata/
│   ├── scenarios/
│   │   ├── auth/
│   │   │   ├── login.yml
│   │   │   ├── register.yml
│   │   │   └── password_reset.yml
│   │   ├── products/
│   │   │   ├── crud.yml
│   │   │   ├── search.yml
│   │   │   └── inventory.yml
│   │   ├── orders/
│   │   │   ├── create_order.yml
│   │   │   ├── order_lifecycle.yml
│   │   │   └── payment_flow.yml
│   │   └── integration/
│   │       ├── user_journey.yml
│   │       └── admin_workflow.yml
│   ├── fixtures/
│   │   ├── users.json
│   │   ├── products.json
│   │   └── categories.json
│   └── sql/
│       ├── schema.sql
│       └── seed.sql
└── api_test.go
```

#### メインテストファイル

```go
// api_test.go
package main

import (
    "context"
    "database/sql"
    "fmt"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/k1LoW/runn"
    _ "github.com/lib/pq"
)

func TestECommerceAPI(t *testing.T) {
    // テスト環境のセットアップ
    testEnv := setupTestEnvironment(t)
    defer testEnv.Cleanup()

    // runnの設定
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", testEnv.ServerURL),
        runn.DBRunner("db", testEnv.DB),
        
        // テストデータ
        runn.Var("admin_email", "admin@example.com"),
        runn.Var("admin_password", "admin123"),
        runn.Var("test_user_email", "user@example.com"),
        runn.Var("test_user_password", "user123"),
        
        // 設定値
        runn.Var("jwt_secret", testEnv.JWTSecret),
        runn.Var("payment_api_key", "test_payment_key"),
    }

    // 全シナリオを実行
    o, err := runn.Load("testdata/scenarios/**/*.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }

    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}

type TestEnvironment struct {
    DB        *sql.DB
    Server    *httptest.Server
    ServerURL string
    JWTSecret string
    cleanup   []func()
}

func (te *TestEnvironment) Cleanup() {
    for _, fn := range te.cleanup {
        fn()
    }
}

func setupTestEnvironment(t *testing.T) *TestEnvironment {
    env := &TestEnvironment{
        JWTSecret: "test-jwt-secret-key",
    }

    // テスト用データベースの作成
    dbName := fmt.Sprintf("test_ecommerce_%d", time.Now().UnixNano())
    db := createTestDatabase(t, dbName)
    env.DB = db
    env.cleanup = append(env.cleanup, func() {
        db.Close()
        dropTestDatabase(t, dbName)
    })

    // スキーマとテストデータの投入
    if err := loadSchema(db, "testdata/sql/schema.sql"); err != nil {
        t.Fatal(err)
    }
    if err := loadTestData(db, "testdata/sql/seed.sql"); err != nil {
        t.Fatal(err)
    }

    // アプリケーションサーバーの起動
    app := NewApp(&Config{
        Database:  db,
        JWTSecret: env.JWTSecret,
        TestMode:  true,
    })
    
    server := httptest.NewServer(app.Handler())
    env.Server = server
    env.ServerURL = server.URL
    env.cleanup = append(env.cleanup, server.Close)

    return env
}
```#### 認証フロ
ーのテスト

```yaml
# testdata/scenarios/auth/login.yml
desc: ユーザー認証フローのテスト
steps:
  # 管理者ログイン
  admin_login:
    req:
      api:///auth/login:
        post:
          body:
            application/json:
              email: "{{ vars.admin_email }}"
              password: "{{ vars.admin_password }}"
    test: |
      current.res.status == 200 &&
      current.res.body.token != null &&
      current.res.body.user.role == "admin" &&
      current.res.body.expires_in > 0

  # 一般ユーザーログイン
  user_login:
    req:
      api:///auth/login:
        post:
          body:
            application/json:
              email: "{{ vars.test_user_email }}"
              password: "{{ vars.test_user_password }}"
    test: |
      current.res.status == 200 &&
      current.res.body.token != null &&
      current.res.body.user.role == "user"

  # 無効な認証情報
  invalid_login:
    req:
      api:///auth/login:
        post:
          body:
            application/json:
              email: "invalid@example.com"
              password: "wrongpassword"
    test: |
      current.res.status == 401 &&
      current.res.body.error != null

  # トークンの検証
  verify_admin_token:
    req:
      api:///auth/verify:
        get:
          headers:
            Authorization: "Bearer {{ steps.admin_login.res.body.token }}"
    test: |
      current.res.status == 200 &&
      current.res.body.user.email == vars.admin_email

  # 期限切れトークンのシミュレーション
  expired_token_test:
    req:
      api:///auth/verify:
        get:
          headers:
            Authorization: "Bearer expired.jwt.token"
    test: current.res.status == 401
```

#### 商品管理のテスト

```yaml
# testdata/scenarios/products/crud.yml
desc: 商品CRUD操作のテスト
vars:
  test_product:
    name: "テスト商品"
    description: "これはテスト用の商品です"
    price: 1999
    category_id: 1
    stock: 100
    sku: "TEST-PRODUCT-001"

steps:
  # 管理者認証
  admin_auth:
    include:
      path: ../auth/login.yml

  # 商品作成（管理者権限必要）
  create_product:
    req:
      api:///products:
        post:
          headers:
            Authorization: "Bearer {{ steps.admin_auth.admin_login.res.body.token }}"
          body:
            application/json: "{{ vars.test_product }}"
    test: |
      current.res.status == 201 &&
      current.res.body.id > 0 &&
      current.res.body.name == vars.test_product.name &&
      current.res.body.sku == vars.test_product.sku

  # 商品一覧取得（認証不要）
  list_products:
    req:
      api:///products:
        get:
          query:
            page: 1
            limit: 10
    test: |
      current.res.status == 200 &&
      len(current.res.body.products) > 0 &&
      current.res.body.pagination.total > 0

  # 特定商品の取得
  get_product:
    req:
      api:///products/{{ steps.create_product.res.body.id }}:
        get:
    test: |
      current.res.status == 200 &&
      current.res.body.id == steps.create_product.res.body.id &&
      current.res.body.name == vars.test_product.name

  # 商品の更新
  update_product:
    req:
      api:///products/{{ steps.create_product.res.body.id }}:
        put:
          headers:
            Authorization: "Bearer {{ steps.admin_auth.admin_login.res.body.token }}"
          body:
            application/json:
              name: "更新されたテスト商品"
              price: 2499
              stock: 150
    test: |
      current.res.status == 200 &&
      current.res.body.name == "更新されたテスト商品" &&
      current.res.body.price == 2499

  # 在庫確認
  check_stock:
    req:
      api:///products/{{ steps.create_product.res.body.id }}/stock:
        get:
    test: |
      current.res.status == 200 &&
      current.res.body.stock == 150

  # 商品検索
  search_products:
    req:
      api:///products/search:
        get:
          query:
            q: "更新されたテスト"
            category: 1
    test: |
      current.res.status == 200 &&
      len(current.res.body.products) > 0 &&
      any(current.res.body.products, {.id == steps.create_product.res.body.id})

  # 商品削除
  delete_product:
    req:
      api:///products/{{ steps.create_product.res.body.id }}:
        delete:
          headers:
            Authorization: "Bearer {{ steps.admin_auth.admin_login.res.body.token }}"
    test: current.res.status == 204

  # 削除確認
  verify_deletion:
    req:
      api:///products/{{ steps.create_product.res.body.id }}:
        get:
    test: current.res.status == 404
```

### 2. マイクロサービスの統合テスト

複数のマイクロサービスが連携するシステムのテスト例です。

```go
func TestMicroservicesIntegration(t *testing.T) {
    // 複数のサービスを起動
    services := setupMicroservices(t)
    defer services.Cleanup()

    opts := []runn.Option{
        runn.T(t),
        runn.Runner("user_service", services.UserService.URL),
        runn.Runner("product_service", services.ProductService.URL),
        runn.Runner("order_service", services.OrderService.URL),
        runn.Runner("notification_service", services.NotificationService.URL),
        runn.DBRunner("user_db", services.UserDB),
        runn.DBRunner("product_db", services.ProductDB),
        runn.DBRunner("order_db", services.OrderDB),
    }

    o, err := runn.Load("testdata/microservices/**/*.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }

    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}
```

```yaml
# testdata/microservices/user_journey.yml
desc: マイクロサービス間の連携テスト
steps:
  # ユーザーサービス：ユーザー作成
  create_user:
    req:
      user_service:///users:
        post:
          body:
            application/json:
              name: "{{ faker.name() }}"
              email: "{{ faker.email() }}"
    test: current.res.status == 201

  # 商品サービス：商品情報取得
  get_products:
    req:
      product_service:///products:
        get:
          query:
            limit: 5
    test: |
      current.res.status == 200 &&
      len(current.res.body.products) > 0

  # 注文サービス：注文作成
  create_order:
    req:
      order_service:///orders:
        post:
          body:
            application/json:
              user_id: "{{ steps.create_user.res.body.id }}"
              items:
                - product_id: "{{ steps.get_products.res.body.products[0].id }}"
                  quantity: 1
    test: current.res.status == 201

  # 通知サービス：通知送信確認
  verify_notification:
    loop:
      count: 5
      until: len(current.res.body.notifications) > 0
      minInterval: 1
    req:
      notification_service:///notifications:
        get:
          query:
            user_id: "{{ steps.create_user.res.body.id }}"
            type: "order_created"
    test: |
      current.res.status == 200 &&
      len(current.res.body.notifications) > 0

  # データ整合性確認
  verify_data_consistency:
    db:
      user_db:///
        query: SELECT * FROM users WHERE id = $1
        params:
          - "{{ steps.create_user.res.body.id }}"
    test: len(current.rows) == 1

  verify_order_data:
    db:
      order_db:///
        query: SELECT * FROM orders WHERE user_id = $1
        params:
          - "{{ steps.create_user.res.body.id }}"
    test: |
      len(current.rows) == 1 &&
      current.rows[0].status == "created"
```

## ベストプラクティス

### 1. テストデータの管理

#### 固定データとランダムデータの使い分け

```yaml
# testdata/data_management.yml
desc: テストデータ管理のベストプラクティス
vars:
  # 固定データ：テストの再現性が重要な場合
  fixed_test_data:
    admin_user:
      email: "admin@example.com"
      password: "admin123"
    test_categories:
      - { id: 1, name: "Electronics" }
      - { id: 2, name: "Books" }
      - { id: 3, name: "Clothing" }
  
  # ランダムデータ：データの多様性が重要な場合
  random_test_data:
    users: |
      map(range(1, 6), {
        "name": faker.name(),
        "email": faker.email(),
        "age": faker.randomInt(18, 65),
        "department": faker.randomChoice(["IT", "Sales", "Marketing"])
      })

steps:
  # 固定データを使用したテスト
  test_with_fixed_data:
    req:
      api:///auth/login:
        post:
          body:
            application/json: "{{ vars.fixed_test_data.admin_user }}"
    test: current.res.status == 200

  # ランダムデータを使用したテスト
  test_with_random_data:
    loop:
      count: len(vars.random_test_data.users)
    req:
      api:///users:
        post:
          body:
            application/json: "{{ vars.random_test_data.users[i] }}"
    test: current.res.status == 201
```

### 2. エラーハンドリングとリトライ戦略

```yaml
# testdata/error_handling.yml
desc: エラーハンドリングのベストプラクティス
steps:
  # 基本的なリトライ
  robust_api_call:
    loop:
      count: 3
      until: current.res.status == 200
      minInterval: 1
      maxInterval: 5
    req:
      api:///unstable-endpoint:
        get:
    test: current.res.status == 200

  # 条件付きエラーハンドリング
  conditional_retry:
    loop:
      count: 5
      until: |
        current.res.status == 200 || 
        current.res.status == 404  # 404は正常として扱う
      minInterval: 2
    req:
      api:///resource/{{ vars.resource_id }}:
        get:
    test: current.res.status in [200, 404]

  # エラー情報の詳細記録
  detailed_error_logging:
    req:
      api:///complex-operation:
        post:
          body:
            application/json: "{{ vars.complex_data }}"
    test: true  # エラーでも続行
    dump:
      error_details: |
        current.res.status >= 400 ? {
          "status": current.res.status,
          "error_message": current.res.body.message ?? "Unknown error",
          "request_id": current.res.headers["X-Request-ID"],
          "timestamp": time("now"),
          "input_data_size": len(toJSON(vars.complex_data))
        } : null
```

### 3. 環境別設定の管理

```yaml
# testdata/environment_config.yml
desc: 環境別設定の管理
vars:
  # 環境変数による設定切り替え
  environment: "{{ env.TEST_ENV || 'development' }}"
  
  # 環境別設定
  config:
    development:
      api_url: "http://localhost:8080"
      timeout: 30
      retry_count: 3
      debug: true
    staging:
      api_url: "https://staging-api.example.com"
      timeout: 10
      retry_count: 2
      debug: false
    production:
      api_url: "https://api.example.com"
      timeout: 5
      retry_count: 1
      debug: false

  # 現在の環境設定を取得
  current_config: "{{ vars.config[vars.environment] }}"

runners:
  api: "{{ vars.current_config.api_url }}"

steps:
  environment_specific_test:
    req:
      api:///health:
        get:
          timeout: "{{ vars.current_config.timeout }}s"
    test: current.res.status == 200
    
    dump:
      environment_info:
        env: "{{ vars.environment }}"
        api_url: "{{ vars.current_config.api_url }}"
        debug_mode: "{{ vars.current_config.debug }}"
```

## デバッグ方法

### 1. 段階的なデバッグ

```yaml
# testdata/debugging.yml
desc: デバッグ技法の実践
steps:
  # ステップ1: 基本的な接続確認
  connectivity_check:
    req:
      api:///health:
        get:
    test: current.res.status == 200
    dump:
      health_status: current.res.body

  # ステップ2: 認証の確認
  auth_debug:
    req:
      api:///auth/login:
        post:
          body:
            application/json:
              email: "{{ vars.test_email }}"
              password: "{{ vars.test_password }}"
    test: true  # エラーでも続行してデバッグ
    dump:
      auth_request:
        url: current.req.url
        headers: current.req.headers
        body: current.req.body
      auth_response:
        status: current.res.status
        headers: current.res.headers
        body: current.res.body

  # ステップ3: 詳細なリクエスト/レスポンス情報
  detailed_debug:
    req:
      api:///complex-endpoint:
        post:
          body:
            application/json: "{{ vars.complex_payload }}"
    test: true
    dump:
      request_analysis:
        method: current.req.method
        url: current.req.url
        headers: current.req.headers
        body_size: len(toJSON(current.req.body))
        body_preview: |
          len(toJSON(current.req.body)) > 1000 ?
          "Large payload (" + string(len(toJSON(current.req.body))) + " bytes)" :
          current.req.body
      
      response_analysis:
        status: current.res.status
        headers: current.res.headers
        body_size: len(toJSON(current.res.body))
        response_time: current.res.response_time
        error_details: |
          current.res.status >= 400 ? {
            "error_code": current.res.body.code ?? "unknown",
            "error_message": current.res.body.message ?? "no message",
            "error_details": current.res.body.details ?? {}
          } : null
```

## パフォーマンスチューニング

### 1. 負荷テストの実装

```go
func TestAPIPerformance(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping performance test in short mode")
    }

    // パフォーマンステスト用の設定
    testConfig := &PerformanceTestConfig{
        ConcurrentUsers:    50,
        RequestsPerUser:    100,
        RampUpDuration:     30 * time.Second,
        TestDuration:       5 * time.Minute,
        AcceptableErrorRate: 0.01, // 1%
        RequiredRPS:        100,
    }

    db := setupTestDB(t)
    defer db.Close()

    // 大量のテストデータを準備
    preparePerformanceData(t, db, 100000)

    server := httptest.NewServer(NewApp(db).Handler())
    defer server.Close()

    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", server.URL),
        runn.DBRunner("db", db),
        runn.Var("concurrent_users", testConfig.ConcurrentUsers),
        runn.Var("requests_per_user", testConfig.RequestsPerUser),
        runn.Var("test_duration", testConfig.TestDuration.Seconds()),
    }

    start := time.Now()
    o, err := runn.Load("testdata/performance/**/*.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }

    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
    
    duration := time.Since(start)
    
    // パフォーマンス指標の評価
    evaluatePerformance(t, testConfig, duration)
}
```

```yaml
# testdata/performance/load_test.yml
desc: 負荷テストシナリオ
steps:
  # 並行ユーザーシミュレーション
  concurrent_load:
    loop:
      count: "{{ vars.concurrent_users }}"
    include:
      path: ./user_simulation.yml
      vars:
        user_id: "{{ i }}"
        requests_count: "{{ vars.requests_per_user }}"
```

## まとめ

この章では、runnを実際のプロジェクトで活用するための実践的な知識を学びました：

1. **よくあるユースケース**: RESTful API、マイクロサービス統合テスト
2. **ベストプラクティス**: テストデータ管理、エラーハンドリング、環境別設定
3. **デバッグ方法**: 段階的デバッグ、詳細ログ出力
4. **パフォーマンスチューニング**: 負荷テスト、メモリ監視

これらの実践的な技法を組み合わせることで、本格的なプロダクション環境で使用できる高品質なテストスイートを構築できます。

次章では、runnの詳細な仕様とFAQをまとめたリファレンスを提供します。

[第9章：リファレンスへ →](chapter09.md)