# 第3章：Expression文法編

## expr-lang/expr - runnの式評価エンジン

runnは式評価エンジンとして[expr-lang/expr](https://expr-lang.org/)を採用しています。これにより、柔軟で強力な式評価が可能になります。

### expr-lang/exprの特徴

- Go風の構文
- 型安全性
- 高速な実行速度
- サンドボックス環境での安全な実行

## 基本的な式の構文

### リテラルと演算子

<!-- TODO: type() 関数などを追記 -->

expr-lang/exprでは、さまざまなリテラルと演算子を使ってデータを表現できます。詳細は[expr-lang Literals（リテラル）](https://expr-lang.org/docs/language-definition)を参照してください。

| 種類       | 例                                      | 説明              |
|-----------|----------------------------------------|------------------|
| コメント      | `// コメント` と `/* コメント */`               | 単一行コメント        |
| ブール値      | `true`、`false`                         | 真偽値            |
| 整数        | `42`、`-37`, `0x2A`, `0o52`, `0b101010` | 整数値            |
| 浮動小数点数    | `0.5`、`.5`                             | 小数値            |
| 文字列       | `"Hello"`、`'World'`                    | 文字列            |
| 配列        | `[1, 2, 3]`                            | 配列・リスト         |
| マップ       | `{a: 1, b: 2, c: 3}`                   | 連想配列・オブジェクト    |
| nil       | `nil`                                  | null値          |

<!-- TODO: 演算子表を追加 -->

```yaml
{{ includex("examples/expr-lang/literals_demo.yml") }}
```

### 比較演算子

```yaml
{{ includex("examples/expr-lang/comparison_operators.yml") }}
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

### 変数アクセスの実践例

```yaml
{{ includex("examples/expr-lang/variable_reference.yml") }}
```

## 高度な式パターン

### 条件式（三項演算子）

```yaml
{{ includex("examples/expr-lang/conditional_expr.yml") }}
```

### フィルタリングとマッピング

```yaml
{{ includex("examples/expr-lang/filter_map_example.yml") }}
```

### 配列・マップ操作

```yaml
{{ includex("examples/expr-lang/array_map_operations.yml") }}
```

## 実践的な式の例

### APIレスポンスの検証

```yaml
{{ includex("examples/expr-lang/api_response_validation.yml") }}
```

## デバッグのテクニック

### dump機能の活用

TBD

## よくあるパターンと落とし穴

### 1. null/undefinedの扱い

<!-- TODO: 外部ファイル化して、テストする -->

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

<!-- TODO: 外部ファイル化して、テストする -->

```yaml
steps:
  type_conversion:
    test: |
      # 文字列から数値への変換は自動では行われない
      current.res.body.count == "10" &&  # 文字列として比較
      int(current.res.body.count) == 10  # 数値として比較
```

### 3. 配列の境界チェック

<!-- TODO: 外部ファイル化して、テストする -->

```yaml
steps:
  safe_array_access:
    test: |
      # 配列が空でないことを確認してからアクセス
      len(current.res.body.items) > 0 &&
      current.res.body.items[0].name == "test"
```


