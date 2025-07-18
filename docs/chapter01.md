# 第1章：基礎編

## runnとは何か

> 「APIテストを書くのが面倒だ」「複雑なシナリオテストをもっとシンプルに書きたい」「テストコードのメンテナンスに疲れた」
> 
> そんなあなたに、**runn**という革命的なツールを紹介します。

[runn](https://github.com/k1LoW/runn)（「Run N」と読み、/rʌ́n én/と発音）は、[k1LoW](https://github.com/k1LoW)氏によって開発された、**シナリオベースのテスト・自動化ツール**です。2022年の登場以来、そのシンプルさと強力さで多くの開発者を魅了し続けています。

### なぜrunnが必要なのか？

従来のAPIテストツールやE2Eテストフレームワークには、こんな課題がありました：

- **学習コストが高い**: 独自のDSLや複雑なAPIを覚える必要がある
- **コードが冗長**: シンプルなテストでも大量のボイラープレートコードが必要
- **メンテナンスが大変**: テストコードの修正に時間がかかる
- **プロトコルごとに別ツール**: HTTP、gRPC、DBテストで異なるツールを使い分ける必要がある

runnは、これらの課題を**YAMLベースの宣言的な記述**で解決します：

```yaml
{{ includex("examples/chapter01/intro-simple-api-test.yml") }}
```

### runnの3つの顔

runnは単なるAPIテストツールではありません。以下の3つの側面を持つ、**マルチパーパスツール**です：

#### 1. 🎯 シナリオベースのテストツール
YAMLでテストシナリオを記述し、CLIから実行。**学習コスト最小**で始められます。

#### 2. 🔧 Go言語のテストヘルパー
`go test`にシームレスに統合。**既存のGoプロジェクトにそのまま導入**できます。

#### 3. 🤖 汎用的な自動化ツール
CI/CDでの活用、定期実行タスク、運用自動化など、**テスト以外でも大活躍**。

### runnの革新的な特徴

#### 🌐 **マルチプロトコル対応**
HTTP、gRPC、データベース、ブラウザ操作、SSHまで、**すべて同じYAML形式**で記述できます。もう複数のツールを使い分ける必要はありません。

```yaml
{{ includex("examples/chapter01/intro-multi-protocol.yml") }}
```

#### 📦 **シングルバイナリ**
依存関係なし、環境構築不要。**ダウンロードしてすぐ使える**シンプルさ。CI環境でも `curl` 一発でインストール完了。

#### 🔗 **強力な式評価エンジン**
前のステップの結果を次のステップで利用する**ステップ間連携**が自由自在。複雑なシナリオも直感的に記述できます。

```yaml
{{ includex("examples/chapter01/intro-step-chaining.yml") }}
```

#### 🚀 **Goテスト統合**

既存のGoプロジェクトに**1ファイル追加するだけ**で導入可能。テストデータベースのセットアップや並列実行も自由自在。

### 誰のためのツール？

- **APIを開発している人**: シンプルで保守しやすいE2Eテストを書きたい
- **QAエンジニア**: プログラミング知識なしでもテストシナリオを作成したい
- **SRE/インフラエンジニア**: 運用タスクを自動化したい
- **Goプログラマー**: 既存のテストフレームワークと統合したい

### なぜ今、runnなのか？

マイクロサービス化が進む現代、複数のサービス間の連携テストはますます重要になっています。しかし、既存のツールでは複雑なシナリオの記述が困難でした。

runnは、**「シンプルに書けて、パワフルに実行できる」**という理想を実現したツールです。YAMLという親しみやすい形式で、誰でも読み書きできるテストシナリオを作成できます。

さあ、次のセクションでrunnをインストールして、この革命的なツールの威力を体感してみましょう！

## インストール方法

runnのインストールには複数の方法があります。環境に応じて選択してください。

### 1. Homebrewを使用（macOS/Linux）

```bash
brew install k1LoW/tap/runn
```

### 2. Go installを使用

```bash
go install github.com/k1LoW/runn/cmd/runn@latest
```

### 3. バイナリを直接ダウンロード

[GitHub Releases](https://github.com/k1LoW/runn/releases)から、お使いのOS・アーキテクチャに対応したバイナリをダウンロードすることもできます。

### 4. Dockerを使用

```bash
docker container run -it --rm --name runn -v $PWD:/books ghcr.io/k1low/runn:latest list /books/*.yml
```

### インストールの確認

```bash
runn --version
```

## 基本的な使い方（CLIコマンド）

### 主要なコマンド

```bash
# シナリオを実行
runn run scenario.yml

# 複数のシナリオを実行
runn run scenarios/**/*.yml

# シナリオ一覧を表示
runn list scenarios/

# curlコマンドからシナリオを生成
runn new --and-run --out first.yml -- curl https://httpbin.org/get

# アクセスログからシナリオを生成
cat access.log | runn new --out generated.yml
```

### 便利なオプション

```bash
# 詳細なログを表示
runn run scenario.yml --verbose

# 特定のラベルのシナリオのみ実行
runn run scenarios/**/*.yml --label api --label critical

# 並列実行数を指定
runn run scenarios/**/*.yml --concurrent 5

# 失敗時に即座に停止
runn run scenarios/**/*.yml --fail-fast
```

## はじめてのシナリオ作成

実際にシナリオを作る前に、テスト対象になるサーバーを用意します。
`go-httpbin` を使うとさまざまなリクエストパターンがためせて便利です。

これ以後のサンプルコードでは、`go-httpbin` が起動していることを前提として進めていきます。

```bash
docker run -p 8080:8080 mccutchen/go-httpbin
```

### 1. 最もシンプルなHTTPリクエスト

`examples/chapter01/first-scenario.yml`を作成：

この例では、`/get` にリクエストし、response code が 200 であることを確認しています。

```yaml
{{ includex("examples/chapter01/first-scenario.yml") }}
```

実行：

```bash
runn run examples/chapter01/first-scenario.yml --verbose
```

実行結果:

```
{{ includex("examples/chapter01/first-scenario.out") }}
```

### 2. レスポンスの検証を追加

```yaml
{{ includex("examples/chapter01/json-validation.yml") }}
```

### 3. 変数を使用したシナリオ

```yaml
{{ includex("examples/chapter01/with-variables.yml") }}
```

### 4. 複数ステップのシナリオ

```yaml
{{ includex("examples/chapter01/multi-step.yml") }}
```

## CLIとGoテストヘルパーの使い分け

### CLIツールとして使うべき場面

- 手動でのAPI動作確認
- CI/CDパイプラインでの簡易的なE2Eテスト
- 外部APIの監視・ヘルスチェック
- 開発中の動作確認

### Goテストヘルパーとして使うべき場面

- **Goで書かれたアプリケーションのAPIテスト**（推奨）
- テストDBの準備や初期化が必要な場合
- モックサーバーとの連携が必要な場合
- 複雑なテストデータの準備が必要な場合
- 既存のGoテストフレームワークと統合したい場合

実際のプロジェクトでは、**Goテストヘルパーとしての利用が圧倒的に強力**です。第7章で詳しく解説しますが、ここで簡単な例を示します：

```go
{{ includex("examples/chapter01/go-test/main_test.go") }}
```

テスト対象のサーバーはこのような実装になっています。

```go
{{ includex("examples/chapter01/go-test/main.go") }}
```

このテストで使用するYAMLシナリオはこのような感じです：

```yaml
{{ includex("examples/chapter01/go-test/user-api-test.yml") }}
```

## まとめ

この章では、runnの基本的な概念とインストール方法、簡単なシナリオの作成方法を学びました。重要なポイント：

1. runnはCLIツールとGoテストヘルパーの2つの顔を持つ
2. YAMLで宣言的にテストシナリオを記述できる
3. 前のステップの結果を次のステップで利用できる
4. Goテストヘルパーとして使用すると、より強力で柔軟なテストが可能

次章では、より詳細なシナリオの記述方法について学んでいきます。

[第2章：シナリオ記述編へ →](chapter02.md)
