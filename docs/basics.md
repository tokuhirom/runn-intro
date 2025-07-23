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
{{ includex("examples/basics/intro/intro-simple-api-test.yml") }}
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
{{ includex("examples/basics/intro/intro-multi-protocol.concept.yml") }}
```

#### 📦 **シングルバイナリ**
依存関係なし、環境構築不要。**ダウンロードしてすぐ使える**シンプルさ。CI環境でも `curl` 一発でインストール完了。

#### 🔗 **強力な式評価エンジン**
前のステップの結果を次のステップで利用する**ステップ間連携**が自由自在。複雑なシナリオも直感的に記述できます。

```yaml
{{ includex("examples/basics/intro/intro-step-chaining.yml") }}
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

## 今すぐrunnを始めよう！

> 「よし、runnを使ってみたい！」
> 
> その意欲に応えるため、**最速でrunnを始められる方法**を紹介します。

### 爆速インストールガイド

#### **Homebrew** - macOS/Linuxユーザーの最速ルート

たった1行で完了！コーヒーを淹れる時間すら必要ありません：

```bash
brew install k1LoW/tap/runn
```

#### **Go install** - Gopherたちの選択

Go環境がある？それなら話は早い：

```bash
go install github.com/k1LoW/runn/cmd/runn@latest
```

**Pro tip**: `go install`なら最新の開発版も簡単に試せます！

#### **直接ダウンロード** - シンプル・イズ・ベスト

[GitHub Releases](https://github.com/k1LoW/runn/releases)から、あなたの環境に合ったバイナリを選んでダウンロード。**解凍して配置するだけ**で準備完了！

- Windows? ✅ 対応
- Mac (Intel/Apple Silicon)? ✅ 対応  
- Linux (各種アーキテクチャ)? ✅ 対応

#### 🐳 **Docker** - 環境を汚したくない慎重派のあなたへ

ローカル環境は綺麗に保ちたい？Dockerで隔離実行：

```bash
docker container run -it --rm --name runn -v $PWD:/books ghcr.io/k1low/runn:latest list /books/*.yml
```

### インストール成功の瞬間

ドキドキの確認タイム：

```bash
runn --version
```

バージョン番号が表示されたら、**おめでとうございます！**  
あなたは今、runnの世界への扉を開きました。

## runnマスターへの第一歩

> インストール完了？素晴らしい！
> 
> では、**runnの基本コマンド**をマスターして、テスト自動化の達人への道を歩み始めましょう。

### 必須コマンド - これだけ覚えれば今すぐ使える！

#### **`runn run`** - シナリオを実行する魔法の呪文

一番使うコマンドがこれ！

```bash
# 単一シナリオの実行
runn run scenario.yml

# ワイルドカードで複数実行 - まとめてテスト！
runn run scenarios/**/*.yml
```

#### **`runn list`** - シナリオの棚卸し

どんなテストがあるか一目瞭然：

```bash
runn list scenarios/
```

#### **`runn new`** - シナリオ自動生成の魔術

**これがrunnの隠れた最強機能！** 既存のcurlコマンドやアクセスログから、自動でシナリオを生成：

```bash
# curlコマンドを即座にシナリオ化！
runn new --and-run --out first.yml -- curl https://httpbin.org/get

# アクセスログから一括生成 - 過去の通信を再現！
cat access.log | runn new --out generated.yml
```

**驚きの事実**: 手動でYAMLを書く前に、まず`runn new`を試してみて！

### パワーユーザー向けオプション

#### **`--verbose`** - 詳細ログで問題を即座に発見

デバッグの強い味方：

```bash
runn run scenario.yml --verbose
```

何が起きているか、すべてが見える！

#### **`--label`** - スマートなシナリオ管理

大規模プロジェクトでの必須テクニック：

```bash
# 重要なAPIテストだけを実行
runn run scenarios/**/*.yml --label api --label critical
```

#### **`--concurrent`** - 並列実行で爆速テスト

時間は金なり！並列度を上げて高速化：

```bash
runn run scenarios/**/*.yml --concurrent 5
```

**注意**: サーバーに優しく。並列度は適切に！

#### **`--fail-fast`** - 失敗したら即停止

CI/CDでの時間節約テクニック：

```bash
runn run scenarios/**/*.yml --fail-fast
```

最初のエラーで停止。無駄な待ち時間とはお別れ！

## あなたの最初のrunnシナリオ！

> 理論は十分。実践の時間です！
> 
> **5分後には、あなたも立派なrunnマスター**になっているでしょう。

### テスト環境をサクッと準備

まずはテスト用のAPIサーバーを立ち上げましょう。**go-httpbin**という便利なツールを使います：

```bash
docker run -p 8080:8080 mccutchen/go-httpbin
```

**起動した？** よし、これであなた専用のテスト環境の完成です！

### Scenario 1: Hello, runn! - 記念すべき第一歩

`examples/basics/first-scenario.yml`として保存：

```yaml
{{ includex("examples/basics/first-scenario.yml") }}
```

**たったこれだけ！** でも、これが立派なE2Eテストなんです。

#### 実行してみよう！

ドキドキの瞬間：

```bash
runn run examples/basics/first-scenario.yml --verbose
```

#### 実行結果

```
{{ includex("examples/basics/first-scenario.out") }}
```

**見ました？** `ok`の文字が！これがあなたの**初めてのテスト成功**です！

### Scenario 2: JSONレスポンスを賢く検証

APIテストの真骨頂、**レスポンスの中身まで検証**してみましょう：

```yaml
{{ includex("examples/basics/json-validation.yml") }}
```

**ポイント**: `current.res.body`で自由自在にJSONの中身をチェック！

### Scenario 3: 変数でDRYに

同じ値を何度も書くのは面倒？**変数を使ってスマートに**：

```yaml
{{ includex("examples/basics/with-variables.yml") }}
```

**魔法のような`{{ vars }}`記法**で、メンテナンスが劇的に楽に！

### Scenario 4: ステップ連携の威力

これぞrunnの真骨頂！**前のステップの結果を次で使う**：

```yaml
{{ includex("examples/basics/multi-step.yml") }}
```

**驚きポイント**: `steps.login.res.body.json.username`で前のレスポンスを参照！
実際のログインフローをそのままテストできます。

## runnの2つの顔 - あなたはどっち派？

> CLIツール？Goテストヘルパー？
> 
> **使い分けをマスターすれば、runnの真の力を引き出せます！**

### CLIツール派 - シンプル＆クイック

こんな時はCLIで決まり！

#### **手動でサクッとAPI確認**
開発中の動作確認に最適。ターミナルから即実行！

#### **CI/CDパイプラインで活躍**
GitHub ActionsやGitLab CIに組み込んで自動テスト。設定も簡単！

#### **外部APIの監視**
定期的にヘルスチェック。cronと組み合わせて24時間監視体制！

#### **開発中のデバッグ**
`--verbose`オプションで詳細ログ。問題箇所を即座に特定！

### Goテストヘルパー派 - パワフル＆フレキシブル

**プロの選択はこっち！** 本格的なプロジェクトでの威力は絶大：

#### **Goアプリケーションとの完璧な統合**
`go test`コマンドでそのまま実行。既存のテストフローに自然に組み込める！

#### **テストDBの自在な制御**
```go
db := setupTestDB(t)  // テスト用DBを準備
defer cleanupDB(db)   // 自動でクリーンアップ
```

#### **モックサーバーとの連携**
```go
mockServer := httptest.NewServer(mockHandler)
// runnでモックサーバーに対してテスト実行
```

#### **複雑なテストデータの準備**
Goコードでデータを生成してから、runnでシナリオテスト。最強の組み合わせ！

### 実践例：Goテストヘルパーの威力を体感

**これが現場で使われている本物のコード**です：

```go
{{ includex("examples/basics/go-test/main_test.go") }}
```

#### テスト対象サーバーの実装

まずは**実際に動くAPIサーバー**を見てみましょう：

```go
{{ includex("examples/basics/go-test/main.go") }}
```

**見どころ**: SQLiteのインメモリDBを使った本格的なREST API！

#### 魔法のYAMLシナリオ

そして、このサーバーをテストする**美しいシナリオ**：

```yaml
{{ includex("examples/basics/user-api-test.yml") }}
```

**注目ポイント**:
- ユーザー作成 → 取得 → エラーケースまで**完璧にカバー**
- `steps.create_user.res.body.id`で**動的にIDを参照**
- たった36行で**包括的なAPIテスト**が完成！

