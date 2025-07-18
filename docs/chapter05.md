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
desc: HTTPランナーの基本設定
runners:
  api: https://api.example.com/v1
  # 複数のエンドポイントを定義可能
  auth: https://auth.example.com
  webhook: http://localhost:8080/webhook
```

### リクエストメソッドとパラメータ

```yaml
steps:
  # GET リクエスト
  get_users:
    req:
      api:///users:  # runners.apiのベースURLを使用
        get:
          query:
            page: 1
            limit: 10
            sort: created_at
          headers:
            Accept: application/json
            User-Agent: runn/test
    test: current.res.status == 200

  # POST リクエスト
  create_user:
    req:
      api:///users:
        post:
          headers:
            Content-Type: application/json
            Authorization: "Bearer {{ vars.token }}"
          body:
            application/json:
              name: "Alice"
              email: "alice@example.com"
              role: "user"
    test: |
      current.res.status == 201 &&
      current.res.body.id != null

  # PUT リクエスト（更新）
  update_user:
    req:
      api:///users/{{ steps.create_user.res.body.id }}:
        put:
          body:
            application/json:
              name: "Alice Smith"
              email: "alice.smith@example.com"
    test: current.res.status == 200

  # PATCH リクエスト（部分更新）
  patch_user:
    req:
      api:///users/{{ steps.create_user.res.body.id }}:
        patch:
          body:
            application/json:
              role: "admin"
    test: current.res.status == 200

  # DELETE リクエスト
  delete_user:
    req:
      api:///users/{{ steps.create_user.res.body.id }}:
        delete:
    test: current.res.status == 204
```

### 様々なボディ形式

```yaml
steps:
  # JSON形式
  json_request:
    req:
      api:///data:
        post:
          body:
            application/json:
              key: "value"
              nested:
                array: [1, 2, 3]

  # フォームデータ
  form_request:
    req:
      api:///form:
        post:
          body:
            application/x-www-form-urlencoded:
              username: alice
              password: secret123

  # マルチパートフォーム
  multipart_request:
    req:
      api:///upload:
        post:
          body:
            multipart/form-data:
              file: "@./testdata/sample.txt"
              description: "Test file upload"

  # プレーンテキスト
  text_request:
    req:
      api:///webhook:
        post:
          body:
            text/plain: |
              This is a plain text message
              with multiple lines

  # XML形式
  xml_request:
    req:
      api:///soap:
        post:
          body:
            application/xml: |
              <?xml version="1.0" encoding="UTF-8"?>
              <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
                <soap:Body>
                  <GetUser>
                    <UserId>123</UserId>
                  </GetUser>
                </soap:Body>
              </soap:Envelope>
```

### 認証の実装

```yaml
steps:
  # Basic認証
  basic_auth:
    req:
      api:///protected:
        get:
          headers:
            Authorization: "Basic {{ toBase64(vars.username + ':' + vars.password) }}"

  # Bearer Token認証
  bearer_auth:
    req:
      api:///protected:
        get:
          headers:
            Authorization: "Bearer {{ vars.access_token }}"

  # API Key認証
  api_key_auth:
    req:
      api:///protected:
        get:
          headers:
            X-API-Key: "{{ vars.api_key }}"
          query:
            api_key: "{{ vars.api_key }}"  # クエリパラメータとしても可能

  # OAuth 2.0フロー
  oauth_login:
    req:
      auth:///oauth/token:
        post:
          body:
            application/x-www-form-urlencoded:
              grant_type: client_credentials
              client_id: "{{ vars.client_id }}"
              client_secret: "{{ vars.client_secret }}"
              scope: "read write"
    test: current.res.status == 200

  use_oauth_token:
    req:
      api:///protected:
        get:
          headers:
            Authorization: "Bearer {{ steps.oauth_login.res.body.access_token }}"
    test: current.res.status == 200
```

### レスポンスの詳細な検証

```yaml
steps:
  detailed_validation:
    req:
      api:///users:
        get:
    test: |
      # ステータスコード
      current.res.status == 200 &&
      
      # ヘッダーの検証
      current.res.headers["Content-Type"] contains "application/json" &&
      current.res.headers["X-Rate-Limit-Remaining"] != null &&
      
      # ボディの構造検証
      "data" in current.res.body &&
      "pagination" in current.res.body &&
      
      # データの詳細検証
      len(current.res.body.data) > 0 &&
      all(current.res.body.data, {
        "id" in . &&
        "name" in . &&
        "email" in . &&
        .id > 0 &&
        contains(.email, "@")
      }) &&
      
      # ページネーション検証
      current.res.body.pagination.page >= 1 &&
      current.res.body.pagination.total >= len(current.res.body.data)
```

### GraphQL API の操作

```yaml
runners:
  graphql: https://api.github.com/graphql

steps:
  graphql_query:
    req:
      graphql:///:
        post:
          headers:
            Authorization: "Bearer {{ env.GITHUB_TOKEN }}"
            Content-Type: application/json
          body:
            application/json:
              query: |
                query {
                  viewer {
                    login
                    name
                    email
                    repositories(first: 5) {
                      nodes {
                        name
                        description
                        stargazerCount
                      }
                    }
                  }
                }
    test: |
      current.res.status == 200 &&
      current.res.body.data.viewer.login != null &&
      len(current.res.body.data.viewer.repositories.nodes) <= 5

  graphql_mutation:
    req:
      graphql:///:
        post:
          headers:
            Authorization: "Bearer {{ env.GITHUB_TOKEN }}"
          body:
            application/json:
              query: |
                mutation($input: AddStarInput!) {
                  addStar(input: $input) {
                    starrable {
                      stargazerCount
                    }
                  }
                }
              variables:
                input:
                  starrableId: "{{ vars.repository_id }}"
    test: current.res.status == 200
```

## gRPCランナー

### 基本的な設定

```yaml
desc: gRPCランナーの使用例
runners:
  grpc: grpc://localhost:50051
  # TLS接続の場合
  secure_grpc: grpcs://api.example.com:443

steps:
  # Unary RPC
  unary_call:
    greq:
      grpc:///helloworld.Greeter/SayHello:
        headers:
          authorization: "Bearer {{ vars.token }}"
        message:
          name: "World"
    test: |
      current.res.status == 0 &&  # gRPCのOKステータス
      current.res.message.message == "Hello World"

  # Server Streaming RPC
  server_streaming:
    greq:
      grpc:///example.StreamService/GetStream:
        message:
          count: 5
    test: |
      current.res.status == 0 &&
      len(current.res.messages) == 5

  # Client Streaming RPC
  client_streaming:
    greq:
      grpc:///example.StreamService/PutStream:
        messages:
          - data: "message1"
          - data: "message2"
          - data: "message3"
    test: |
      current.res.status == 0 &&
      current.res.message.count == 3

  # Bidirectional Streaming RPC
  bidirectional_streaming:
    greq:
      grpc:///example.ChatService/Chat:
        messages:
          - user: "Alice"
            message: "Hello"
          - user: "Alice"
            message: "How are you?"
    test: |
      current.res.status == 0 &&
      len(current.res.messages) >= 2
```

### プロトコルバッファの動的読み込み

```yaml
runners:
  grpc:
    addr: grpc://localhost:50051
    # .protoファイルを指定
    protos:
      - ./proto/service.proto
      - ./proto/common.proto
    # インポートパスを指定
    import_paths:
      - ./proto
      - ./third_party/proto

steps:
  dynamic_proto:
    greq:
      grpc:///myservice.MyService/GetData:
        message:
          id: 123
          filters:
            - field: "name"
              value: "test"
```

## データベースランナー

### 対応データベース

```yaml
runners:
  # PostgreSQL
  postgres: postgres://user:password@localhost:5432/testdb?sslmode=disable
  
  # MySQL
  mysql: mysql://user:password@localhost:3306/testdb
  
  # SQLite
  sqlite: sqlite:///path/to/database.db
  
  # Cloud Spanner
  spanner: spanner://projects/my-project/instances/my-instance/databases/my-db
```

### 基本的なクエリ操作

```yaml
steps:
  # SELECT クエリ
  select_users:
    db:
      postgres:///
        query: |
          SELECT id, name, email, created_at
          FROM users
          WHERE active = true
          ORDER BY created_at DESC
          LIMIT 10
    test: |
      len(current.rows) <= 10 &&
      all(current.rows, {.id > 0})

  # パラメータ付きクエリ
  select_user_by_email:
    db:
      postgres:///
        query: |
          SELECT id, name, email
          FROM users
          WHERE email = $1
        params:
          - "{{ vars.test_email }}"
    test: |
      len(current.rows) == 1 &&
      current.rows[0].email == vars.test_email

  # INSERT クエリ
  insert_user:
    db:
      postgres:///
        query: |
          INSERT INTO users (name, email, password_hash)
          VALUES ($1, $2, $3)
          RETURNING id, created_at
        params:
          - "{{ faker.name() }}"
          - "{{ faker.email() }}"
          - "{{ toBase64(faker.password(true, true, true, false, false, 12)) }}"
    test: |
      current.rows[0].id > 0 &&
      current.rows[0].created_at != null

  # UPDATE クエリ
  update_user:
    db:
      postgres:///
        query: |
          UPDATE users
          SET name = $1, updated_at = NOW()
          WHERE id = $2
          RETURNING name, updated_at
        params:
          - "Updated Name"
          - "{{ steps.insert_user.rows[0].id }}"
    test: |
      current.rows[0].name == "Updated Name" &&
      current.rows[0].updated_at != null

  # DELETE クエリ
  delete_user:
    db:
      postgres:///
        query: |
          DELETE FROM users
          WHERE id = $1
        params:
          - "{{ steps.insert_user.rows[0].id }}"
    test: current.rowsAffected == 1
```

### トランザクション処理

```yaml
steps:
  # トランザクション開始
  begin_transaction:
    db:
      postgres:///
        query: BEGIN

  # 複数の操作を実行
  insert_order:
    db:
      postgres:///
        query: |
          INSERT INTO orders (user_id, total_amount)
          VALUES ($1, $2)
          RETURNING id
        params:
          - "{{ vars.user_id }}"
          - "{{ vars.total_amount }}"

  insert_order_items:
    loop:
      count: len(vars.order_items)
    db:
      postgres:///
        query: |
          INSERT INTO order_items (order_id, product_id, quantity, price)
          VALUES ($1, $2, $3, $4)
        params:
          - "{{ steps.insert_order.rows[0].id }}"
          - "{{ vars.order_items[i].product_id }}"
          - "{{ vars.order_items[i].quantity }}"
          - "{{ vars.order_items[i].price }}"

  # トランザクションコミット
  commit_transaction:
    db:
      postgres:///
        query: COMMIT
    test: current.rowsAffected == 0  # COMMITは行に影響しない

  # エラー時のロールバック例
  rollback_on_error:
    if: vars.should_rollback
    db:
      postgres:///
        query: ROLLBACK
```

### 複雑なデータ検証

```yaml
steps:
  data_integrity_check:
    db:
      postgres:///
        query: |
          SELECT 
            u.id,
            u.name,
            u.email,
            COUNT(o.id) as order_count,
            COALESCE(SUM(o.total_amount), 0) as total_spent
          FROM users u
          LEFT JOIN orders o ON u.id = o.user_id
          WHERE u.active = true
          GROUP BY u.id, u.name, u.email
          HAVING COUNT(o.id) > 0
          ORDER BY total_spent DESC
    test: |
      len(current.rows) > 0 &&
      all(current.rows, {
        .id > 0 &&
        .name != "" &&
        contains(.email, "@") &&
        .order_count > 0 &&
        .total_spent > 0
      }) &&
      # 最も多く購入したユーザーが先頭にいることを確認
      current.rows[0].total_spent >= current.rows[-1].total_spent
```

## CDPランナー（ブラウザ自動化）

### 基本的な設定

```yaml
desc: ブラウザ自動化テスト
runners:
  browser: chrome://new  # 新しいChromeインスタンスを起動
  # 既存のブラウザに接続する場合
  # browser: chrome://localhost:9222

steps:
  # ページナビゲーション
  navigate_to_page:
    cdp:
      browser:///
        actions:
          - navigate: https://example.com
          - waitVisible: 'h1'
    test: current.url == "https://example.com/"

  # 要素の操作
  interact_with_elements:
    cdp:
      browser:///
        actions:
          # テキスト入力
          - type:
              selector: 'input[name="username"]'
              text: "testuser"
          
          # クリック
          - click: 'button[type="submit"]'
          
          # 要素が表示されるまで待機
          - waitVisible: '.success-message'
          
          # スクリーンショット撮影
          - screenshot: './screenshots/login-success.png'
    test: |
      current.screenshot != null &&
      current.url contains "/dashboard"

  # フォーム操作
  form_interaction:
    cdp:
      browser:///
        actions:
          - navigate: https://example.com/form
          
          # 複数の入力フィールド
          - type:
              selector: '#name'
              text: "{{ faker.name() }}"
          - type:
              selector: '#email'
              text: "{{ faker.email() }}"
          
          # セレクトボックス
          - select:
              selector: '#country'
              value: 'JP'
          
          # チェックボックス
          - check: '#agree-terms'
          
          # ラジオボタン
          - click: 'input[name="gender"][value="other"]'
          
          # フォーム送信
          - click: 'button[type="submit"]'
          
          # 結果の確認
          - waitVisible: '.form-success'
          - text: '.form-success'
    test: |
      current.text contains "successfully" ||
      current.text contains "thank you"
```

### 高度なブラウザ操作

```yaml
steps:
  advanced_browser_actions:
    cdp:
      browser:///
        actions:
          - navigate: https://example.com/app
          
          # JavaScript実行
          - evaluate: |
              window.scrollTo(0, document.body.scrollHeight);
              return document.title;
          
          # 要素のテキスト取得
          - text: 'h1.main-title'
          
          # 要素の属性取得
          - attribute:
              selector: 'img.logo'
              name: 'src'
          
          # 複数要素の取得
          - textAll: '.product-name'
          
          # 要素の存在確認
          - exists: '.error-message'
          
          # 要素が非表示になるまで待機
          - waitNotVisible: '.loading-spinner'
          
          # カスタム待機条件
          - wait: |
              document.querySelectorAll('.product-item').length >= 10
          
          # ページのPDF出力
          - pdf: './output/page.pdf'
    
    test: |
      current.text != "" &&
      current.attribute != "" &&
      len(current.textAll) > 0 &&
      current.exists == false  # エラーメッセージが存在しないことを確認
```

### SPAアプリケーションのテスト

```yaml
steps:
  spa_testing:
    cdp:
      browser:///
        actions:
          - navigate: https://spa-app.example.com
          
          # 初期ローディング完了まで待機
          - wait: |
              window.app && window.app.initialized === true
          
          # ルーティングのテスト
          - click: 'a[href="/products"]'
          - waitVisible: '.product-list'
          
          # 動的コンテンツの読み込み待機
          - wait: |
              document.querySelectorAll('.product-item').length > 0
          
          # 検索機能のテスト
          - type:
              selector: '.search-input'
              text: 'laptop'
          - click: '.search-button'
          
          # 検索結果の待機
          - wait: |
              document.querySelector('.search-results') &&
              !document.querySelector('.loading')
          
          - textAll: '.product-item .product-name'
    
    test: |
      all(current.textAll, {
        lower(.) contains "laptop"
      })
```

## SSHランナー

### 基本的な設定

```yaml
desc: SSH経由でのリモート操作
runners:
  server: ssh://user@example.com:22
  # 秘密鍵を使用する場合
  secure_server:
    type: ssh
    addr: user@secure.example.com:22
    key: /path/to/private_key
    passphrase: "{{ env.SSH_PASSPHRASE }}"

steps:
  # 基本的なコマンド実行
  basic_command:
    ssh:
      server:///
        command: ls -la /home/user
    test: |
      current.exit_code == 0 &&
      current.stdout contains "total"

  # 複数コマンドの実行
  multiple_commands:
    ssh:
      server:///
        command: |
          cd /var/log
          ls -la *.log | head -5
          df -h
    test: current.exit_code == 0

  # ファイル操作
  file_operations:
    ssh:
      server:///
        command: |
          echo "Test content" > /tmp/test.txt
          cat /tmp/test.txt
          rm /tmp/test.txt
    test: |
      current.exit_code == 0 &&
      current.stdout contains "Test content"

  # システム情報の取得
  system_info:
    ssh:
      server:///
        command: |
          uname -a
          uptime
          free -m
          ps aux | head -10
    test: current.exit_code == 0
    dump:
      system_output: current.stdout
```

### サーバー監視とヘルスチェック

```yaml
steps:
  health_check:
    ssh:
      server:///
        command: |
          # サービスの状態確認
          systemctl is-active nginx
          systemctl is-active postgresql
          
          # ポートの確認
          netstat -tlnp | grep :80
          netstat -tlnp | grep :5432
          
          # ディスク使用量
          df -h | grep -E '(Filesystem|/dev/)'
          
          # メモリ使用量
          free -m
          
          # CPU負荷
          uptime
    test: |
      current.exit_code == 0 &&
      current.stdout contains "active" &&
      !(current.stdout contains "failed")
    
    dump:
      health_status: |
        {
          "services_active": current.stdout contains "active",
          "ports_open": current.stdout contains ":80" && current.stdout contains ":5432",
          "timestamp": time.now()
        }
```

## Execランナー（ローカルコマンド実行）

### 基本的な使用方法

```yaml
desc: ローカルコマンドの実行
steps:
  # 基本的なコマンド実行
  basic_exec:
    exec:
      command: echo "Hello, World!"
    test: |
      current.exit_code == 0 &&
      current.stdout == "Hello, World!\n"

  # 環境変数を設定してコマンド実行
  exec_with_env:
    exec:
      command: env | grep TEST_VAR
      env:
        TEST_VAR: "test_value"
        PATH: "{{ env.PATH }}"
    test: |
      current.exit_code == 0 &&
      current.stdout contains "TEST_VAR=test_value"

  # 作業ディレクトリを指定
  exec_with_workdir:
    exec:
      command: pwd
      dir: /tmp
    test: current.stdout contains "/tmp"

  # 複雑なシェルコマンド
  complex_shell:
    exec:
      command: |
        for i in {1..5}; do
          echo "Count: $i"
        done | grep "Count: [35]"
    test: |
      current.exit_code == 0 &&
      current.stdout contains "Count: 3" &&
      current.stdout contains "Count: 5"
```

### ファイル操作とテスト

```yaml
steps:
  # テストファイルの作成
  create_test_file:
    exec:
      command: |
        cat > /tmp/test_data.json << 'EOF'
        {
          "users": [
            {"id": 1, "name": "Alice"},
            {"id": 2, "name": "Bob"}
          ]
        }
        EOF
    test: current.exit_code == 0

  # ファイル内容の検証
  validate_file:
    exec:
      command: cat /tmp/test_data.json
    test: |
      current.exit_code == 0 &&
      fromJSON(current.stdout).users[0].name == "Alice"

  # ファイルのクリーンアップ
  cleanup:
    exec:
      command: rm -f /tmp/test_data.json
    test: current.exit_code == 0
```

## ランナーの組み合わせ

### マルチプロトコルテスト

```yaml
desc: 複数のランナーを組み合わせたE2Eテスト
runners:
  api: https://api.example.com
  db: postgres://user:pass@localhost:5432/testdb
  browser: chrome://new

vars:
  test_user:
    name: "{{ faker.name() }}"
    email: "{{ faker.email() }}"

steps:
  # 1. データベースにテストデータを準備
  setup_test_data:
    db:
      db:///
        query: |
          INSERT INTO users (name, email, active)
          VALUES ($1, $2, true)
          RETURNING id
        params:
          - "{{ vars.test_user.name }}"
          - "{{ vars.test_user.email }}"
    test: current.rows[0].id > 0

  # 2. APIでユーザー情報を取得
  api_get_user:
    req:
      api:///users/{{ steps.setup_test_data.rows[0].id }}:
        get:
    test: |
      current.res.status == 200 &&
      current.res.body.name == vars.test_user.name

  # 3. ブラウザでユーザー管理画面を確認
  browser_verify_user:
    cdp:
      browser:///
        actions:
          - navigate: https://admin.example.com/users
          - waitVisible: '.user-list'
          - type:
              selector: '.search-input'
              text: "{{ vars.test_user.email }}"
          - click: '.search-button'
          - waitVisible: '.user-row'
          - text: '.user-row .user-name'
    test: current.text == vars.test_user.name

  # 4. データベースからテストデータを削除
  cleanup_test_data:
    db:
      db:///
        query: |
          DELETE FROM users
          WHERE id = $1
        params:
          - "{{ steps.setup_test_data.rows[0].id }}"
    test: current.rowsAffected == 1
```

### 障害テストシナリオ

```yaml
desc: 障害時の動作確認
steps:
  # 正常時のレスポンス確認
  normal_request:
    req:
      api:///health:
        get:
    test: current.res.status == 200

  # サーバーに負荷をかける
  load_test:
    loop:
      count: 100
      concurrent: 10
    req:
      api:///heavy-operation:
        post:
          body:
            application/json:
              data: "{{ faker.randomString(1000) }}"
    test: current.res.status in [200, 202, 429]  # 成功またはレート制限

  # 障害後の復旧確認
  recovery_check:
    loop:
      count: 5
      until: current.res.status == 200
      minInterval: 1
      maxInterval: 5
    req:
      api:///health:
        get:
    test: current.res.status == 200
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