# 第2章：シナリオ記述編

## Runbook（シナリオファイル）とは

runnでは、テストシナリオをYAML形式のファイルで記述します。このファイルを**Runbook**と呼びます。Runbookは、実行すべきステップを順番に記述した「台本」のようなものです。

## Runbookの基本構造

Runbookは大きく分けて以下のセクションから構成されます：

```yaml
desc: シナリオの説明          # 必須: このシナリオが何をテストするか
labels:                      # オプション: シナリオの分類ラベル
  - api
  - user
runners:                     # 必須: 使用するランナーの定義
  req: https://api.example.com
vars:                        # オプション: シナリオ全体で使用する変数
  baseURL: https://api.example.com
  timeout: 30
steps:                       # 必須: 実行するステップの配列
  - req:
      /users:
        get:
    test: current.res.status == 200
```

### 主要なセクション

1. **desc**: シナリオの説明（必須）
2. **labels**: シナリオの分類用ラベル
3. **runners**: 使用するランナー（HTTP、DB、gRPCなど）の定義
4. **vars**: シナリオ全体で使用する変数の定義
5. **steps**: 実行するステップの配列

## ステップの記述方法：リスト形式とマップ形式

runnでは、ステップを**リスト形式**と**マップ形式**の2つの方法で記述できます。どちらを使うかは、シナリオの複雑さや好みによって選択できます。

### リスト形式

インデックスベースでステップを参照します。シンプルで直感的です。

```yaml
desc: ユーザー作成とログインのテスト（リスト形式）
runners:
  req: https://api.example.com
steps:
  - req:                              # steps[0]
      /users:
        post:
          body:
            application/json:
              name: "Alice"
              email: "alice@example.com"
    test: steps[0].res.status == 201
  
  - req:                              # steps[1]
      /login:
        post:
          body:
            application/json:
              email: "{{ steps[0].res.body.email }}"
              password: "password123"
    test: |
      steps[1].res.status == 200 &&
      steps[1].res.body.token != null
```

### マップ形式

名前付きステップで、可読性が高く複雑なシナリオに適しています。

```yaml
desc: ユーザー作成とログインのテスト（マップ形式）
runners:
  req: https://api.example.com
steps:
  create_user:                        # 名前付きステップ
    req:
      /users:
        post:
          body:
            application/json:
              name: "Alice"
              email: "alice@example.com"
    test: steps.create_user.res.status == 201
  
  login_user:                         # 名前付きステップ
    req:
      /login:
        post:
          body:
            application/json:
              email: "{{ steps.create_user.res.body.email }}"
              password: "password123"
    test: |
      steps.login_user.res.status == 200 &&
      steps.login_user.res.body.token != null
```

### どちらを使うべきか？

- **リスト形式**: シンプルなシナリオ、ステップ数が少ない場合
- **マップ形式**: 複雑なシナリオ、ステップ間の依存関係が多い場合、可読性を重視する場合

## 変数の定義と参照

### varsセクションでの変数定義

```yaml
vars:
  # 静的な値
  apiVersion: v1
  timeout: 30
  
  # 環境変数から取得
  apiKey: ${API_KEY}
  environment: ${ENV:-development}  # デフォルト値付き
  
  # 複雑なデータ構造
  testUser:
    name: "Test User"
    email: "test@example.com"
    roles:
      - admin
      - user
```

### 変数の参照方法

```yaml
steps:
  - req:
      /{{ vars.apiVersion }}/users:  # パス内での変数展開
        get:
          headers:
            X-API-Key: "{{ vars.apiKey }}"
            X-Timeout: "{{ vars.timeout }}"
    test: |
      current.res.status == 200
```

## ステップの詳細な記述

### HTTPリクエストの完全な例

```yaml
steps:
  api_call:
    desc: ユーザー情報を更新する    # ステップの説明
    if: vars.environment == "test"  # 条件付き実行
    req:
      /users/{{ vars.userId }}:
        put:
          headers:
            Content-Type: application/json
            Authorization: "Bearer {{ vars.token }}"
          body:
            application/json:
              name: "Updated Name"
              email: "updated@example.com"
          timeout: 10s              # タイムアウト設定
    test: |                        # テストアサーション
      current.res.status == 200 &&
      current.res.body.name == "Updated Name"
    dump:                          # デバッグ用の値出力
      status: current.res.status
      body: current.res.body
```

### データベースクエリの例

```yaml
runners:
  db: postgres://user:pass@localhost:5432/testdb
steps:
  check_user:
    db:
      query: |
        SELECT id, name, email, created_at
        FROM users
        WHERE email = $1
      params:
        - "{{ vars.testEmail }}"
    test: |
      len(steps.check_user.rows) == 1 &&
      steps.check_user.rows[0].name == "Alice"
```

## 実践的なシナリオ例

### 例1: RESTful APIのCRUD操作

```yaml
desc: ブログ記事のCRUD操作をテスト
runners:
  api: https://blog-api.example.com
vars:
  authorId: "author-123"
steps:
  # 1. 記事を作成
  create_post:
    req:
      /posts:
        post:
          body:
            application/json:
              title: "テスト記事"
              content: "これはテスト記事です"
              authorId: "{{ vars.authorId }}"
    test: |
      steps.create_post.res.status == 201 &&
      steps.create_post.res.body.id != null

  # 2. 作成した記事を取得
  get_post:
    req:
      /posts/{{ steps.create_post.res.body.id }}:
        get:
    test: |
      steps.get_post.res.status == 200 &&
      steps.get_post.res.body.title == "テスト記事"

  # 3. 記事を更新
  update_post:
    req:
      /posts/{{ steps.create_post.res.body.id }}:
        put:
          body:
            application/json:
              title: "更新されたテスト記事"
              content: "内容も更新しました"
    test: steps.update_post.res.status == 200

  # 4. 更新を確認
  verify_update:
    req:
      /posts/{{ steps.create_post.res.body.id }}:
        get:
    test: |
      steps.verify_update.res.body.title == "更新されたテスト記事"

  # 5. 記事を削除
  delete_post:
    req:
      /posts/{{ steps.create_post.res.body.id }}:
        delete:
    test: steps.delete_post.res.status == 204

  # 6. 削除を確認
  verify_delete:
    req:
      /posts/{{ steps.create_post.res.body.id }}:
        get:
    test: steps.verify_delete.res.status == 404
```

### 例2: 認証フローのテスト

```yaml
desc: JWT認証フローの完全なテスト
runners:
  auth: https://auth.example.com
  api: https://api.example.com
vars:
  testUser:
    email: "test@example.com"
    password: "Test123!@#"
steps:
  # ユーザー登録
  register:
    req:
      auth:///register:
        post:
          body:
            application/json:
              email: "{{ vars.testUser.email }}"
              password: "{{ vars.testUser.password }}"
              name: "Test User"
    test: |
      steps.register.res.status == 201

  # ログイン
  login:
    req:
      auth:///login:
        post:
          body:
            application/json:
              email: "{{ vars.testUser.email }}"
              password: "{{ vars.testUser.password }}"
    test: |
      steps.login.res.status == 200 &&
      steps.login.res.body.accessToken != null &&
      steps.login.res.body.refreshToken != null

  # 認証が必要なAPIにアクセス
  access_protected:
    req:
      api:///profile:
        get:
          headers:
            Authorization: "Bearer {{ steps.login.res.body.accessToken }}"
    test: |
      steps.access_protected.res.status == 200 &&
      steps.access_protected.res.body.email == vars.testUser.email

  # トークンをリフレッシュ
  refresh:
    req:
      auth:///refresh:
        post:
          body:
            application/json:
              refreshToken: "{{ steps.login.res.body.refreshToken }}"
    test: |
      steps.refresh.res.status == 200 &&
      steps.refresh.res.body.accessToken != null

  # 新しいトークンでアクセス
  access_with_new_token:
    req:
      api:///profile:
        get:
          headers:
            Authorization: "Bearer {{ steps.refresh.res.body.accessToken }}"
    test: steps.access_with_new_token.res.status == 200
```

## YAMLの記述Tips

### 1. 複数行文字列の記述

```yaml
steps:
  - desc: |
      これは複数行の
      説明文です
    req:
      /test:
        post:
          body:
            text/plain: |
              複数行の
              テキストデータ
```

### 2. アンカーとエイリアスの活用

```yaml
# 共通のヘッダーを定義
commonHeaders: &headers
  Content-Type: application/json
  X-API-Version: "1.0"

steps:
  - req:
      /users:
        get:
          headers:
            <<: *headers          # 共通ヘッダーを使用
            Authorization: "Bearer token123"
```

### 3. 環境別の設定

```yaml
vars:
  baseURL: ${BASE_URL:-https://api.example.com}
  apiKey: ${API_KEY}
  environment: ${ENV:-development}
  
  # 環境別の設定をマップで管理
  config:
    development:
      timeout: 60
      retries: 3
    production:
      timeout: 30
      retries: 1
```

## まとめ

この章では、runnのシナリオ記述について以下を学びました：

1. **Runbookの基本構造**: desc、runners、vars、stepsの各セクション
2. **2つの記述形式**: リスト形式とマップ形式、それぞれの利点
3. **変数の活用**: 定義方法と参照方法、環境変数の利用
4. **実践的なシナリオ**: CRUD操作、認証フローなどの実例

次章では、runnの強力な式評価エンジンについて詳しく見ていきます。

[第3章：Expression文法編へ →](chapter03.md)