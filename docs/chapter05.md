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
{{ includex("examples/chapter05/http_basic_setup.yml") }}
```

### 🎨 リクエストメソッドとパラメータ - あらゆるHTTPメソッドを制覇！

```yaml
{{ includex("examples/chapter05/http_request_methods.yml") }}
```

### 📦 様々なボディ形式 - どんなデータ形式もお手のもの！

```yaml
{{ includex("examples/chapter05/http_body_formats.yml") }}
```

### 🔐 認証の実装 - セキュリティを完璧にテスト！

```yaml
{{ includex("examples/chapter05/http_authentication.yml") }}
```

### 🔍 レスポンスの詳細な検証 - 一分の隙も見逃さない！

```yaml
{{ includex("examples/chapter05/http_response_validation.yml") }}
```

### 🚀 GraphQL API の操作 - 次世代APIも完全サポート！

```yaml
{{ includex("examples/chapter05/graphql_example.yml") }}
```

## 🔗 gRPCランナー - マイクロサービスの強い味方！

### ⚡ 基本的な設定 - 高速通信の世界へ！

```yaml
{{ includex("examples/chapter05/grpc_basic.yml") }}
```

### 📚 プロトコルバッファの動的読み込み - protoファイルを瞬時に理解！

```yaml
{{ includex("examples/chapter05/grpc_dynamic_proto.yml") }}
```

## 🗄️ データベースランナー - SQLの魔術師になれ！

### 🌍 対応データベース - あらゆるDBを制覇！

```yaml
{{ includex("examples/chapter05/db_connections.yml") }}
```

### 📝 基本的なクエリ操作 - SQLを思いのままに！

```yaml
{{ includex("examples/chapter05/db_basic_queries.yml") }}
```

### 🔄 トランザクション処理 - データ整合性を守り抜け！

```yaml
{{ includex("examples/chapter05/db_transactions.yml") }}
```

### 🧩 複雑なデータ検証 - どんなデータも完璧にテスト！

```yaml
{{ includex("examples/chapter05/db_complex_validation.yml") }}
```

## 🌐 CDPランナー（ブラウザ自動化） - ブラウザを完全支配！

### 🎮 基本的な設定 - Chromeを思いのままに操ろう！

```yaml
{{ includex("examples/chapter05/cdp_basic.yml") }}
```

### 🎪 高度なブラウザ操作 - プロ級のテクニック！

```yaml
{{ includex("examples/chapter05/cdp_advanced.yml") }}
```

### ✨ SPAアプリケーションのテスト - モダンWebアプリも余裕！

```yaml
{{ includex("examples/chapter05/cdp_spa_testing.yml") }}
```

## 💻 SSHランナー - リモートサーバーの絶対的支配者！

### 🔑 基本的な設定 - サーバーへのセキュアアクセス！

```yaml
{{ includex("examples/chapter05/ssh_basic.yml") }}
```

### 📏 サーバー監視とヘルスチェック - 24時間365日の番人！

```yaml
{{ includex("examples/chapter05/ssh_health_check.yml") }}
```

## ⚙️ Execランナー（ローカルコマンド実行） - シェルコマンドの魔術師！

### 🚀 基本的な使用方法 - コマンドを瞬時に実行！

```yaml
{{ includex("examples/chapter05/exec_basic.yml") }}
```

### 📁 ファイル操作とテスト - ローカルファイルを完全管理！

```yaml
{{ includex("examples/chapter05/exec_file_operations.yml") }}
```

## 🎆 ランナーの組み合わせ - 最強のコンボ技！

### 🌈 マルチプロトコルテスト - 複数プロトコルを華麗に連携！

```yaml
{{ includex("examples/chapter05/multi_protocol_test.yml") }}
```

### 💥 障害テストシナリオ - カオスエンジニアリングの極意！

```yaml
{{ includex("examples/chapter05/failure_test_scenario.yml") }}
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

**次章では、ループ処理や条件分岐などの高度な機能をマスターして、真のrunnマスターへの道を進もう！** 準備はいいか？

[第6章：高度な機能編へ →](chapter06.md)