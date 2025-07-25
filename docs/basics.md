# 第1章：基礎編

## runnとは何か

[runn](https://github.com/k1LoW/runn)（「Run N」/rʌ́n én/）は、[k1LoW](https://github.com/k1LoW)氏が開発したシナリオベースのテスト・自動化ツールです。

### 特徴

#### マルチプロトコル対応
HTTP、gRPC、データベース、CDP（Chrome DevTools Protocol）、SSHを同一のYAML形式で記述できます。

```yaml
{{ includex("examples/basics/intro/intro-multi-protocol.concept.yml") }}
```

#### シングルバイナリ
単体で実行可能。ダウンロードしてすぐ使えます。

#### 強力な式評価エンジン
前のステップの結果を次のステップで利用できます。

```yaml
{{ includex("examples/basics/intro/intro-step-chaining.yml") }}
```

#### Go言語統合
`go test`にシームレスに統合可能です。

### 用途

- APIのE2Eテスト
- CI/CDでの自動テスト
- 運用タスクの自動化
- APIの動作確認

## インストール

### Homebrew
```bash
brew install k1LoW/tap/runn
```

### Go install
```bash
go install github.com/k1LoW/runn/cmd/runn@latest
```

### 直接ダウンロード
[GitHub Releases](https://github.com/k1LoW/runn/releases)から環境に合ったバイナリをダウンロード。

### Docker
```bash
docker container run -it --rm --name runn -v $PWD:/books ghcr.io/k1low/runn:latest list /books/*.yml
```

### 確認
```bash
runn --version
```

## 基本的な使い方

### 主要コマンド

#### `runn run` - シナリオ実行
```bash
runn run scenario.yml
runn run scenarios/**/*.yml
```

#### `runn list` - シナリオ一覧
```bash
runn list scenarios/
```

#### `runn new` - シナリオ生成
```bash
# curlコマンドから生成
runn new --and-run --out first.yml -- curl https://httpbin.org/get

# アクセスログから生成
cat access.log | runn new --out generated.yml
```

### 便利なオプション

- `--verbose` - 詳細ログ表示
- `--label` - ラベルでフィルタリング
- `--concurrent` - 並列実行数指定
- `--fail-fast` - 最初のエラーで停止

## はじめてのシナリオ作成

### テスト環境準備
```bash
docker run -p 8080:8080 mccutchen/go-httpbin
```

### 基本的なGETリクエスト

`examples/basics/first-scenario.yml`:

```yaml
{{ includex("examples/basics/first-scenario.yml") }}
```

実行:
```bash
runn run examples/basics/first-scenario.yml --verbose
```

結果:
```
{{ includex("examples/basics/first-scenario.out") }}
```

### JSONレスポンスの検証

```yaml
{{ includex("examples/basics/json-validation.yml") }}
```

### 変数の使用

```yaml
{{ includex("examples/basics/with-variables.yml") }}
```

### ステップ間の連携

```yaml
{{ includex("examples/basics/multi-step.yml") }}
```

`steps.login.res.body.json.username`で前のステップの結果を参照できます。

## CLIツール vs Goテストヘルパー

### CLIツールとして

適した用途:
- 手動でのAPI動作確認
- CI/CDパイプラインでの自動テスト
- 外部APIの監視
- デバッグ作業

### Goテストヘルパーとして

適した用途:
- Goアプリケーションのテスト統合
- テストDBのセットアップ/クリーンアップ
- モックサーバーとの連携
- 複雑なテストデータの準備

### Goテストヘルパーの実装例

`main_test.go`:
```go
{{ includex("examples/basics/go-test/main_test.go") }}
```

テスト対象のAPIサーバー（`main.go`）:
```go
{{ includex("examples/basics/go-test/main.go") }}
```

テストシナリオ（`user-api-test.yml`）:
```yaml
{{ includex("examples/basics/user-api-test.yml") }}
```

このように、SQLiteを使った本格的なREST APIのテストが、わずかなコードで実現できます。