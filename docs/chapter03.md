# 第3章：Expression文法編

## expr-lang/exprとは

runnは式評価エンジンとして[expr-lang/expr](https://expr-lang.org/)を採用しています。これは、Go言語で実装された高速で安全な式評価エンジンで、以下の特徴があります：

- **Go風の構文**: Goプログラマーには馴染みやすい
- **型安全**: 実行時エラーを最小限に
- **高速**: コンパイル済み式の実行
- **サンドボックス**: 安全な実行環境

## 基本的な式の構文

### リテラルと演算子

```yaml
steps:
  literals_demo:
    test: |
      # 数値
      42 == 42 &&
      3.14 < 4 &&
      
      # 文字列
      "hello" + " world" == "hello world" &&
      "runn" * 2 == "runnrunn" &&
      
      # ブール値
      true && !false &&
      
      # 配列
      [1, 2, 3][0] == 1 &&
      len([1, 2, 3]) == 3 &&
      
      # マップ
      {"name": "alice", "age": 30}.name == "alice"
```

### 比較演算子

```yaml
steps:
  comparison_demo:
    test: |
      # 基本的な比較
      10 > 5 &&
      "apple" < "banana" &&
      100 >= 100 &&
      50 <= 100 &&
      
      # 等価性
      "test" == "test" &&
      100 != 99 &&
      
      # 包含チェック
      "run" in "running" &&
      2 in [1, 2, 3] &&
      "key" in {"key": "value"}
```

## 変数参照の詳細

### 利用可能な変数一覧

| 変数名 | スコープ | 説明 |
|--------|----------|------|
| `vars` | グローバル | Runbookで定義された変数 |
| `env` | グローバル | 環境変数 |
| `steps` | グローバル | すべてのステップの結果 |
| `current` | ステップ内 | 現在のステップの結果 |
| `previous` | ステップ内 | 直前のステップの結果 |
| `i` | ループ内 | ループのインデックス |
| `parent` | Include内 | 親Runbookの変数 |

### 変数アクセスの例

```yaml
desc: 変数参照の包括的な例
vars:
  baseURL: https://api.example.com
  users:
    - id: 1
      name: Alice
    - id: 2
      name: Bob
steps:
  # varsへのアクセス
  access_vars:
    dump: |
      {
        "url": vars.baseURL,
        "firstUser": vars.users[0].name,
        "userCount": len(vars.users)
      }
  
  # 環境変数へのアクセス
  access_env:
    test: |
      env.HOME != "" &&
      env.USER != ""
  
  # ステップ結果へのアクセス（マップ形式）
  make_request:
    req:
      /users:
        get:
    test: current.res.status == 200
  
  # 前のステップの結果を参照
  use_previous:
    test: |
      previous.res.status == 200 &&
      len(steps.make_request.res.body) > 0
```

## 高度な式パターン

### 条件式（三項演算子）

```yaml
steps:
  conditional_expr:
    dump: |
      # 三項演算子
      vars.environment == "prod" ? "https://api.example.com" : "http://localhost:8080"
    
    test: |
      # if式を使った条件分岐
      (current.res.status == 200 ? "success" : "failure") == "success"
```

### フィルタリングとマッピング

```yaml
vars:
  products:
    - name: "iPhone"
      price: 999
      category: "electronics"
    - name: "Book"
      price: 20
      category: "books"
    - name: "MacBook"
      price: 1999
      category: "electronics"
steps:
  filter_example:
    dump: |
      # 価格が100以上の商品をフィルタ
      filter(vars.products, {.price >= 100})
    
    test: |
      # カテゴリが"electronics"の商品数をカウント
      len(filter(vars.products, {.category == "electronics"})) == 2
  
  map_example:
    dump: |
      # 商品名のリストを作成
      map(vars.products, {.name})
    
    test: |
      # すべての商品の価格が0より大きいことを確認
      all(vars.products, {.price > 0})
```

### 配列・マップ操作

```yaml
vars:
  numbers: [1, 2, 3, 4, 5]
  person:
    name: "Alice"
    skills:
      - "Go"
      - "Python"
      - "JavaScript"
steps:
  array_operations:
    test: |
      # スライス操作
      vars.numbers[1:3] == [2, 3] &&
      vars.numbers[:2] == [1, 2] &&
      vars.numbers[3:] == [4, 5] &&
      
      # 要素の存在確認
      3 in vars.numbers &&
      !(10 in vars.numbers) &&
      
      # 配列の結合
      vars.numbers + [6, 7] == [1, 2, 3, 4, 5, 6, 7]
  
  map_operations:
    test: |
      # ネストしたアクセス
      vars.person.skills[0] == "Go" &&
      len(vars.person.skills) == 3 &&
      
      # キーの存在確認
      "name" in vars.person &&
      !("age" in vars.person)
```

## 実践的な式の例

### APIレスポンスの検証

```yaml
desc: 複雑なAPIレスポンスの検証
steps:
  get_users:
    req:
      /users:
        get:
          query:
            page: 1
            limit: 10
    test: |
      # ステータスコードの確認
      current.res.status == 200 &&
      
      # レスポンスボディの構造確認
      "data" in current.res.body &&
      "pagination" in current.res.body &&
      
      # データの検証
      len(current.res.body.data) <= 10 &&
      all(current.res.body.data, {
        "id" in . &&
        "email" in . &&
        .id > 0
      }) &&
      
      # ページネーションの検証
      current.res.body.pagination.page == 1 &&
      current.res.body.pagination.limit == 10
```

### 動的なリクエスト構築

```yaml
vars:
  testUsers:
    - username: "alice"
      role: "admin"
    - username: "bob"
      role: "user"
    - username: "charlie"
      role: "user"
steps:
  # 管理者ユーザーのみを抽出してリクエスト
  create_admin_session:
    req:
      /sessions:
        post:
          body:
            application/json:
              # 管理者ユーザーの最初の1人を取得
              username: filter(vars.testUsers, {.role == "admin"})[0].username
              password: "admin123"
    test: current.res.status == 201
  
  # すべてのユーザーに対してループ処理
  create_all_users:
    loop:
      count: len(vars.testUsers)
    req:
      /users:
        post:
          body:
            application/json:
              username: vars.testUsers[i].username
              role: vars.testUsers[i].role
    test: current.res.status == 201
```

### エラーハンドリング

```yaml
steps:
  api_call_with_retry:
    loop:
      count: 3
      until: current.res.status == 200
    req:
      /unstable-endpoint:
        get:
    test: |
      # 最終的に成功したか、または特定のエラーコード
      current.res.status == 200 ||
      (current.res.status == 503 && i == 2)  # 3回目でも503なら許容
  
  check_error_response:
    req:
      /invalid-endpoint:
        get:
    test: |
      # エラーレスポンスの構造を確認
      current.res.status >= 400 &&
      "error" in current.res.body &&
      current.res.body.error.code != "" &&
      current.res.body.error.message != ""
```

## デバッグのテクニック

### dump機能の活用

```yaml
steps:
  debug_step:
    req:
      /complex-endpoint:
        get:
    dump:
      # 複雑な式の中間結果を出力
      filtered_items: filter(current.res.body.items, {.active == true})
      item_count: len(current.res.body.items)
      first_item_name: current.res.body.items[0].name
      status_check: current.res.status == 200
```

### 式の段階的な構築

```yaml
steps:
  # 複雑な条件を段階的に構築
  complex_validation:
    test: |
      # 基本的な検証
      current.res.status == 200
    
  detailed_validation:
    test: |
      # より詳細な検証を追加
      previous.res.status == 200 &&
      len(previous.res.body.data) > 0 &&
      all(previous.res.body.data, {.id != null})
```

## よくあるパターンと落とし穴

### 1. null/undefinedの扱い

```yaml
steps:
  null_handling:
    test: |
      # nullチェック
      current.res.body.optional_field != null &&
      
      # デフォルト値の設定
      (current.res.body.optional_field ?? "default") != "default" &&
      
      # ネストしたnullチェック
      current.res.body.user?.profile?.bio != null
```

### 2. 型変換

```yaml
steps:
  type_conversion:
    test: |
      # 文字列から数値への変換は自動では行われない
      current.res.body.count == "10" &&  # 文字列として比較
      int(current.res.body.count) == 10  # 数値として比較
```

### 3. 配列の境界チェック

```yaml
steps:
  safe_array_access:
    test: |
      # 配列が空でないことを確認してからアクセス
      len(current.res.body.items) > 0 &&
      current.res.body.items[0].name == "test"
```

## まとめ

この章では、runnの式評価エンジンについて以下を学びました：

1. **基本的な構文**: リテラル、演算子、比較
2. **変数参照**: vars、steps、current、previousなど
3. **高度なパターン**: フィルタリング、マッピング、条件式
4. **実践的な使用例**: APIレスポンスの検証、動的リクエスト構築
5. **デバッグテクニック**: dump機能の活用、段階的な構築

expr-lang/exprの強力な機能により、複雑なテストシナリオも簡潔に記述できます。次章では、これらの式で使用できる豊富なビルトイン関数について詳しく見ていきます。

[第4章：ビルトイン関数編へ →](chapter04.md)