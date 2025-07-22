# 第5章：ランナー詳細編 - 6大プロトコルを完全制覇！

**ついに来た、runnの真骨頂！** 他のテストツールが**一つのプロトコルしか扱えない**中、runnは**6つのプロトコルを統一的に操る**！この章では、各ランナーの**秘伝の技**を伝授しよう！

## 🎆 6大ランナー軍団 - あなたの最強の武器庫！

runnは**6つの強力なランナー**を搭載！それぞれが**特定のプロトコルのスペシャリスト**だ！

| ランナー | プロトコル | 用途 |
|----------|------------|------|
| **🌐 HTTP** | HTTP/HTTPS | REST API、GraphQL、Webhookを**完全制圧**！ |
| **🔗 gRPC** | gRPC | マイクロサービス間通信を**爆速テスト**！ |
| **🗄️ DB** | SQL | データベースを**自在に操作**！ |
| **🌐 CDP** | Chrome DevTools Protocol | ブラウザを**完全支配**！ |
| **💻 SSH** | SSH | リモートサーバーを**思いのままに**！ |
| **⚙️ Exec** | プロセス実行 | ローカルコマンドを**瞬時に実行**！ |

## 🌐 HTTPランナー - Web APIテストの王者！

### 🚀 基本的な設定 - まずはここから始めよう！

```yaml
{{ includex("examples/runners/http_basic_setup.yml") }}
```

### 🎨 リクエストメソッドとパラメータ - あらゆるHTTPメソッドを制覇！

```yaml
{{ includex("examples/runners/http_request_methods.yml") }}
```

### 📦 様々なボディ形式 - どんなデータ形式もお手のもの！

```yaml
{{ includex("examples/runners/http_body_formats.yml") }}
```

## 🔗 gRPCランナー - マイクロサービスの強い味方！

TODO: grpc の例を追加する

## 🗄️ データベースランナー - SQLの魔術師になれ！

### 🌍 対応データベース - あらゆるDBを制覇！

```yaml
{{ includex("examples/runners/db_connections.yml") }}
```

### 📝 基本的なクエリ操作 - SQLを思いのままに！

<!-- TODO: INSERT のあと RETRUNING 使えてない: https://github.com/k1LoW/runn/issues/1276 -->

```yaml
{{ includex("examples/runners/db_basic_queries.yml") }}
```

## 🌐 CDPランナー（ブラウザ自動化） - ブラウザを完全支配！

### 主なCDP actions一覧

| アクション名      | 概要                                   | 記述例（YAML）           |
|------------------|----------------------------------------|-------------------------|
| attributes       | 要素の属性取得                         | actions: - attributes: "h1" |
| click            | 要素をクリック                         | actions: - click: "nav > a" |
| doubleClick      | 要素をダブルクリック                   | actions: - doubleClick: "nav > li" |
| evaluate         | JS式の評価                             | actions: - evaluate: "document.querySelector('h1').textContent = 'hello'" |
| fullHTML         | ページ全体のHTML取得                   | actions: - fullHTML     |
| innerHTML        | 要素のinnerHTML取得                    | actions: - innerHTML: "h1" |
| localStorage     | localStorage取得                       | actions: - localStorage: "https://github.com" |
| location         | 現在のURL取得                          | actions: - location     |
| navigate         | 指定URLへ遷移                          | actions: - navigate: "https://pkg.go.dev/time" |
| outerHTML        | 要素のouterHTML取得                    | actions: - outerHTML: "h1" |
| screenshot       | スクリーンショット取得                 | actions: - screenshot   |
| scroll           | 要素までスクロール                     | actions: - scroll: "body > footer" |
| sendKeys         | 要素にキー入力                         | actions: - sendKeys: {sel: "input[name=username]", value: "xxx"} |
| sessionStorage   | sessionStorage取得                     | actions: - sessionStorage: "https://github.com" |
| setUploadFile    | ファイルアップロード                   | actions: - setUploadFile: {sel: "input[name=avator]", path: "/path/to/image.png"} |
| setUserAgent     | User-Agent設定                         | actions: - setUserAgent: "Mozilla/5.0 ..." |
| submit           | フォーム送信                           | actions: - submit: "form.login" |
| tabTo            | タブ切り替え                           | actions: - tabTo: "https://pkg.go.dev/time" |
| text             | 要素のテキスト取得                     | actions: - text: "h1"   |
| textContent      | 要素のtextContent取得                  | actions: - textContent: "h1" |
| title            | ページタイトル取得                     | actions: - title        |
| value            | 要素のvalue取得                        | actions: - value: "input[name=address]" |
| wait             | 指定時間待機                           | actions: - wait: "10sec"|
| waitReady        | 要素の準備完了まで待機                 | actions: - waitReady: "body > footer" |
| waitVisible      | 要素の表示まで待機                     | actions: - waitVisible: "body > footer" |

※詳細・最新情報は[公式README](https://github.com/k1LoW/runn?tab=readme-ov-file#functions-for-action-to-control-browser)をご参照ください。

### 🎮 基本的な使い方

```yaml
{{ includex("examples/runners/cdp_basic.yml") }}
```

## 💻 SSHランナー - リモートサーバーの絶対的支配者！

### 🔑 基本的な設定 - サーバーへのセキュアアクセス！

```yaml
{{ includex("examples/runners/ssh_basic.yml") }}
```

### 📏 サーバー監視とヘルスチェック - 24時間365日の番人！

```yaml
{{ includex("examples/runners/ssh_health_check.yml") }}
```

## ⚙️ Execランナー（ローカルコマンド実行） - シェルコマンドの魔術師！

### 🚀 基本的な使用方法 - コマンドを瞬時に実行！

```yaml
{{ includex("examples/runners/exec_basic.yml") }}
```

### 📁 ファイル操作とテスト - ローカルファイルを完全管理！

```yaml
{{ includex("examples/runners/exec_file_operations.yml") }}
```

## 🎆 ランナーの組み合わせ - 最強のコンボ技！

### 🌈 マルチプロトコルテスト - 複数プロトコルを華麗に連携！

```yaml
{{ includex("examples/runners/multi_protocol_test.yml") }}
```

### 💥 障害テストシナリオ - カオスエンジニアリングの極意！

```yaml
{{ includex("examples/runners/failure_test_scenario.yml") }}
```

## 🎆 まとめ - ランナーマスター誕生！

**やったぞ！** あなたは今、**6大ランナーを完全にマスター**した！

### 🏆 この章で身につけた7つの必殺技：

1. **🌐 HTTPランナー**: REST API、GraphQL、認証フローを**完全制圧**！
2. **🔗 gRPCランナー**: マイクロサービス間通信を**爆速テスト**！
3. **🗄️ DBランナー**: データベースを**思いのままに操作**！
4. **🌐 CDPランナー**: ブラウザを**完全自動化**！
5. **💻 SSHランナー**: リモートサーバーを**自由自在に監視**！
6. **⚙️ Execランナー**: ローカルコマンドを**瞬時に実行**！
7. **🌈 ランナーの組み合わせ**: **マルチプロトコルテスト**の達人に！

各ランナーを**絶妙に組み合わせれば**、どんなに複雑なシステムも**完璧にテスト**できる。あなたはもう、**プロトコルの支配者**だ！
