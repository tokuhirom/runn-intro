---
layout: default
title: 第6章：高度な機能編
---

# 第6章：高度な機能編

この章では、runnの高度な機能について詳しく解説します。これらの機能を使いこなすことで、より複雑で実用的なテストシナリオを構築できます。

## ループ処理

### 基本的なループ

```yaml
desc: 基本的なループ処理
steps:
  # 単純な回数指定ループ
  simple_loop:
    loop: 5
    req:
      /api/ping:
        get:
    test: current.res.status == 200
    dump:
      iteration: i  # ループインデックス（0から開始）

  # 配列の各要素に対するループ
  array_loop:
    loop:
      count: len(vars.test_users)
    req:
      /users:
        post:
          body:
            application/json:
              name: "{{ vars.test_users[i].name }}"
              email: "{{ vars.test_users[i].email }}"
    test: current.res.status == 201
```

### 条件付きループ（リトライ機能）

```yaml
steps:
  # 成功するまでリトライ
  retry_until_success:
    loop:
      count: 10  # 最大10回
      until: current.res.status == 200  # 成功条件
      minInterval: 1  # 最小間隔（秒）
      maxInterval: 5  # 最大間隔（秒）
    req:
      /unstable-endpoint:
        get:
    test: current.res.status == 200

  # 複雑な条件でのリトライ
  complex_retry:
    loop:
      count: 5
      until: |
        current.res.status == 200 &&
        current.res.body.status == "ready" &&
        len(current.res.body.items) > 0
      minInterval: 2
      maxInterval: 10
    req:
      /async-operation/{{ vars.operation_id }}/status:
        get:
    test: |
      current.res.body.status == "ready"

  # エラー条件でのループ終了
  stop_on_error:
    loop:
      count: 100
      until: current.res.status >= 400  # エラーが発生したら停止
    req:
      /batch-process:
        post:
          body:
            application/json:
              batch_id: "{{ i }}"
    test: |
      current.res.status < 400 ||  # 成功
      (current.res.status >= 400 && i > 0)  # エラーだが少なくとも1回は成功
```

### 動的なループ制御

```yaml
vars:
  page_size: 10
  max_pages: 100

steps:
  # ページネーションを使った全データ取得
  paginated_fetch:
    loop:
      count: vars.max_pages
      until: len(current.res.body.data) < vars.page_size  # 最後のページに到達
    req:
      /users:
        get:
          query:
            page: "{{ i + 1 }}"
            limit: "{{ vars.page_size }}"
    test: current.res.status == 200
    dump:
      page_number: i + 1
      items_count: len(current.res.body.data)
      total_fetched: |
        sum(map(steps.paginated_fetch, {len(.res.body.data)}))

  # 条件に基づく動的ループ
  conditional_processing:
    loop:
      count: len(vars.items_to_process)
    if: vars.items_to_process[i].needs_processing
    req:
      /process:
        post:
          body:
            application/json: "{{ vars.items_to_process[i] }}"
    test: current.res.status == 200
```

## 条件付き実行

### 基本的な条件分岐

```yaml
steps:
  # 環境による条件分岐
  environment_specific:
    if: env.ENVIRONMENT == "production"
    req:
      /production-only-endpoint:
        get:
    test: current.res.status == 200

  # 変数による条件分岐
  feature_flag_check:
    if: vars.enable_new_feature
    req:
      /new-feature:
        get:
    test: current.res.status == 200

  # 前のステップの結果による条件分岐
  conditional_on_previous:
    if: steps.user_creation.res.status == 201
    req:
      /users/{{ steps.user_creation.res.body.id }}/activate:
        post:
    test: current.res.status == 200
```

### 複雑な条件式

```yaml
steps:
  complex_conditions:
    if: |
      (env.ENVIRONMENT == "test" || env.ENVIRONMENT == "staging") &&
      vars.user_role == "admin" &&
      len(vars.test_data) > 0
    req:
      /admin/test-data:
        post:
          body:
            application/json: "{{ vars.test_data }}"
    test: current.res.status == 201

  # 複数条件の組み合わせ
  multi_condition_check:
    if: |
      vars.api_version >= 2 &&
      (vars.user_permissions contains "write" || vars.user_role == "admin") &&
      time.now() > time.parse(vars.maintenance_end, time.RFC3339)
    req:
      /v2/protected-resource:
        post:
          body:
            application/json:
              action: "create"
              data: "{{ vars.resource_data }}"
    test: current.res.status in [200, 201]
```

### エラーハンドリングと条件分岐

```yaml
steps:
  # エラー時の代替処理
  primary_request:
    req:
      /primary-endpoint:
        get:
    test: true  # エラーでも続行

  fallback_request:
    if: steps.primary_request.res.status != 200
    req:
      /fallback-endpoint:
        get:
    test: current.res.status == 200

  # 成功時の追加処理
  success_processing:
    if: |
      steps.primary_request.res.status == 200 ||
      steps.fallback_request.res.status == 200
    req:
      /process-result:
        post:
          body:
            application/json:
              source: |
                steps.primary_request.res.status == 200 ? "primary" : "fallback"
              data: |
                steps.primary_request.res.status == 200 ? 
                steps.primary_request.res.body : 
                steps.fallback_request.res.body
    test: current.res.status == 200
```

## シナリオのインクルード

### 基本的なインクルード

```yaml
# main.yml
desc: メインシナリオ
vars:
  base_url: https://api.example.com
  user_id: 123

steps:
  # 共通の認証処理をインクルード
  - include:
      path: ./common/auth.yml
      vars:
        username: "{{ vars.test_username }}"
        password: "{{ vars.test_password }}"

  # ユーザー操作をインクルード
  - include:
      path: ./user/user_operations.yml
      vars:
        user_id: "{{ vars.user_id }}"
        auth_token: "{{ steps[0].auth_token }}"

  # クリーンアップ処理をインクルード
  - include:
      path: ./common/cleanup.yml
```

```yaml
# common/auth.yml
desc: 認証処理
steps:
  login:
    req:
      "{{ parent.vars.base_url }}/auth/login":
        post:
          body:
            application/json:
              username: "{{ vars.username }}"
              password: "{{ vars.password }}"
    test: current.res.status == 200
    dump:
      auth_token: current.res.body.token
```

### 動的なインクルード

```yaml
steps:
  # 条件に基づくインクルード
  conditional_include:
    if: vars.test_type == "integration"
    include:
      path: ./integration/full_test.yml
      vars:
        test_data: "{{ vars.integration_data }}"

  # ループ内でのインクルード
  multiple_scenarios:
    loop:
      count: len(vars.test_scenarios)
    include:
      path: "{{ vars.test_scenarios[i].path }}"
      vars:
        scenario_data: "{{ vars.test_scenarios[i].data }}"
        iteration: "{{ i }}"
```

### ネストしたインクルード

```yaml
# level1.yml
desc: レベル1のシナリオ
steps:
  setup:
    include:
      path: ./setup/database.yml
      vars:
        db_name: "test_{{ time.unix(time.now()) }}"

  main_test:
    include:
      path: ./level2.yml
      vars:
        db_name: "{{ steps.setup.db_name }}"
```

```yaml
# level2.yml
desc: レベル2のシナリオ
steps:
  data_preparation:
    include:
      path: ./data/prepare.yml
      vars:
        target_db: "{{ parent.vars.db_name }}"

  execute_test:
    include:
      path: ./level3.yml
      vars:
        prepared_data: "{{ steps.data_preparation.result }}"
```

## 並行実行制御

### 基本的な並行実行

```yaml
desc: 並行実行の制御
# 同時実行数の制限
concurrency: 5

steps:
  # 複数のAPIエンドポイントを並行テスト
  parallel_api_tests:
    loop:
      count: len(vars.api_endpoints)
    req:
      "{{ vars.api_endpoints[i] }}":
        get:
    test: current.res.status == 200
```

### 共有リソースの制御

```yaml
# データベースを使用するテストは同時に1つだけ実行
concurrency: use-database

steps:
  database_test:
    db:
      postgres:///
        query: |
          INSERT INTO test_table (data) VALUES ($1)
        params:
          - "{{ faker.randomString(10) }}"
    test: current.rowsAffected == 1
```

### 複雑な並行制御

```yaml
desc: 複雑な並行実行制御
vars:
  worker_count: 10
  batch_size: 100

steps:
  # ワーカープロセスのシミュレーション
  worker_simulation:
    loop:
      count: vars.worker_count
    include:
      path: ./worker/process_batch.yml
      vars:
        worker_id: "{{ i }}"
        batch_start: "{{ i * vars.batch_size }}"
        batch_end: "{{ (i + 1) * vars.batch_size }}"
        
  # 結果の集約
  aggregate_results:
    dump:
      total_processed: |
        sum(map(steps.worker_simulation, {.processed_count}))
      success_rate: |
        sum(map(steps.worker_simulation, {.success_count})) / 
        sum(map(steps.worker_simulation, {.processed_count}))
```

## 依存関係の定義

### 基本的な依存関係

```yaml
desc: 依存関係のあるテスト
needs:
  setup: ./setup/environment.yml
  data: ./setup/test_data.yml

steps:
  main_test:
    req:
      /api/test:
        get:
          headers:
            Authorization: "Bearer {{ needs.setup.auth_token }}"
    test: current.res.status == 200

  data_validation:
    test: |
      current.res.body.count == needs.data.expected_count
```

### 複雑な依存関係

```yaml
desc: 複雑な依存関係の管理
needs:
  # 基本セットアップ
  base_setup: ./setup/base.yml
  
  # データベースセットアップ（base_setupに依存）
  db_setup:
    path: ./setup/database.yml
    needs:
      - base_setup
  
  # ユーザーデータ準備（db_setupに依存）
  user_data:
    path: ./data/users.yml
    needs:
      - db_setup
  
  # 商品データ準備（db_setupに依存）
  product_data:
    path: ./data/products.yml
    needs:
      - db_setup

steps:
  integration_test:
    req:
      /api/orders:
        post:
          body:
            application/json:
              user_id: "{{ needs.user_data.test_user_id }}"
              product_id: "{{ needs.product_data.test_product_id }}"
              quantity: 1
    test: current.res.status == 201
```

## カスタムランナーの作成

### プラグインランナーの定義

```yaml
desc: カスタムランナーの使用例
runners:
  # カスタムHTTPクライアント
  custom_http:
    type: http
    base_url: https://api.example.com
    default_headers:
      User-Agent: "MyApp/1.0"
      Accept: "application/json"
    timeout: 30s
    retry_count: 3

  # カスタムデータベース接続
  custom_db:
    type: db
    dsn: "{{ env.DATABASE_URL }}"
    max_connections: 10
    connection_timeout: 5s

steps:
  custom_request:
    req:
      custom_http:///users:
        get:
          headers:
            X-Custom-Header: "custom-value"
    test: current.res.status == 200
```

### 外部コマンドランナー

```yaml
runners:
  # 外部ツールの実行
  kubectl:
    type: exec
    command: kubectl
    default_args:
      - --kubeconfig={{ env.KUBECONFIG }}
      - --namespace={{ vars.namespace }}

  # Docker操作
  docker:
    type: exec
    command: docker
    working_dir: "{{ vars.project_root }}"

steps:
  k8s_deployment:
    exec:
      kubectl:///
        args:
          - apply
          - -f
          - deployment.yaml
    test: current.exit_code == 0

  docker_build:
    exec:
      docker:///
        args:
          - build
          - -t
          - "{{ vars.image_name }}:{{ vars.tag }}"
          - .
    test: current.exit_code == 0
```

## 高度なデータ処理

### 動的なテストデータ生成

```yaml
vars:
  # 動的なテストデータ生成
  test_users: |
    map(range(1, 11), {
      "id": .,
      "name": faker.name(),
      "email": faker.email(),
      "age": faker.randomInt(18, 65),
      "department": faker.randomChoice(["IT", "Sales", "Marketing", "HR"]),
      "salary": faker.randomInt(30000, 100000),
      "active": faker.randomBool()
    })

steps:
  bulk_user_creation:
    loop:
      count: len(vars.test_users)
    req:
      /users:
        post:
          body:
            application/json: "{{ vars.test_users[i] }}"
    test: current.res.status == 201

  # 生成されたデータの統計分析
  data_analysis:
    dump:
      total_users: len(vars.test_users)
      active_users: len(filter(vars.test_users, {.active}))
      avg_age: |
        sum(map(vars.test_users, {.age})) / len(vars.test_users)
      departments: |
        groupBy(vars.test_users, {.department})
      salary_stats:
        min: min(map(vars.test_users, {.salary}))
        max: max(map(vars.test_users, {.salary}))
        avg: |
          sum(map(vars.test_users, {.salary})) / len(vars.test_users)
```

### 複雑なデータ変換

```yaml
steps:
  data_transformation:
    req:
      /api/raw-data:
        get:
    
    dump:
      # 生データの変換
      transformed_data: |
        map(current.res.body.items, {
          merge(
            pick(., ["id", "name", "created_at"]),
            {
              "display_name": upper(.name),
              "age_days": time.sub(time.now(), time.parse(.created_at, time.RFC3339)) / time.day,
              "category": .price > 1000 ? "premium" : "standard",
              "tags": split(.tag_string, ","),
              "metadata": fromJSON(.metadata_json ?? "{}")
            }
          )
        })
      
      # 集約データの生成
      summary: |
        let items = current.res.body.items;
        {
          "total_count": len(items),
          "premium_count": len(filter(items, {.price > 1000})),
          "categories": unique(map(items, {.category})),
          "price_ranges": groupBy(items, {
            .price < 100 ? "low" :
            .price < 1000 ? "medium" : "high"
          }),
          "recent_items": filter(items, {
            time.sub(time.now(), time.parse(.created_at, time.RFC3339)) < time.day * 7
          })
        }
```

## エラーハンドリングとデバッグ

### 包括的なエラーハンドリング

```yaml
steps:
  robust_api_call:
    loop:
      count: 3
      until: |
        current.res.status == 200 ||
        current.res.status == 404  # 404は正常な結果として扱う
      minInterval: 1
      maxInterval: 5
    req:
      /api/resource/{{ vars.resource_id }}:
        get:
    test: |
      current.res.status in [200, 404]
    
    dump:
      # エラー情報の詳細記録
      error_info: |
        current.res.status >= 400 ? {
          "status": current.res.status,
          "error": current.res.body.error ?? "Unknown error",
          "timestamp": time.now(),
          "attempt": i + 1
        } : null
```

### デバッグ情報の出力

```yaml
steps:
  debug_step:
    req:
      /api/complex-operation:
        post:
          body:
            application/json: "{{ vars.complex_data }}"
    
    # 詳細なデバッグ情報
    dump:
      request_info:
        url: current.req.url
        method: current.req.method
        headers: current.req.headers
        body_size: len(toJSON(current.req.body))
      
      response_info:
        status: current.res.status
        headers: current.res.headers
        body_size: len(toJSON(current.res.body))
        response_time: current.res.response_time
      
      validation_details:
        expected_fields: ["id", "name", "status"]
        actual_fields: keys(current.res.body)
        missing_fields: |
          filter(["id", "name", "status"], {
            !(. in keys(current.res.body))
          })
        extra_fields: |
          filter(keys(current.res.body), {
            !(. in ["id", "name", "status"])
          })
```

## まとめ

この章では、runnの高度な機能について学びました：

1. **ループ処理**: 基本的なループから条件付きリトライまで
2. **条件付き実行**: 複雑な条件分岐とエラーハンドリング
3. **シナリオのインクルード**: 再利用可能なシナリオの構築
4. **並行実行制御**: パフォーマンステストと共有リソース管理
5. **依存関係の定義**: 複雑なテストシナリオの順序制御
6. **カスタムランナー**: 独自の実行環境の構築
7. **高度なデータ処理**: 動的データ生成と複雑な変換
8. **エラーハンドリング**: 堅牢なテストシナリオの作成

これらの機能を組み合わせることで、実際のプロダクション環境で使用できる高品質なテストスイートを構築できます。次章では、runnの真の力を発揮するGoテストヘルパーとしての活用方法について詳しく解説します。

[第7章：Goテストヘルパー編へ →](chapter07.md)