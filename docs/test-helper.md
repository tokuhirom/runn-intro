# 第5章：テストヘルパーとしての利用

runnはGoテストヘルパーとして使用でき、`go test`と統合してシナリオベースのテストを実行できます。

## 基本的な統合方法

### プロジェクト構造

```
myproject/
├── main.go
├── handler.go
├── go.mod
└── testdata/
    └── scenarios/
        ├── user_test.yml
        └── auth_test.yml
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
        runn.T(t),
        runn.Runner("api", server.URL),
        runn.DBRunner("db", db),
    }
    
    // シナリオの実行
    o, err := runn.Load("testdata/scenarios/user_test.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}
```

### YAMLシナリオ例

```yaml
{{ includex("examples/test-helper/user_test.yml") }}
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
    
    app := NewApp(db)
    server := httptest.NewServer(app.Handler())
    defer server.Close()
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", server.URL),
        runn.DBRunner("db", db),
    }
    
    o, err := runn.Load("testdata/user_crud.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
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
                },
            })
        default:
            w.WriteHeader(404)
        }
    }))
    defer mockServer.Close()
    
    db := setupTestDB(t)
    defer db.Close()
    
    app := NewApp(db)
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

### テストデータの準備

```go
func TestComplexScenario(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()
    
    // テストデータの準備
    testData := prepareTestData(t, db)
    
    server := httptest.NewServer(NewApp(db).Handler())
    defer server.Close()
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", server.URL),
        runn.DBRunner("db", db),
        runn.Var("admin_user_id", testData.AdminUser.ID),
        runn.Var("regular_user_id", testData.RegularUser.ID),
    }
    
    o, err := runn.Load("testdata/complex_scenario.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}
```

## 実践的なテストパターン

### 認証フローのテスト

```yaml
{{ includex("examples/test-helper/auth_flow.yml") }}
```

### E2Eワークフローテスト

```yaml
{{ includex("examples/test-helper/e2e_workflow.yml") }}
```

## パフォーマンステスト

```go
func TestAPIPerformance(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping performance test in short mode")
    }
    
    db := setupTestDB(t)
    defer db.Close()
    
    // パフォーマンステスト用のデータを準備
    preparePerformanceTestData(t, db, 10000)
    
    app := NewApp(db)
    server := httptest.NewServer(app.Handler())
    defer server.Close()
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("api", server.URL),
        runn.DBRunner("db", db),
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
}
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
        runn.Debug(true),
        runn.Verbose(true),
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

## CI/CDとの統合

### GitHub Actions

```yaml
{{ includex("examples/test-helper/github_actions.yml") }}
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