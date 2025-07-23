# 第5章：ランナー詳細編

runnの特徴の一つは、複数のプロトコルを統一的に扱えることです。この章では、各ランナーの使い方を詳しく説明します。

## runnが対応するランナー

runnは6つのランナーを提供しています。それぞれが特定のプロトコルに対応します。

| ランナー | プロトコル | 用途 |
|----------|------------|------|
| HTTP | HTTP/HTTPS | REST API、GraphQL、Webhookのテスト |
| gRPC | gRPC | マイクロサービス間通信のテスト |
| DB | SQL | データベース操作のテスト |
| CDP | Chrome DevTools Protocol | ブラウザ自動化 |
| SSH | SSH | リモートサーバーの操作 |
| Exec | プロセス実行 | ローカルコマンドの実行 |

## HTTPランナー

### 基本的な設定

```yaml
{{ includex("examples/runners/http_basic_setup.yml") }}
```

### リクエストメソッドとパラメータ

```yaml
{{ includex("examples/runners/http_request_methods.yml") }}
```

### 様々なボディ形式

```yaml
{{ includex("examples/runners/http_body_formats.yml") }}
```

## gRPCランナー

TODO: grpc の例を追加する

## データベースランナー

### 対応データベース

```yaml
{{ includex("examples/runners/db_connections.yml") }}
```

### 基本的なクエリ操作

<!-- TODO: INSERT のあと RETRUNING 使えてない: https://github.com/k1LoW/runn/issues/1276 -->

```yaml
{{ includex("examples/runners/db_basic_queries.yml") }}
```

## CDPランナー（ブラウザ自動化）

### 主なCDP actions一覧

| アクション名      | 概要                                   |
|------------------|----------------------------------------|
| attributes       | 要素の属性取得                         |
| click            | 要素をクリック                         |
| doubleClick      | 要素をダブルクリック                   |
| evaluate         | JS式の評価                             |
| fullHTML         | ページ全体のHTML取得                   |
| innerHTML        | 要素のinnerHTML取得                    |
| localStorage     | localStorage取得                       |
| location         | 現在のURL取得                          |
| navigate         | 指定URLへ遷移                          |
| outerHTML        | 要素のouterHTML取得                    |
| screenshot       | スクリーンショット取得                 |
| scroll           | 要素までスクロール                     |
| sendKeys         | 要素にキー入力                         |
| sessionStorage   | sessionStorage取得                     |
| setUploadFile    | ファイルアップロード                   |
| setUserAgent     | User-Agent設定                         |
| submit           | フォーム送信                           |
| tabTo            | タブ切り替え                           |
| text             | 要素のテキスト取得                     |
| textContent      | 要素のtextContent取得                  |
| title            | ページタイトル取得                     |
| value            | 要素のvalue取得                        |
| wait             | 指定時間待機                           |
| waitReady        | 要素の準備完了まで待機                 |
| waitVisible      | 要素の表示まで待機                     |

※詳細・最新情報は[公式README](https://github.com/k1LoW/runn?tab=readme-ov-file#functions-for-action-to-control-browser)をご参照ください。

### 基本的な使い方

```yaml
{{ includex("examples/runners/cdp_basic.yml") }}
```

## SSHランナー

### 基本的な設定

```yaml
{{ includex("examples/runners/ssh_basic.yml") }}
```

### サーバー監視とヘルスチェック

```yaml
{{ includex("examples/runners/ssh_health_check.yml") }}
```

## Execランナー（ローカルコマンド実行）

### 基本的な使用方法

```yaml
{{ includex("examples/runners/exec_basic.yml") }}
```

### ファイル操作とテスト

```yaml
{{ includex("examples/runners/exec_file_operations.yml") }}
```

## まとめ

この章では、runnが提供する6つのランナーについて説明しました。

### この章で学んだこと：

1. HTTPランナー: REST API、GraphQL、認証フローのテスト
2. gRPCランナー: マイクロサービス間通信のテスト
3. DBランナー: データベース操作のテスト
4. CDPランナー: ブラウザ自動化
5. SSHランナー: リモートサーバーの監視と操作
6. Execランナー: ローカルコマンドの実行
7. ランナーの組み合わせ: マルチプロトコルテスト

各ランナーを組み合わせることで、複雑なシステムのテストも実現できます。
