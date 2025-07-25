# 第8章：CLIオプション

## コマンド一覧

### 基本コマンド

#### `runn run` - シナリオ実行
```bash
runn run [PATH_PATTERN ...] [flags]
```

#### `runn list` - シナリオ一覧表示
```bash
runn list [PATH_PATTERN ...]
```

#### `runn new` - シナリオ生成
```bash
runn new [flags]
```

#### `runn loadt` - 負荷テスト実行
```bash
runn loadt [PATH_PATTERN ...] [flags]
```

#### `runn coverage` - カバレッジ表示
```bash
runn coverage [flags]
```

## run コマンドのオプション

### 基本オプション

#### `--verbose`
詳細なログを出力します。デバッグ時に便利です。
```bash
runn run scenario.yml --verbose
```

#### `--debug`
デバッグモードで実行します。より詳細な情報が出力されます。
```bash
runn run scenario.yml --debug
```

### 実行制御

#### `--fail-fast`
最初のエラーで実行を停止します。CI/CDで時間を節約できます。
```bash
runn run tests/**/*.yml --fail-fast
```

#### `--skip-test`
`test:` セクションをスキップします。動作確認時に便利です。
```bash
runn run scenario.yml --skip-test
```

#### `--skip-included`
インクルードされたrunbookの単体実行をスキップします。
```bash
runn run --skip-included
```

### 並列実行

#### `--concurrent`
並列実行を制御します。デフォルトは "off" です。
```bash
# 5並列で実行
runn run tests/**/*.yml --concurrent 5

# 並列実行を有効化（デフォルト並列数）
runn run tests/**/*.yml --concurrent on
```

#### `--shuffle`
実行順序をランダム化します。
```bash
# ランダム化を有効化
runn run tests/**/*.yml --shuffle on

# シード値を指定してランダム化
runn run tests/**/*.yml --shuffle 42
```

### フィルタリング

#### `--label`
ラベルでrunbookをフィルタリングします。
```bash
# 単一ラベル
runn run tests/**/*.yml --label api

# 複数ラベル（AND条件）
runn run tests/**/*.yml --label api --label critical

# 複雑なラベル条件
runn run tests/**/*.yml --label "users and auth"
runn run tests/**/*.yml --label "(users or projects) and not slow"
```

#### `--id`
IDでrunbookをフィルタリングします。
```bash
runn run tests/**/*.yml --id test-user-creation
```

#### `--run`
ファイルパスの正規表現でフィルタリングします。
```bash
runn run tests/**/*.yml --run "user.*test"
```

#### `--sample`
指定した数のrunbookをサンプリングして実行します。
```bash
runn run tests/**/*.yml --sample 10
```

#### `--random`
指定した数のrunbookをランダムに選択して実行します。
```bash
runn run tests/**/*.yml --random 5
```

### 変数とデータ

#### `--var`
変数を設定します。
```bash
runn run scenario.yml --var "base_url:https://api.example.com" --var "api_key:secret123"
```

#### `--env-file`
環境変数をファイルから読み込みます。
```bash
runn run scenario.yml --env-file .env.test
```

#### `--overlay`
runbookに値を上書きします。
```bash
runn run scenario.yml --overlay overlay.yml
```

#### `--underlay`
runbookの下に値を配置します。
```bash
runn run scenario.yml --underlay defaults.yml
```

### ランナー設定

#### `--runner`
ランナーをrunbookに設定します。
```bash
runn run scenario.yml --runner "db:postgres://user:pass@localhost/testdb"
```

#### `--host-rules`
ホストルールを設定します。
```bash
runn run scenario.yml --host-rules "api.example.com=localhost:8080,cdn.example.com=localhost:8081"
```

### HTTPランナーオプション

#### `--http-openapi3`
OpenAPI v3ドキュメントのパスを設定します。
```bash
runn run scenario.yml --http-openapi3 openapi.yml
runn run scenario.yml --http-openapi3 api:openapi.yml
```

### gRPCランナーオプション

#### `--grpc-proto`
protoソースファイルを指定します。
```bash
runn run scenario.yml --grpc-proto service.proto
```

#### `--grpc-import-path`
protoファイルのインポートパスを設定します。
```bash
runn run scenario.yml --grpc-import-path ./proto
```

#### `--grpc-no-tls`
すべてのgRPCランナーでTLSを無効化します。
```bash
runn run scenario.yml --grpc-no-tls
```

#### `--grpc-buf-*`
buf関連の設定を行います。
```bash
runn run scenario.yml --grpc-buf-config buf.yaml --grpc-buf-dir ./proto
```

### 出力とレポート

#### `--format`
結果の出力フォーマットを指定します。
```bash
# JSON形式で出力
runn run scenario.yml --format json

# JUnit XML形式で出力
runn run scenario.yml --format junit
```

#### `--capture`
実行結果をキャプチャして保存します。
```bash
runn run scenario.yml --capture results/
```

#### `--profile`
プロファイリングを有効化します。
```bash
runn run scenario.yml --profile --profile-out profile.prof
```

### 分散実行

#### `--shard-n` / `--shard-index`
runbookを分散実行します。
```bash
# 3つのシャードに分割し、1番目のシャードを実行
runn run tests/**/*.yml --shard-n 3 --shard-index 0
```

### その他のオプション

#### `--attach`
runnプロセスにアタッチします。
```bash
runn run scenario.yml --attach
```

#### `--force-color`
非TTY環境でも色付き出力を強制します。
```bash
runn run scenario.yml --force-color
```

#### `--wait-timeout`
クリーンアップ処理のタイムアウトを設定します。
```bash
runn run scenario.yml --wait-timeout 30s
```

#### `--cache-dir`
リモートrunbookのキャッシュディレクトリを指定します。
```bash
runn run https://example.com/scenario.yml --cache-dir .cache/
```

#### `--retain-cache-dir`
リモートrunbookのキャッシュを保持します。
```bash
runn run https://example.com/scenario.yml --retain-cache-dir
```

## new コマンドのオプション

#### `--and-run`
作成後すぐに実行します。
```bash
runn new --and-run --out test.yml -- curl https://api.example.com
```

#### `--desc`
シナリオに説明を追加します。
```bash
runn new --desc "User API test" --out user-test.yml
```

#### `--out`
出力ファイル名を指定します。
```bash
runn new --out scenario.yml -- curl https://api.example.com
```

## グローバルオプション

#### `--scopes`
runnの追加スコープを設定します。
```bash
runn run scenario.yml --scopes "run:exec"
```

## 活用例

### CI/CDでの利用
```bash
# 高速フェイルでAPIテストを並列実行
runn run tests/api/**/*.yml \
  --label api \
  --label "not slow" \
  --concurrent 10 \
  --fail-fast \
  --format junit > test-results.xml
```

### デバッグ時の利用
```bash
# 詳細ログ付きで特定のテストを実行
runn run tests/user-api.yml \
  --verbose \
  --debug \
  --var "base_url:http://localhost:8080" \
  --skip-test
```

### 負荷テスト
```bash
# 10並列で30秒間負荷テストを実行
runn loadt tests/api/**/*.yml \
  --concurrent 10 \
  --duration 30s \
  --threshold "error_rate < 0.01"
```