# 第8章：実践編 - 現場で使える最強テクニック！

**あなたの番が来た！** これまで学んできた知識を**現実のプロジェクトで解放**しよう！この章では、**実際の現場で直ぐに使える**テクニックを伝授する。よくあるユースケースから、**プロのベストプラクティス**、**最強のデバッグ法**、**爆速パフォーマンスチューニング**まで、**すべてを教える**！

## 💼 よくあるユースケース - 現場で即戦力のシナリオ！

### 1. 🌐 RESTful APIの包括的テスト - ECサイトを完全制覇！

**これがプロのテストだ！** 実際のECサイトAPIを例に、**企業レベルの完璧なテストスイート**を構築しよう！

#### 📁 プロジェクト構造 - 理想的なディレクトリ設計

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

#### 🚀 メインテストファイル - これがプロのコード！

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

#### 🔐 認証フローのテスト - セキュリティを完璧に検証！

```yaml
{{ includex("examples/chapter08/auth/login.yml") }}
```

#### 📦 商品管理のテスト - CRUD操作を完全網羅！

```yaml
{{ includex("examples/chapter08/products/crud.yml") }}
```

### 2. 🌍 マイクロサービスの統合テスト - 複雑なシステムを完全支配！

**マイクロサービス時代の最強テスト！** 複数のサービスが**美しく連携**するシステムを完璧にテストしよう！

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

## 🏆 ベストプラクティス - プロの流儀を伝授！

### 1. 📊 テストデータの管理 - データを完璧にコントロール！

#### 🎲 固定データとランダムデータの使い分け - プロの技！

```yaml
{{ includex("examples/chapter08/data_management.yml") }}
```

### 2. 🛡️ エラーハンドリングとリトライ戦略 - 失敗を成功に変えろ！

```yaml
{{ includex("examples/chapter08/error_handling.yml") }}
```

### 3. 🌍 環境別設定の管理 - どんな環境でも完璧に動く！

```yaml
{{ includex("examples/chapter08/environment_config.yml") }}
```

## 🔍 デバッグ方法 - 問題を瞬時に特定する魔法！

### 1. 📈 段階的なデバッグ - ステップ・バイ・ステップで確実に！

```yaml
{{ includex("examples/chapter08/debugging.yml") }}
```

## ⚡ パフォーマンスチューニング - 速度の限界を突破！

### 1. 🚀 負荷テストの実装 - システムの真の力を測れ！

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

## 🎆 まとめ - 実践テクニックのマスター誕生！

**やったぞ！** あなたは今、**現場で即戦力のテクニック**をすべて手に入れた！

### 🏆 この章でマスターした4つの奥義：

1. **💼 よくあるユースケース**: RESTful API、マイクロサービスを**完全制覇**！
2. **🏆 ベストプラクティス**: データ管理、エラーハンドリング、環境設定を**極めた**！
3. **🔍 デバッグ方法**: 問題を**瞬時に特定**する技を習得！
4. **⚡ パフォーマンスチューニング**: **爆速テスト**の秘訣を伝授！

これらの**実践的な武器**を組み合わせれば、どんなに複雑なプロダクション環境でも**完璧なテストスイート**を構築できる。あなたはもう、**ただのテスターではない**。**テストのアーキテクト**だ！
