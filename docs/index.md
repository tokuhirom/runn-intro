# runn入門 - シナリオベースのE2Eテストツール

## runnの2つの顔：CLIツールとGoテストヘルパー

runnには大きく分けて2つの使い方があります：

1. **CLIツールとしての利用** - YAMLファイルに記述したシナリオを`runn`コマンドで実行
2. **Goテストヘルパーとしての利用** - Go言語のテストコードに組み込んで実行

実は、**著者はGoテストヘルパーとしての利用を特に重視しています**。なぜなら、この方法により以下のような強力な機能が実現できるからです：

### Goテストヘルパーとしての利点

```go
func TestAPIWithRunn(t *testing.T) {
    // 1. テスト用のデータベースを準備
    db := setupTestDB(t)
    defer db.Close()
    
    // 2. テストデータを投入
    seedTestData(t, db)
    
    // 3. アプリケーションサーバーを起動
    app := NewApp(db)
    ts := httptest.NewServer(app.Handler())
    defer ts.Close()
    
    // 4. runnでシナリオを実行
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("req", ts.URL),        // テストサーバーのURLを動的に設定
        runn.DBRunner("db", db),            // テスト用DBを直接渡せる
        runn.Var("testuser", "alice@example.com"),
    }
    
    o, err := runn.Load("testdata/scenarios/**/*.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(context.Background()); err != nil {
        t.Fatal(err)
    }
}
```

この方法により：
- テスト環境の初期化・後片付けをGoで完全制御
- テストデータの動的な準備
- 環境に依存しない安定したテスト実行
- IDEのデバッガでブレークポイントを設定可能
- `go test`の豊富なオプションを活用

## 本書の構成

本書では、まずCLIツールとしての基本的な使い方を学び、その後Goテストヘルパーとしての活用方法を深く掘り下げていきます。

### 目次

1. [第1章：基礎編](chapter01.md) - runnの概要とインストール
2. [第2章：シナリオ記述編](chapter02.md) - YAMLでのシナリオ記述方法
3. [第3章：Expression文法編](chapter03.md) - 式評価エンジンの詳細
4. [第4章：ビルトイン関数編](chapter04.md) - 便利な組み込み関数群
5. [第5章：ランナー詳細編](chapter05.md) - 各種プロトコルのサポート
6. [第6章：高度な機能編](chapter06.md) - ループ、条件分岐、並行実行
7. [第7章：Goテストヘルパー編](chapter07.md) - **本書の核心**
8. [第8章：実践編](chapter08.md) - 実際のプロジェクトでの活用例
9. [第9章：リファレンス](chapter09.md) - 詳細な仕様とFAQ

## なぜrunnなのか？

従来のAPIテストツールと比較して、runnは以下の点で優れています：

### 1. 宣言的なシナリオ記述
```yaml
steps:
  - req:
      /users:
        post:
          body:
            username: alice
    test: current.res.status == 201
```

### 2. 複数プロトコルの統一的なサポート
- HTTP/HTTPS
- gRPC
- データベース
- ブラウザ操作（CDP）
- SSH

### 3. 強力な式評価エンジン
前のステップの結果を次のステップで利用できる柔軟な変数参照システム。

### 4. Goエコシステムとの完璧な統合
`go test`の一部として実行できるため、既存のCI/CDパイプラインにシームレスに統合。

## 始めましょう

それでは、[第1章：基礎編](chapter01.md)から始めましょう。runnの世界へようこそ！