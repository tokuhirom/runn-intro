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
```

#### 認証フローのテスト

```yaml
{{ includex("examples/chapter08/auth/login.yml") }}
```

#### 商品管理のテスト

```yaml
{{ includex("examples/chapter08/products/crud.yml") }}
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
{{ includex("examples/chapter08/microservices/user_journey.yml") }}
```

## ベストプラクティス

### 1. テストデータの管理

#### 固定データとランダムデータの使い分け

```yaml
{{ includex("examples/chapter08/data_management.yml") }}
```

### 2. エラーハンドリングとリトライ戦略

```yaml
{{ includex("examples/chapter08/error_handling.yml") }}
```

### 3. 環境別設定の管理

```yaml
{{ includex("examples/chapter08/environment_config.yml") }}
```

## デバッグ方法

### 1. 段階的なデバッグ

```yaml
{{ includex("examples/chapter08/debugging.yml") }}
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
{{ includex("examples/chapter08/performance/load_test.yml") }}
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