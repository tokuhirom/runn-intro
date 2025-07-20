# 第3章：Expression文法編 - runnの真の力を解き放て！

## 🚀 expr-lang/expr - 最強の式評価エンジンとの出会い

**ついに来た！** runnの心臓部、[expr-lang/expr](https://expr-lang.org/)の世界へようこそ！これは単なる式評価エンジンじゃない。**テストシナリオに魔法をかける最強の相棒**だ！

### なぜexpr-lang/exprが最高なのか？

- **⚡ Go風の構文**: Goプログラマーなら**5秒で理解できる**直感的な文法！
- **🛡️ 型安全**: 実行時エラー？**そんなものは過去の話**だ！
- **🏃‍♂️ 爆速実行**: コンパイル済み式で**ミリ秒単位の処理**を実現！
- **🔒 完全サンドボックス**: 安全な実行環境で**何も心配いらない**！

## 💪 基本的な式の構文 - これさえ覚えれば無敵！

### 🎯 リテラルと演算子 - あらゆるデータ型を自在に操れ！

<!-- TODO: typeof はないが is はあることを追記 -->

expr-lang/exprでは、さまざまなリテラルと演算子を使って、データを表現することができる。以下の表は、リテラルの種類と例を示している。これらを駆使すれば、**複雑なデータ構造も簡単に扱える**ようになるよ！
詳しくは [expr-lang Literals（リテラル）](https://expr-lang.org/docs/language-definition) を参考にしてね！

| 種類         | 例                                      | 説明              |
|------------|----------------------------------------|-----------------|
| **コメント**   | `// コメント` と `/* コメント */`               | 単一行コメント         |
| **ブール値**   | `true`、`false`                         | 真偽値             |
| **整数**     | `42`、`-37`, `0x2A`, `0o52`, `0b101010` | 整数値             |
| **浮動小数点数** | `0.5`、`.5`                             | 小数値             |
| **文字列**    | `"Hello"`、`'World'`                    | 文字列（単一または二重引用符） |
| **配列**     | `[1, 2, 3]`                            | 配列・リスト          |
| **マップ**    | `{a: 1, b: 2, c: 3}`                   | 連想配列・オブジェクト     |
| **nil**    | `nil`                                  | null値           |

```yaml
{{ includex("examples/chapter03/literals_demo.yml") }}
```

### ⚖️ 比較演算子 - 真偽を見極める審判の目！

```yaml
{{ includex("examples/chapter03/comparison_operators.yml") }}
```

## 🔥 変数参照の詳細 - データの海を自由に泳げ！

### 📊 利用可能な変数一覧 - 7つの強力な武器

| 変数名 | スコープ | 説明 |
|--------|----------|------|
| `vars` | グローバル | Runbookで定義された変数 |
| `env` | グローバル | 環境変数 |
| `steps` | グローバル | すべてのステップの結果 |
| `current` | ステップ内 | 現在のステップの結果 |
| `previous` | ステップ内 | 直前のステップの結果 |
| `i` | ループ内 | ループのインデックス |
| `parent` | Include内 | 親Runbookの変数 |

### 💡 変数アクセスの実践例 - これが本物のパワーだ！

```yaml
{{ includex("examples/chapter03/variable_reference.yml") }}
```

## 🎨 高度な式パターン - プロフェッショナルへの道

### 🔀 条件式（三項演算子） - スマートな分岐処理の極意！

```yaml
{{ includex("examples/chapter03/conditional_expr.yml") }}
```

### 🔍 フィルタリングとマッピング - データ操作の魔術師になれ！

```yaml
{{ includex("examples/chapter03/filter_map_example.yml") }}
```

### 📦 配列・マップ操作 - コレクションを思いのままに！

```yaml
{{ includex("examples/chapter03/array_map_operations.yml") }}
```

## 💼 実践的な式の例 - 現場で使える最強テクニック！

### 🎯 APIレスポンスの検証 - 完璧な検証の極意

```yaml
{{ includex("examples/chapter03/api_response_validation.yml") }}
```

## 🔧 デバッグのテクニック - 問題解決のマスターになる！

### 🔍 dump機能の活用 - すべてを可視化せよ！

TBD

## ⚠️ よくあるパターンと落とし穴 - 達人への必修科目！

### 1. 💀 null/undefinedの扱い - 空の罠を回避せよ！

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

### 2. 🔄 型変換 - データ型の壁を打ち破れ！

```yaml
steps:
  type_conversion:
    test: |
      # 文字列から数値への変換は自動では行われない
      current.res.body.count == "10" &&  # 文字列として比較
      int(current.res.body.count) == 10  # 数値として比較
```

### 3. 🚧 配列の境界チェック - 安全第一の鉄則！

```yaml
steps:
  safe_array_access:
    test: |
      # 配列が空でないことを確認してからアクセス
      len(current.res.body.items) > 0 &&
      current.res.body.items[0].name == "test"
```

## 🎊 まとめ - Expression文法マスターへの道

**おめでとう！** あなたは今、**runnの式評価エンジンの達人**への第一歩を踏み出した！

### 🏆 この章で手に入れた5つの武器：

1. **⚡ 基本的な構文**: リテラル、演算子、比較 - **基礎こそが最強の土台**！
2. **🔑 変数参照**: vars、steps、current、previousなど - **データへの完全アクセス権**！
3. **🎯 高度なパターン**: フィルタリング、マッピング、条件式 - **プロ級のテクニック**！
4. **💪 実践的な使用例**: APIレスポンスの検証、動的リクエスト構築 - **現場で即戦力**！
5. **🔧 デバッグテクニック**: dump機能の活用、段階的な構築 - **問題解決の秘訣**！

**expr-lang/exprの強力な機能**により、どんなに複雑なテストシナリオも**エレガントに記述**できる。でも、これはまだ序章に過ぎない！

**次章では、これらの式で使用できる豊富なビルトイン関数の世界へ飛び込もう！** 準備はいいか？

[第4章：ビルトイン関数編へ →](chapter04.md)
