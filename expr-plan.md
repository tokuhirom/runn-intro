# expr-lang.md 実装計画

## 概要

`docs/expr-lang.md`を、expr-langの公式言語定義（https://expr-lang.org/docs/language-definition）をベースに作成する。

### 基本方針

1. **構成**: expr-langの公式ドキュメントと同じ順番・構成を維持
2. **言語**: すべて日本語に翻訳
3. **サンプルコード**: すべてrunnのシナリオ（YAMLファイル）として作成
4. **配置**: サンプルは`examples/expr-lang/`ディレクトリに配置
5. **テスト**: `chapters_test.go`から実行できるようにする
6. **既存コンテンツ**: 現在の`expr-lang.md`の内容は破棄可能

## ドキュメント構成

以下の順番で章立てを行う（expr-lang公式ドキュメントと同一）：

### 1. リテラル (Literals)
- コメント
- ブール値
- 整数（10進数、16進数、8進数、2進数）
- 浮動小数点数
- 文字列
- 配列
- マップ
- nil

#### 1.1 文字列 (Strings)
- エスケープシーケンス
- 複数行文字列（バッククォート）

### 2. 演算子 (Operators)
- 算術演算子
- 比較演算子
- 論理演算子
- 条件演算子
- メンバーシップ演算子
- 文字列演算子
- 正規表現演算子
- 範囲演算子
- スライス演算子
- パイプ演算子

#### 2.1 メンバーシップ演算子 (Membership Operator)
- フィールドアクセス（`.`と`[]`）
- 配列・スライスアクセス
- `in`演算子
- オプショナルチェイニング（`?.`）
- nil合体演算子（`??`）

#### 2.2 スライス演算子 (Slice Operator)
#### 2.3 パイプ演算子 (Pipe Operator)
#### 2.4 範囲演算子 (Range Operator)

### 3. 変数 (Variables)
- `let`キーワード
- 複数変数の宣言
- `$env`変数

### 4. 述語 (Predicate)
- フィルタ関数での使用
- `#`シンボルの省略
- ネストした述語

### 5. 文字列関数 (String Functions)
- trim, trimPrefix, trimSuffix
- upper, lower
- split, splitAfter
- replace, repeat
- indexOf, lastIndexOf
- hasPrefix, hasSuffix

### 6. 日時関数 (Date Functions)
- now, duration, date
- timezone
- 日時の演算

### 7. 数値関数 (Number Functions)
- max, min, abs
- ceil, floor, round

### 8. 配列関数 (Array Functions)
- all, any, one, none
- map, filter
- find, findIndex, findLast, findLastIndex
- groupBy, count
- concat, flatten, uniq
- join
- reduce, sum, mean, median
- first, last, take
- reverse, sort, sortBy

### 9. マップ関数 (Map Functions)
- keys, values

### 10. 型変換関数 (Type Conversion Functions)
- type
- int, float, string
- toJSON, fromJSON
- toBase64, fromBase64
- toPairs, fromPairs

### 11. その他の関数 (Miscellaneous Functions)
- len, get

### 12. ビット演算関数 (Bitwise Functions)
- bitand, bitor, bitxor, bitnand
- bitnot
- bitshl, bitshr, bitushr

## サンプルコード作成方針

各セクションに対して、以下の形式でrunnシナリオを作成：

1. **ファイル名規則**:
   - `examples/expr-lang/literals-boolean.yml`
   - `examples/expr-lang/operators-arithmetic.yml`
   - `examples/expr-lang/functions-string-trim.yml`
   - など、セクション名と機能を組み合わせた名前

2. **YAMLフォーマット**:
   ```yaml
   desc: "expr-lang: ブール値リテラルの使用例"
   runners:
     httpbin: https://httpbin.org
   steps:
     boolean_example:
       desc: "ブール値の評価"
       bind:
         is_active: true
         is_disabled: false
       test: |
         // ブール値の評価
         current.is_active == true &&
         current.is_disabled == false
   ```

3. **テストケース**:
   - 各関数・演算子の基本的な使い方
   - エッジケース
   - 実践的な使用例

## テスト実装

`chapters_test.go`に以下の関数を追加：

```go
func TestExprLang(t *testing.T) {
    testFiles(t, "examples/expr-lang")
}
```

## 作業手順

1. 現在の`docs/expr-lang.md`をバックアップ（必要な場合）
2. 新しい`docs/expr-lang.md`を作成開始
3. `examples/expr-lang/`ディレクトリを作成
4. 各セクションごとに：
   - 日本語翻訳
   - runnシナリオ作成
   - テスト実行確認
5. `chapters_test.go`にテスト関数追加
6. 全体のテスト実行確認

## 注意事項

- expr-lang式内のコメントは`//`スタイルを使用（`#`は非推奨）
- YAMLのtestセクションでは複数行の式を使用可能
- 各サンプルは実際に動作することを確認
- runnの変数（vars, steps, current等）と組み合わせた実践的な例を含める