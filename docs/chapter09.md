# 第9章：リファレンス

この章では、runnの詳細な仕様、全ビルトイン関数一覧、エラーメッセージ一覧、よくある質問（FAQ）をまとめています。日常的な開発で参照できるリファレンスとしてご活用ください。

## YAMLスキーマ

### Runbookの基本構造

```yaml
# 必須フィールド
desc: string                    # シナリオの説明

# オプションフィールド
labels: [string]               # シナリオの分類ラベル
runners: map[string]Runner     # 使用するランナーの定義
vars: map[string]any          # シナリオ全体で使用する変数
needs: map[string]Need        # 依存関係の定義
concurrency: string|int       # 並行実行制御
if: string                    # 条件付き実行
loop: Loop                    # ループ設定

# 必須フィールド
steps: [Step]|map[string]Step  # 実行するステップ
```

### Step（ステップ）の構造

```yaml
# ステップの基本構造
step_name:
  desc: string                 # ステップの説明（オプション）
  if: string                   # 条件付き実行（オプション）
  loop: Loop                   # ループ設定（オプション）
  
  # 実行内容（いずれか一つ）
  req: HTTPRequest            # HTTPリクエスト
  greq: GRPCRequest          # gRPCリクエスト
  db: DBQuery                # データベースクエリ
  cdp: CDPAction             # ブラウザ操作
  ssh: SSHCommand            # SSH実行
  exec: ExecCommand          # ローカルコマンド実行
  include: Include           # 他のシナリオをインクルード
  
  # 検証・出力
  test: string               # テストアサーション（オプション）
  dump: map[string]any       # デバッグ出力（オプション）
```

### Runner（ランナー）の定義

```yaml
runners:
  # HTTPランナー
  api_name: "https://api.example.com"
  
  # または詳細設定
  api_name:
    type: http
    base_url: "https://api.example.com"
    timeout: 30s
    headers:
      User-Agent: "MyApp/1.0"
    
  # データベースランナー
  db_name: "postgres://user:pass@host:5432/dbname"
  
  # gRPCランナー
  grpc_name: "grpc://grpc.example.com:443"
  
  # CDPランナー
  browser_name: "chrome://new"
  
  # SSHランナー
  ssh_name: "ssh://user@host:22"
```

### Loop（ループ）の設定

```yaml
# 単純な回数指定
loop: 5

# 詳細設定
loop:
  count: 10                    # 最大実行回数
  until: string               # 終了条件（式）
  minInterval: 1              # 最小間隔（秒）
  maxInterval: 10             # 最大間隔（秒）
  jitter: true               # ランダムな間隔を追加
  multiplier: 1.5            # 間隔の倍率
  concurrent: 1              # 並行実行数
```

### HTTPRequest（HTTPリクエスト）の構造

```yaml
req:
  /path/to/endpoint:
    method:                    # get, post, put, patch, delete
      headers:
        Header-Name: "value"
      query:
        param1: "value1"
        param2: "value2"
      body:
        application/json:      # JSON形式
          key: "value"
        application/x-www-form-urlencoded:  # フォーム形式
          key: "value"
        multipart/form-data:   # マルチパート形式
          file: "@/path/to/file"
          field: "value"
        text/plain: |          # プレーンテキスト
          テキストデータ
      timeout: 30s             # タイムアウト
      followRedirect: true     # リダイレクト追従
```

### DBQuery（データベースクエリ）の構造

```yaml
db:
  runner_name:///
    query: |
      SELECT * FROM users
      WHERE id = $1
    params:
      - "{{ vars.user_id }}"
    timeout: 30s
```

### CDPAction（ブラウザ操作）の構造

```yaml
cdp:
  runner_name:///
    actions:
      - navigate: "https://example.com"
      - waitVisible: "selector"
      - click: "button"
      - type:
          selector: "input"
          text: "入力テキスト"
      - screenshot: "/path/to/screenshot.png"
      - evaluate: "JavaScript code"
```

## 全ビルトイン関数一覧

### 比較・差分関数

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `compare(x, y, ignorePaths...)` | 2つの値を比較 | `compare(actual, expected)` |
| `diff(x, y, ignorePaths...)` | 2つの値の差分を表示 | `diff(actual, expected, [".timestamp"])` |

### データ操作関数

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `pick(obj, keys...)` | 指定キーのみ抽出 | `pick(user, "id", "name")` |
| `omit(obj, keys...)` | 指定キーを除外 | `omit(user, "password")` |
| `merge(obj1, obj2...)` | オブジェクトをマージ | `merge(defaults, config)` |
| `intersect(arr1, arr2)` | 配列の積集合 | `intersect([1,2,3], [2,3,4])` |
| `union(arr1, arr2)` | 配列の和集合 | `union([1,2], [2,3])` |
| `unique(arr)` | 重複を除去 | `unique([1,2,2,3])` |
| `groupBy(arr, expr)` | 配列をグループ化 | `groupBy(users, {.department})` |

### 配列・文字列関数

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `len(x)` | 長さを取得 | `len(array)`, `len(string)` |
| `map(arr, expr)` | 配列の各要素を変換 | `map(users, {.name})` |
| `filter(arr, expr)` | 配列をフィルタリング | `filter(users, {.active})` |
| `sort(arr, expr)` | 配列をソート | `sort(users, {.name})` |
| `reverse(arr)` | 配列を逆順 | `reverse([1,2,3])` |
| `join(arr, sep)` | 配列を文字列に結合 | `join(names, ", ")` |
| `split(str, sep)` | 文字列を分割 | `split("a,b,c", ",")` |
| `contains(str, substr)` | 文字列が含まれるか | `contains(text, "hello")` |
| `startsWith(str, prefix)` | 文字列が指定文字で始まるか | `startsWith(text, "http")` |
| `endsWith(str, suffix)` | 文字列が指定文字で終わるか | `endsWith(text, ".com")` |
| `upper(str)` | 大文字に変換 | `upper("hello")` |
| `lower(str)` | 小文字に変換 | `lower("HELLO")` |
| `trim(str)` | 前後の空白を除去 | `trim(" hello ")` |

### 数値・計算関数

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `sum(arr)` | 配列の合計 | `sum([1,2,3])` |
| `min(arr)` | 配列の最小値 | `min([1,2,3])` |
| `max(arr)` | 配列の最大値 | `max([1,2,3])` |
| `avg(arr)` | 配列の平均値 | `avg([1,2,3])` |
| `abs(num)` | 絶対値 | `abs(-5)` |
| `ceil(num)` | 切り上げ | `ceil(3.14)` |
| `floor(num)` | 切り下げ | `floor(3.14)` |
| `round(num)` | 四捨五入 | `round(3.14)` |

### 型変換関数

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `string(x)` | 文字列に変換 | `string(123)` |
| `int(x)` | 整数に変換 | `int("123")` |
| `float(x)` | 浮動小数点数に変換 | `float("3.14")` |
| `bool(x)` | 真偽値に変換 | `bool("true")` |
| `toJSON(x)` | JSON文字列に変換 | `toJSON(object)` |
| `fromJSON(str)` | JSON文字列から変換 | `fromJSON(jsonStr)` |

### エンコーディング関数

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `urlencode(str)` | URLエンコード | `urlencode("hello world")` |
| `urldecode(str)` | URLデコード | `urldecode("hello%20world")` |
| `toBase64(str)` | Base64エンコード | `toBase64("hello")` |
| `fromBase64(str)` | Base64デコード | `fromBase64("aGVsbG8=")` |
| `toHex(str)` | 16進数エンコード | `toHex("hello")` |
| `fromHex(str)` | 16進数デコード | `fromHex("68656c6c6f")` |

### 時間関数

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `time(x)` | 時刻オブジェクトに変換 | `time("2024-01-01T00:00:00Z")` |
| `now()` | 現在時刻 | `now()` |
| `unix(t)` | Unix timestamp | `unix(time("2024-01-01"))` |
| `format(t, layout)` | 時刻をフォーマット | `format(now(), "2006-01-02")` |
| `parse(str, layout)` | 文字列を時刻に変換 | `parse("2024-01-01", "2006-01-02")` |

### ファイル・URL関数

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `file(path)` | ファイルを読み込み | `file("data.json")` |
| `url(rawURL)` | URLを解析 | `url("https://example.com/path")` |

### Faker関数（テストデータ生成）

#### 個人情報

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `faker.name()` | ランダムな名前 | `faker.name()` |
| `faker.firstName()` | 名前（名） | `faker.firstName()` |
| `faker.lastName()` | 名前（姓） | `faker.lastName()` |
| `faker.email()` | メールアドレス | `faker.email()` |
| `faker.phone()` | 電話番号 | `faker.phone()` |
| `faker.username()` | ユーザー名 | `faker.username()` |

#### 住所情報

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `faker.address()` | 住所 | `faker.address()` |
| `faker.city()` | 都市名 | `faker.city()` |
| `faker.state()` | 州・県名 | `faker.state()` |
| `faker.country()` | 国名 | `faker.country()` |
| `faker.zipCode()` | 郵便番号 | `faker.zipCode()` |

#### インターネット関連

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `faker.url()` | URL | `faker.url()` |
| `faker.domainName()` | ドメイン名 | `faker.domainName()` |
| `faker.ipv4()` | IPv4アドレス | `faker.ipv4()` |
| `faker.ipv6()` | IPv6アドレス | `faker.ipv6()` |
| `faker.macAddress()` | MACアドレス | `faker.macAddress()` |

#### 数値・文字列

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `faker.number(min, max)` | ランダムな数値 | `faker.number(1, 100)` |
| `faker.float(min, max, precision)` | ランダムな浮動小数点数 | `faker.float(0.0, 1.0, 2)` |
| `faker.uuid()` | UUID | `faker.uuid()` |
| `faker.randomString(length)` | ランダムな文字列 | `faker.randomString(10)` |
| `faker.password(lower, upper, numeric, special, space, length)` | パスワード | `faker.password(true, true, true, false, false, 12)` |

#### 日時

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `faker.date()` | ランダムな日付 | `faker.date()` |
| `faker.dateRange(start, end)` | 期間内のランダムな日付 | `faker.dateRange("2024-01-01", "2024-12-31")` |
| `faker.time()` | ランダムな時刻 | `faker.time()` |

#### 選択・真偽値

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `faker.randomChoice(arr)` | 配列からランダム選択 | `faker.randomChoice(["A", "B", "C"])` |
| `faker.randomBool()` | ランダムな真偽値 | `faker.randomBool()` |
| `faker.randomInt(min, max)` | ランダムな整数 | `faker.randomInt(1, 100)` |

### その他の関数

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `input(prompt, default)` | ユーザー入力を取得 | `input("Enter name", "default")` |
| `env(key, default)` | 環境変数を取得 | `env("HOME", "/tmp")` |
| `range(start, end, step)` | 数値の範囲を生成 | `range(1, 10, 2)` |
| `keys(obj)` | オブジェクトのキー一覧 | `keys({"a": 1, "b": 2})` |
| `values(obj)` | オブジェクトの値一覧 | `values({"a": 1, "b": 2})` |

## エラーメッセージ一覧

### 一般的なエラー

| エラーメッセージ | 原因 | 対処法 |
|------------------|------|--------|
| `failed to load runbook` | YAMLファイルの読み込みエラー | ファイルパスとYAML構文を確認 |
| `invalid YAML format` | YAML構文エラー | YAMLの構文を確認（インデント、引用符など） |
| `step not found` | 存在しないステップを参照 | ステップ名のスペルミスを確認 |
| `runner not found` | 存在しないランナーを参照 | ランナー名のスペルミスを確認 |

### HTTP関連エラー

| エラーメッセージ | 原因 | 対処法 |
|------------------|------|--------|
| `connection refused` | サーバーに接続できない | サーバーが起動しているか確認 |
| `timeout` | リクエストがタイムアウト | タイムアウト値を調整またはサーバー応答を確認 |
| `invalid URL` | 不正なURL形式 | URL形式を確認 |
| `unsupported content type` | サポートされていないContent-Type | 適切なContent-Typeを指定 |

### データベース関連エラー

| エラーメッセージ | 原因 | 対処法 |
|------------------|------|--------|
| `database connection failed` | データベース接続エラー | 接続文字列とデータベース状態を確認 |
| `SQL syntax error` | SQL構文エラー | SQLクエリの構文を確認 |
| `table not found` | テーブルが存在しない | テーブル名とスキーマを確認 |
| `permission denied` | 権限不足 | データベースユーザーの権限を確認 |

### 式評価エラー

| エラーメッセージ | 原因 | 対処法 |
|------------------|------|--------|
| `expression evaluation failed` | 式の評価エラー | 式の構文と変数の存在を確認 |
| `variable not found` | 変数が存在しない | 変数名のスペルミスと定義を確認 |
| `type mismatch` | 型の不一致 | 期待される型と実際の型を確認 |
| `division by zero` | ゼロ除算エラー | 除数がゼロでないことを確認 |

### ブラウザ操作エラー

| エラーメッセージ | 原因 | 対処法 |
|------------------|------|--------|
| `element not found` | 要素が見つからない | セレクターと要素の存在を確認 |
| `element not visible` | 要素が表示されていない | 要素の表示状態を確認 |
| `browser not responding` | ブラウザが応答しない | ブラウザの状態とリソース使用量を確認 |
| `navigation failed` | ページ遷移に失敗 | URLとネットワーク接続を確認 |

## FAQ（よくある質問）

### 基本的な使い方

**Q: runnとは何ですか？**
A: runnは、シナリオベースのテスト・自動化ツールです。YAMLでテストシナリオを記述し、HTTP、gRPC、データベース、ブラウザ操作などを統一的に扱えます。

**Q: CLIツールとGoテストヘルパーの違いは何ですか？**
A: CLIツールは単独でシナリオを実行しますが、Goテストヘルパーは`go test`と統合してより柔軟なテスト環境を構築できます。本書では特にGoテストヘルパーとしての利用を推奨しています。

**Q: どのようなプロトコルをサポートしていますか？**
A: HTTP/HTTPS、gRPC、データベース（PostgreSQL、MySQL、SQLite、Cloud Spanner）、ブラウザ操作（CDP）、SSH、ローカルコマンド実行をサポートしています。

### インストールと設定

**Q: インストール方法を教えてください。**
A: 複数の方法があります：
- Homebrew: `brew install k1LoW/tap/runn`
- Go install: `go install github.com/k1LoW/runn/cmd/runn@latest`
- バイナリダウンロード: GitHub Releasesから取得

**Q: Goプロジェクトに統合するにはどうすればよいですか？**
A: `go.mod`に`github.com/k1LoW/runn`を追加し、テストファイルで`runn.Load()`を使用してシナリオを実行します。

### シナリオ記述

**Q: リスト形式とマップ形式のどちらを使うべきですか？**
A: シンプルなシナリオはリスト形式、複雑なシナリオや可読性を重視する場合はマップ形式を推奨します。

**Q: 変数はどのように定義・参照しますか？**
A: `vars:`セクションで定義し、`{{ vars.変数名 }}`で参照します。環境変数は`{{ env.環境変数名 }}`で参照できます。

**Q: 前のステップの結果を次のステップで使用するには？**
A: `steps.ステップ名.res.body`（マップ形式）または`steps[インデックス].res.body`（リスト形式）で参照します。

### エラーハンドリング

**Q: テストが失敗した時の詳細情報を取得するには？**
A: `dump:`セクションを使用してデバッグ情報を出力し、`test: true`でエラーでも続行させることができます。

**Q: リトライ機能はありますか？**
A: `loop:`セクションで`until:`条件と`minInterval:`、`maxInterval:`を指定することでリトライ機能を実装できます。

**Q: 条件付きでステップを実行するには？**
A: `if:`フィールドに条件式を記述することで、条件付き実行が可能です。

### パフォーマンス

**Q: 並列実行は可能ですか？**
A: `concurrency:`フィールドで並列実行数を制御できます。共有リソースを使用する場合は同じキーを指定して順次実行も可能です。

**Q: 大量のテストデータを効率的に処理するには？**
A: `loop:`を使用した繰り返し処理、Faker関数による動的データ生成、外部ファイルからのデータ読み込みを組み合わせます。

**Q: パフォーマンステストは実行できますか？**
A: 可能です。`loop:`で大量のリクエストを生成し、レスポンス時間やスループットを測定できます。

### 統合・連携

**Q: CI/CDパイプラインで実行するには？**
A: GitHub ActionsやJenkinsなどで`go test`コマンドを実行するだけです。環境変数でテスト環境を切り替えることも可能です。

**Q: 既存のテストフレームワークと併用できますか？**
A: はい。runnはGoの標準的なテストフレームワークと統合されているため、既存のテストと併用できます。

**Q: モックサーバーとの連携は可能ですか？**
A: 可能です。`httptest.NewServer()`で作成したモックサーバーのURLをランナーに設定することで連携できます。

### トラブルシューティング

**Q: 「runner not found」エラーが発生します。**
A: `runners:`セクションでランナーが正しく定義されているか、ステップでのランナー名にスペルミスがないか確認してください。

**Q: データベース接続エラーが発生します。**
A: 接続文字列の形式、データベースサーバーの起動状態、ネットワーク接続、認証情報を確認してください。

**Q: ブラウザ操作で要素が見つからないエラーが発生します。**
A: セレクターが正しいか、要素が表示されるまで適切に待機しているか、ページの読み込みが完了しているかを確認してください。

**Q: 式評価でエラーが発生します。**
A: 変数名のスペルミス、変数の存在、型の不一致、構文エラーを確認してください。`dump:`を使用してデバッグ情報を出力することも有効です。

### ベストプラクティス

**Q: テストデータはどのように管理すべきですか？**
A: 固定データは再現性のために、ランダムデータは多様性のために使い分けます。外部ファイルからの読み込みも活用しましょう。

**Q: 大規模なプロジェクトでの構成はどうすべきですか？**
A: 機能別にディレクトリを分け、共通処理は`include:`で再利用し、環境別設定を適切に管理することを推奨します。

**Q: セキュリティ面で注意すべき点はありますか？**
A: 認証情報は環境変数で管理し、ログに機密情報が出力されないよう注意し、テスト用データベースは本番から分離してください。

## バージョン情報と互換性

### サポートされるバージョン

- **runn**: v0.100.0以降
- **Go**: 1.21以降
- **PostgreSQL**: 12以降
- **MySQL**: 8.0以降
- **Chrome/Chromium**: 90以降

### 変更履歴の重要なポイント

- v0.100.0: 安定版リリース、APIの大幅な変更
- v0.90.0: Goテストヘルパー機能の強化
- v0.80.0: CDP（ブラウザ操作）サポート追加
- v0.70.0: 並行実行制御機能追加

## 参考リンク

- **公式リポジトリ**: https://github.com/k1LoW/runn
- **ドキュメント**: https://github.com/k1LoW/runn/tree/main/docs
- **サンプル**: https://github.com/k1LoW/runn/tree/main/testdata
- **Issue報告**: https://github.com/k1LoW/runn/issues
- **expr-lang/expr**: https://github.com/expr-lang/expr

## まとめ

この章では、runnの詳細な仕様とリファレンス情報を提供しました：

1. **YAMLスキーマ**: Runbook、Step、Runnerの詳細な構造
2. **全ビルトイン関数一覧**: 比較、データ操作、文字列、数値、Faker関数など
3. **エラーメッセージ一覧**: よくあるエラーとその対処法
4. **FAQ**: 基本的な使い方からトラブルシューティングまで

このリファレンスを活用して、効率的にrunnを使いこなしてください。runnは継続的に開発が進められているため、最新の情報は公式リポジトリも併せてご確認ください。

---

**本書の完了**

これで「runn入門」の全9章が完成しました。runnの基本的な使い方から高度な活用方法、実践的なテクニック、詳細なリファレンスまでを網羅しています。

特に第7章「Goテストヘルパー編」で解説したように、runnの真の価値はGoのテストフレームワークと統合することで発揮されます。YAMLの宣言的な記述とGoの強力なテスト機能を組み合わせることで、保守性が高く理解しやすい高品質なテストスイートを構築できます。

本書がrunnを活用したテスト自動化の一助となれば幸いです。