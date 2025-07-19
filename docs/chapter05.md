# 第5章：ランナー詳細編

runnの最大の特徴の一つは、複数のプロトコルを統一的に扱えることです。この章では、各ランナーの詳細な使い方と実践的な活用方法について解説します。

## ランナーの概要

runnは以下のランナーをサポートしています：

| ランナー | プロトコル | 用途 |
|----------|------------|------|
| **HTTP** | HTTP/HTTPS | REST API、GraphQL、Webhook |
| **gRPC** | gRPC | マイクロサービス間通信 |
| **DB** | SQL | データベース操作・検証 |
| **CDP** | Chrome DevTools Protocol | ブラウザ自動化・E2Eテスト |
| **SSH** | SSH | リモートサーバー操作 |
| **Exec** | プロセス実行 | ローカルコマンド実行 |

## HTTPランナー

### 基本的な設定

```yaml
{{ includex("examples/chapter05/http_basic_setup.yml") }}
```

### リクエストメソッドとパラメータ

```yaml
{{ includex("examples/chapter05/http_request_methods.yml") }}
```

### 様々なボディ形式

```yaml
{{ includex("examples/chapter05/http_body_formats.yml") }}
```

### 認証の実装

```yaml
{{ includex("examples/chapter05/http_authentication.yml") }}
```

### レスポンスの詳細な検証

```yaml
{{ includex("examples/chapter05/http_response_validation.yml") }}
```

### GraphQL API の操作

```yaml
{{ includex("examples/chapter05/graphql_example.yml") }}
```

## gRPCランナー

### 基本的な設定

```yaml
{{ includex("examples/chapter05/grpc_basic.yml") }}
```

### プロトコルバッファの動的読み込み

```yaml
{{ includex("examples/chapter05/grpc_dynamic_proto.yml") }}
```

## データベースランナー

### 対応データベース

```yaml
{{ includex("examples/chapter05/db_connections.yml") }}
```

### 基本的なクエリ操作

```yaml
{{ includex("examples/chapter05/db_basic_queries.yml") }}
```

### トランザクション処理

```yaml
{{ includex("examples/chapter05/db_transactions.yml") }}
```

### 複雑なデータ検証

```yaml
{{ includex("examples/chapter05/db_complex_validation.yml") }}
```

## CDPランナー（ブラウザ自動化）

### 基本的な設定

```yaml
{{ includex("examples/chapter05/cdp_basic.yml") }}
```

### 高度なブラウザ操作

```yaml
{{ includex("examples/chapter05/cdp_advanced.yml") }}
```

### SPAアプリケーションのテスト

```yaml
{{ includex("examples/chapter05/cdp_spa_testing.yml") }}
```

## SSHランナー

### 基本的な設定

```yaml
{{ includex("examples/chapter05/ssh_basic.yml") }}
```

### サーバー監視とヘルスチェック

```yaml
{{ includex("examples/chapter05/ssh_health_check.yml") }}
```

## Execランナー（ローカルコマンド実行）

### 基本的な使用方法

```yaml
{{ includex("examples/chapter05/exec_basic.yml") }}
```

### ファイル操作とテスト

```yaml
{{ includex("examples/chapter05/exec_file_operations.yml") }}
```

## ランナーの組み合わせ

### マルチプロトコルテスト

```yaml
{{ includex("examples/chapter05/multi_protocol_test.yml") }}
```

### 障害テストシナリオ

```yaml
{{ includex("examples/chapter05/failure_test_scenario.yml") }}
```

## まとめ

この章では、runnの各ランナーについて詳しく学びました：

1. **HTTPランナー**: REST API、GraphQL、認証フローの実装
2. **gRPCランナー**: マイクロサービス間通信のテスト
3. **DBランナー**: データベース操作とトランザクション処理
4. **CDPランナー**: ブラウザ自動化とE2Eテスト
5. **SSHランナー**: リモートサーバーの操作と監視
6. **Execランナー**: ローカルコマンドの実行
7. **ランナーの組み合わせ**: マルチプロトコルテストの実現

各ランナーを適切に組み合わせることで、包括的なテストシナリオを構築できます。次章では、ループ処理や条件分岐などの高度な機能について学んでいきます。

[第6章：高度な機能編へ →](chapter06.md)