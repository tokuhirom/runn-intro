# runn - シナリオベースのテスト・自動化ツール

## 概要

**runn**（「Run N」と読み、/rʌ́n én/と発音）は、シナリオに従って操作を実行するためのパッケージ/ツールです。

- **GitHubリポジトリ**: https://github.com/k1LoW/runn
- **作者**: k1LoW
- **ライセンス**: MITライセンス

## 主な特徴

1. **シナリオベースのテストツール**として使用可能
2. **Go言語のテストヘルパーパッケージ**として使用可能
3. **ワークフロー自動化ツール**として使用可能
4. **HTTP/gRPC/DB/CDP/SSHの実行をサポート**
5. **OpenAPI Document風の構文**でHTTPリクエストテストが可能
6. **シングルバイナリ**でCI/CDフレンドリー

## 基本的な使い方

### クイックスタート

#### curlコマンドからシナリオを作成
```bash
# curlコマンドを実行してレスポンスを確認
$ curl https://httpbin.org/json -H "accept: application/json"

# runnでシナリオを作成して実行
$ runn new --and-run --desc 'httpbin.org GET' --out http.yml -- curl https://httpbin.org/json -H "accept: application/json"
```

#### アクセスログからシナリオを作成
```bash
$ cat access_log | runn new --out axslog.yml
```

### 基本コマンド

```bash
# シナリオ一覧を表示
$ runn list path/to/**/*.yml

# シナリオを実行
$ runn run path/to/**/*.yml

# ラベルでフィルタリングして実行
$ runn run path/to/**/*.yml --label users --label projects
```

## Runbook（シナリオファイル）の構造

### 基本構造（リスト形式）

```yaml
desc: ログインしてプロジェクトを取得
runners:
  req: https://example.com/api/v1
  db: mysql://root:mypass@localhost:3306/testdb
vars:
  username: alice
  password: ${TEST_PASS}
steps:
  - db:
      query: SELECT * FROM users WHERE name = '{{ vars.username }}'
  - req:
      /login:
        post:
          body:
            application/json:
              email: "{{ steps[0].rows[0].email }}"
              password: "{{ vars.password }}"
    test: steps[1].res.status == 200
  - req:
      /projects:
        get:
          headers:
            Authorization: "token {{ steps[1].res.body.session_token }}"
    test: steps[2].res.status == 200
  - test: len(steps[2].res.body.projects) > 0
```

### 基本構造（マップ形式）

```yaml
desc: ログインしてプロジェクトを取得
runners:
  req: https://example.com/api/v1
  db: mysql://root:mypass@localhost:3306/testdb
vars:
  username: alice
  password: ${TEST_PASS}
steps:
  find_user:
    db:
      query: SELECT * FROM users WHERE name = '{{ vars.username }}'
  login:
    req:
      /login:
        post:
          body:
            application/json:
              email: "{{ steps.find_user.rows[0].email }}"
              password: "{{ vars.password }}"
    test: steps.login.res.status == 200
  list_projects:
    req:
      /projects:
        get:
          headers:
            Authorization: "token {{ steps.login.res.body.session_token }}"
    test: steps.list_projects.res.status == 200
  count_projects:
    test: len(steps.list_projects.res.body.projects) > 0
```

## 式評価エンジン（Expression Evaluation Engine）

runnは**expr-lang/expr**を式評価エンジンとして使用しています。

### 基本的な式の例

```yaml
steps:
  filterItemId:
    dump: filter(vars.items, {.itemId == vars.itemId})[0].name
  testAllPriceGte50:
    test: all(vars.items, {.price >= 50}) == true
  lenPriceEqual1000:
    test: len(filter(vars.items, {.price == 100})) == 1
  getProductWithKey:
    dump: vars.products["A"]
  getProductWithKey2:
    dump: vars.products.B
  concatString:
    dump: ("k1LoW/" + vars.keyString)
  containsString:
    test: (vars.keyString startsWith "run") == true
  whereIn:
    test: (vars.keyString in ["runn", "hoge", "fuga"]) == true
```

### 変数の参照

| 変数名 | 説明 |
|--------|------|
| `vars` | `vars:`セクションで設定された値 |
| `steps` | 各ステップの戻り値 |
| `i` | ループインデックス（`loop:`セクション内のみ） |
| `env` | 環境変数 |
| `current` | 現在のステップの戻り値 |
| `previous` | 前のステップの戻り値 |
| `parent` | 親runbookの変数（includeされた場合のみ） |

## ビルトイン関数

### 比較・差分関数
- `compare(x, y, ignorePaths)` - 2つの値を比較（ignorePaths引数でjqシンタックスパス式を指定して除外可能）
- `diff(x, y, ignorePaths)` - 2つの値の差分を表示（ignorePaths引数でjqシンタックスパス式を指定して除外可能）

### データ操作関数
- `pick(x, keys...)` - マップから指定したキーの値を抽出
- `omit(x, keys...)` - マップから指定したキーを除外
- `merge(x...)` - 複数のマップをマージ
- `intersect(x, y)` - 2つの配列の積集合を返す

### 文字列・URL関数
- `urlencode(str)` - URLエンコード
- `url(rawURL)` - URL解析
- `toBase64(str)` - Base64エンコード

### 時間関数
- `time(v)` - 文字列や数値を時刻に変換

### ファイル関数
- `file(path)` - ファイルを読み込み

### Faker関数（テストデータ生成）
- `faker.name()` - ランダムな名前を生成
- `faker.email()` - ランダムなメールアドレスを生成
- `faker.uuid()` - UUIDを生成
- `faker.password(lower, upper, numeric, special, space, num)` - パスワードを生成
- その他多数のFaker関数が利用可能

### その他の関数
- `input(prompt, default)` - ユーザー入力を取得

## ランナー（Runners）

### HTTPランナー
```yaml
runners:
  req: https://example.com
steps:
  - req:
      /users:
        post:
          headers:
            Authorization: 'Bearer xxxxx'
          body:
            application/json:
              username: alice
              password: passw0rd
    test: current.res.status == 201
```

### gRPCランナー
```yaml
runners:
  greq: grpc://grpc.example.com:80
steps:
  - greq:
      grpctest.GrpcTestService/Hello:
        headers:
          authentication: tokenhello
        message:
          name: alice
          num: 3
```

### DBランナー
対応データベース：
- PostgreSQL（`postgres://`, `pg://`）
- MySQL（`mysql://`, `my://`）
- SQLite3（`sqlite://`, `sq://`）
- Cloud Spanner（`spanner://`, `sp://`）

```yaml
runners:
  db: postgres://dbuser:dbpass@hostname:5432/dbname
steps:
  - db:
      query: SELECT * FROM users;
```

### CDPランナー（Chrome DevTools Protocol）
```yaml
runners:
  cc: chrome://new
steps:
  - cc:
      actions:
        - navigate: https://pkg.go.dev/time
        - click: 'body > header > div.go-Header-inner > nav > div > ul > li:nth-child(2) > a'
        - waitVisible: 'body > footer'
        - text: 'h1'
  - test: previous.text == 'Install the latest version of Go'
```

### SSHランナー
```yaml
runners:
  ssh: ssh://username@hostname:22
steps:
  - ssh:
      command: ls -la
```

## 高度な機能

### ループ処理
```yaml
# シンプルなループ
loop: 10
steps:
  - req:
      /api/test:
        get:
          body: null

# リトライメカニズム
loop:
  count: 10
  until: 'outcome == "success"'
  minInterval: 0.5
  maxInterval: 10
```

### 条件付き実行
```yaml
steps:
  login:
    if: 'len(vars.token) == 0'  # tokenが設定されていない場合のみ実行
    req:
      /login:
        post:
          body:
```

### インクルード機能
```yaml
steps:
  - include:
      path: ./common/login.yml
      vars:
        username: alice
```

### 並行実行制御
```yaml
concurrency: use-shared-db  # 同じキーを持つrunbookは同時に1つしか実行されない
```

### 依存関係の定義
```yaml
needs:
  prebook: path/to/prebook.yml
  prebook2: path/to/prebook2.yml
```

## Go言語でのテストヘルパーとしての使用

```go
func TestRouter(t *testing.T) {
    ctx := context.Background()
    dsn := "username:password@tcp(localhost:3306)/testdb"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    
    ts := httptest.NewServer(NewRouter(db))
    t.Cleanup(func() {
        ts.Close()
        db.Close()
    })
    
    opts := []runn.Option{
        runn.T(t),
        runn.Runner("req", ts.URL),
        runn.DBRunner("db", db),
    }
    
    o, err := runn.Load("testdata/books/**/*.yml", opts...)
    if err != nil {
        t.Fatal(err)
    }
    
    if err := o.RunN(ctx); err != nil {
        t.Fatal(err)
    }
}
```

## まとめ

runnは、HTTPリクエスト、gRPC、データベースクエリ、ブラウザ操作、SSH実行など、多様なプロトコルをサポートする統合的なテスト・自動化ツールです。YAML形式でシナリオを記述し、強力な式評価エンジンとビルトイン関数により、複雑なテストシナリオも簡潔に表現できます。

シングルバイナリで提供されるため、CI/CD環境での利用も容易で、Go言語のテストヘルパーとしても活用できる汎用性の高いツールです。