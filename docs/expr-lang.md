# 第3章：Expression（式）文法編

**Expr**は、式を評価するためのシンプルな式言語です。runnではテストの条件評価や変数の操作にこの式言語を使用します。

## リテラル

<table>
    <tr>
        <td><strong>コメント</strong></td>
        <td>
             <code>/* */</code> または <code>//</code>
        </td>
    </tr>
    <tr>
        <td><strong>ブール値</strong></td>
        <td>
            <code>true</code>, <code>false</code>
        </td>
    </tr>
    <tr>
        <td><strong>整数</strong></td>
        <td>
            <code>42</code>, <code>0x2A</code>, <code>0o52</code>, <code>0b101010</code>
        </td>
    </tr>
    <tr>
        <td><strong>浮動小数点数</strong></td>
        <td>
            <code>0.5</code>, <code>.5</code>
        </td>
    </tr>
    <tr>
        <td><strong>文字列</strong></td>
        <td>
            <code>"foo"</code>, <code>'bar'</code>
        </td>
    </tr>
    <tr>
        <td><strong>配列</strong></td>
        <td>
            <code>[1, 2, 3]</code>
        </td>
    </tr>
    <tr>
        <td><strong>マップ</strong></td>
        <td>
            <code>&#123;a: 1, b: 2, c: 3&#125;</code>
        </td>
    </tr>
    <tr>
        <td><strong>nil</strong></td>
        <td>
            <code>nil</code>
        </td>
    </tr>
</table>

### ブール値リテラル

```yaml
{{ includex("examples/expr-lang/literals-boolean.yml") }}
```

### 数値リテラル

整数は10進数、16進数、8進数、2進数で表現できます：

```yaml
{{ includex("examples/expr-lang/literals-number.yml") }}
```

### 文字列

文字列はシングルクォートまたはダブルクォートで囲むことができます。文字列にはエスケープシーケンス（`\n`で改行、`\t`でタブ、`\uXXXX`でUnicodeコードポイント）を含めることができます。

```yaml
{{ includex("examples/expr-lang/literals-string.yml") }}
```

複数行文字列にはバッククォートを使用します：

```yaml
{{ includex("examples/expr-lang/literals-string-multiline.yml") }}
```

バッククォート文字列は生文字列であり、エスケープシーケンスをサポートしません。

### 配列とマップ

```yaml
{{ includex("examples/expr-lang/literals-array-map.yml") }}
```

## 演算子

<table>
    <tr>
        <td><strong>算術演算子</strong></td>
        <td>
            <code>+</code>, <code>-</code>, <code>*</code>, <code>/</code>, <code>%</code> (剰余), <code>^</code> または <code>**</code> (累乗)
        </td>
    </tr>
    <tr>
        <td><strong>比較演算子</strong></td>
        <td>
            <code>==</code>, <code>!=</code>, <code>&lt;</code>, <code>&gt;</code>, <code>&lt;=</code>, <code>&gt;=</code>
        </td>
    </tr>
    <tr>
        <td><strong>論理演算子</strong></td>
        <td>
            <code>not</code> または <code>!</code>, <code>and</code> または <code>&amp;&amp;</code>, <code>or</code> または <code>||</code>
        </td>
    </tr>
    <tr>
        <td><strong>条件演算子</strong></td>
        <td>
            <code>?:</code> (三項演算子), <code>??</code> (nil合体演算子), <code>if {} else {}</code> (複数行)
        </td>
    </tr>
    <tr>
        <td><strong>メンバーシップ演算子</strong></td>
        <td>
            <code>[]</code>, <code>.</code>, <code>?.</code>, <code>in</code>
        </td>
    </tr>
    <tr>
        <td><strong>文字列演算子</strong></td>
        <td>
            <code>+</code> (連結), <code>contains</code>, <code>startsWith</code>, <code>endsWith</code>
        </td>
    </tr>
    <tr>
        <td><strong>正規表現演算子</strong></td>
        <td>
            <code>matches</code>
        </td>
    </tr>
    <tr>
        <td><strong>範囲演算子</strong></td>
        <td>
            <code>..</code>
        </td>
    </tr>
    <tr>
        <td><strong>スライス演算子</strong></td>
        <td>
            <code>[:]</code>
        </td>
    </tr>
    <tr>
        <td><strong>パイプ演算子</strong></td>
        <td>
            <code>|</code>
        </td>
    </tr>
</table>

### 算術演算子

```yaml
{{ includex("examples/expr-lang/operators-arithmetic.yml") }}
```

### 比較演算子

```yaml
{{ includex("examples/expr-lang/operators-comparison.yml") }}
```

### 論理演算子

```yaml
{{ includex("examples/expr-lang/operators-logical.yml") }}
```

### 条件演算子

```yaml
{{ includex("examples/expr-lang/operators-conditional.yml") }}
```

### メンバーシップ演算子

構造体のフィールドやマップの項目には`.`演算子または`[]`演算子でアクセスできます。次の2つの式は同等です：

```yaml
{{ includex("examples/expr-lang/operators-membership.yml") }}
```

配列やスライスの要素には`[]`演算子でアクセスできます。負のインデックスがサポートされており、`-1`は最後の要素を表します。

```yaml
{{ includex("examples/expr-lang/operators-array-access.yml") }}
```

`in`演算子を使用して、項目が配列やマップに含まれているかを確認できます。

```yaml
{{ includex("examples/expr-lang/operators-in.yml") }}
```

#### オプショナルチェイニング

`?.`演算子を使用すると、構造体やマップが`nil`かどうかをチェックせずに、構造体のフィールドやマップの項目にアクセスできます。構造体やマップが`nil`の場合、式の結果は`nil`になります。

```yaml
{{ includex("examples/expr-lang/operators-optional-chaining.yml") }}
```

#### nil合体演算子

`??`演算子を使用すると、左側が`nil`でない場合は左側を返し、そうでない場合は右側を返します。

```yaml
{{ includex("examples/expr-lang/operators-nil-coalescing.yml") }}
```

### スライス演算子

スライス演算子`[:]`を使用して配列のスライスにアクセスできます。

```yaml
{{ includex("examples/expr-lang/operators-slice.yml") }}
```

### パイプ演算子

パイプ演算子`|`を使用すると、左側の式の結果を右側の式の最初の引数として渡すことができます。

```yaml
{{ includex("examples/expr-lang/operators-pipe.yml") }}
```

### 範囲演算子

範囲演算子`..`を使用して整数の範囲を作成できます。

```yaml
{{ includex("examples/expr-lang/operators-range.yml") }}
```

### 文字列演算子

```yaml
{{ includex("examples/expr-lang/operators-string.yml") }}
```

### 正規表現演算子

```yaml
{{ includex("examples/expr-lang/operators-regex.yml") }}
```

## 変数

変数は`let`キーワードで宣言できます。変数名は文字またはアンダースコアで始まる必要があります。変数名には文字、数字、アンダースコアを含めることができます。変数が宣言された後、式内で使用できます。

```yaml
{{ includex("examples/expr-lang/variables-let.yml") }}
```

複数の変数をセミコロンで区切って複数の`let`文で宣言できます。

```yaml
{{ includex("examples/expr-lang/variables-multiple.yml") }}
```

パイプ演算子を使用した変数の例：

```yaml
{{ includex("examples/expr-lang/variables-pipe.yml") }}
```

### $env

`$env`変数は、式に渡されたすべての変数のマップです。

```yaml
{{ includex("examples/expr-lang/variables-env.yml") }}
```

`$env`はすべての変数を含むグローバル変数と考えることができます。

`$env`を使用して変数が定義されているかを確認できます：

```yaml
{{ includex("examples/expr-lang/variables-env-check.yml") }}
```

## 述語

述語は式です。述語は`filter`、`all`、`any`、`one`、`none`などの関数で使用できます。
例えば、次の式は0から9までの新しい配列を作成し、偶数でフィルタリングします：

```yaml
{{ includex("examples/expr-lang/predicate-basic.yml") }}
```

配列の項目が構造体やマップの場合、`#`シンボルを省略してフィールドにアクセスできます（`#.Value`は`.Value`になります）。

```yaml
{{ includex("examples/expr-lang/predicate-omit.yml") }}
```

波括弧`{` `}`は省略できます：

```yaml
{{ includex("examples/expr-lang/predicate-no-braces.yml") }}
```

ネストした述語では、外側の変数にアクセスするために[変数](#変数)を使用します。

```yaml
{{ includex("examples/expr-lang/predicate-nested.yml") }}
```

## 文字列関数

### trim(str[, chars])

文字列`str`の両端から空白文字を削除します。
オプションの`chars`引数が指定された場合、削除する文字のセットを指定する文字列です。

```yaml
{{ includex("examples/expr-lang/functions-string-trim.yml") }}
```

### trimPrefix(str, prefix)

文字列`str`が指定されたプレフィックスで始まる場合、そのプレフィックスを削除します。

```yaml
{{ includex("examples/expr-lang/functions-string-trimprefix.yml") }}
```

### trimSuffix(str, suffix)

文字列`str`が指定されたサフィックスで終わる場合、そのサフィックスを削除します。

```yaml
{{ includex("examples/expr-lang/functions-string-trimsuffix.yml") }}
```

### upper(str)

文字列`str`のすべての文字を大文字に変換します。

```yaml
{{ includex("examples/expr-lang/functions-string-upper.yml") }}
```

### lower(str)

文字列`str`のすべての文字を小文字に変換します。

```yaml
{{ includex("examples/expr-lang/functions-string-lower.yml") }}
```

### split(str, delimiter[, n])

文字列`str`を区切り文字の各インスタンスで分割し、部分文字列の配列を返します。

```yaml
{{ includex("examples/expr-lang/functions-string-split.yml") }}
```

### splitAfter(str, delimiter[, n])

文字列`str`を区切り文字の各インスタンスの後で分割します。

```yaml
{{ includex("examples/expr-lang/functions-string-splitafter.yml") }}
```

### replace(str, old, new)

文字列`str`内のすべての`old`の出現を`new`で置換します。

```yaml
{{ includex("examples/expr-lang/functions-string-replace.yml") }}
```

### repeat(str, n)

文字列`str`を`n`回繰り返します。

```yaml
{{ includex("examples/expr-lang/functions-string-repeat.yml") }}
```

### indexOf(str, substring)

文字列`str`内の部分文字列の最初の出現のインデックスを返します。見つからない場合は-1を返します。

```yaml
{{ includex("examples/expr-lang/functions-string-indexof.yml") }}
```

### lastIndexOf(str, substring)

文字列`str`内の部分文字列の最後の出現のインデックスを返します。見つからない場合は-1を返します。

```yaml
{{ includex("examples/expr-lang/functions-string-lastindexof.yml") }}
```

### hasPrefix(str, prefix)

文字列`str`が指定されたプレフィックスで始まる場合、`true`を返します。

```yaml
{{ includex("examples/expr-lang/functions-string-hasprefix.yml") }}
```

### hasSuffix(str, suffix)

文字列`str`が指定されたサフィックスで終わる場合、`true`を返します。

```yaml
{{ includex("examples/expr-lang/functions-string-hassuffix.yml") }}
```

## 日時関数

ExprはGoの[timeパッケージ](https://pkg.go.dev/time)の組み込みサポートを持っています。
2つの日付を減算して、その間の期間を取得できます：

```yaml
{{ includex("examples/expr-lang/functions-date-subtract.yml") }}
```

日付に期間を追加できます：

```yaml
{{ includex("examples/expr-lang/functions-date-add.yml") }}
```

日付を比較できます：

```yaml
{{ includex("examples/expr-lang/functions-date-compare.yml") }}
```

### now()

現在の日付を[time.Time](https://pkg.go.dev/time#Time)値として返します。

```yaml
{{ includex("examples/expr-lang/functions-date-now.yml") }}
```

### duration(str)

指定された文字列`str`の[time.Duration](https://pkg.go.dev/time#Duration)値を返します。

有効な時間単位は"ns"、"us"（または"µs"）、"ms"、"s"、"m"、"h"です。

```yaml
{{ includex("examples/expr-lang/functions-date-duration.yml") }}
```

### date(str[, format[, timezone]])

指定された文字列`str`を日付表現に変換します。

オプションの`format`引数が指定された場合、日付のフォーマットを指定する文字列です。
フォーマット文字列は、標準のGo [timeパッケージ](https://pkg.go.dev/time#pkg-constants)と同じフォーマットルールを使用します。

オプションの`timezone`引数が指定された場合、日付のタイムゾーンを指定する文字列です。

`format`引数が指定されない場合、`v`引数は次のいずれかのフォーマットである必要があります：

- 2006-01-02
- 15:04:05
- 2006-01-02 15:04:05
- RFC3339
- RFC822
- RFC850
- RFC1123

```yaml
{{ includex("examples/expr-lang/functions-date-date.yml") }}
```

日付で使用可能なメソッド：

- `Year()` - 年を返す
- `Month()` - 月を返す（1から開始）
- `Day()` - 月の日を返す
- `Hour()` - 時間を返す
- `Minute()` - 分を返す
- `Second()` - 秒を返す
- `Weekday()` - 曜日を返す
- `YearDay()` - 年の日を返す
- その他（[詳細](https://pkg.go.dev/time#Time)）

```yaml
{{ includex("examples/expr-lang/functions-date-methods.yml") }}
```

### timezone(str)

指定された文字列`str`のタイムゾーンを返します。利用可能なタイムゾーンのリストは[こちら](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)で確認できます。

```yaml
{{ includex("examples/expr-lang/functions-date-timezone.yml") }}
```

日付を別のタイムゾーンに変換するには、[`In()`](https://pkg.go.dev/time#Time.In)メソッドを使用します：

```yaml
{{ includex("examples/expr-lang/functions-date-timezone-convert.yml") }}
```

## 数値関数

### max(n1, n2)

2つの数値`n1`と`n2`の最大値を返します。

```yaml
{{ includex("examples/expr-lang/functions-number-max.yml") }}
```

### min(n1, n2)

2つの数値`n1`と`n2`の最小値を返します。

```yaml
{{ includex("examples/expr-lang/functions-number-min.yml") }}
```

### abs(n)

数値の絶対値を返します。

```yaml
{{ includex("examples/expr-lang/functions-number-abs.yml") }}
```

### ceil(n)

x以上の最小の整数値を返します。

```yaml
{{ includex("examples/expr-lang/functions-number-ceil.yml") }}
```

### floor(n)

x以下の最大の整数値を返します。

```yaml
{{ includex("examples/expr-lang/functions-number-floor.yml") }}
```

### round(n)

最も近い整数を返し、0から離れる方向に丸めます。

```yaml
{{ includex("examples/expr-lang/functions-number-round.yml") }}
```

## 配列関数

### all(array, predicate)

すべての要素が[述語](#述語)を満たす場合、**true**を返します。
配列が空の場合、**true**を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-all.yml") }}
```

### any(array, predicate)

いずれかの要素が[述語](#述語)を満たす場合、**true**を返します。
配列が空の場合、**false**を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-any.yml") }}
```

### one(array, predicate)

正確に1つの要素が[述語](#述語)を満たす場合、**true**を返します。
配列が空の場合、**false**を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-one.yml") }}
```

### none(array, predicate)

すべての要素が[述語](#述語)を満たさない場合、**true**を返します。
配列が空の場合、**true**を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-none.yml") }}
```

### map(array, predicate)

配列の各要素に[述語](#述語)を適用して新しい配列を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-map.yml") }}
```

### filter(array, predicate)

配列の要素を[述語](#述語)でフィルタリングして新しい配列を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-filter.yml") }}
```

### find(array, predicate)

配列内で[述語](#述語)を満たす最初の要素を見つけます。

```yaml
{{ includex("examples/expr-lang/functions-array-find.yml") }}
```

### findIndex(array, predicate)

配列内で[述語](#述語)を満たす最初の要素のインデックスを見つけます。

```yaml
{{ includex("examples/expr-lang/functions-array-findindex.yml") }}
```

### findLast(array, predicate)

配列内で[述語](#述語)を満たす最後の要素を見つけます。

```yaml
{{ includex("examples/expr-lang/functions-array-findlast.yml") }}
```

### findLastIndex(array, predicate)

配列内で[述語](#述語)を満たす最後の要素のインデックスを見つけます。

```yaml
{{ includex("examples/expr-lang/functions-array-findlastindex.yml") }}
```

### groupBy(array, predicate)

[述語](#述語)の結果によって配列の要素をグループ化します。

```yaml
{{ includex("examples/expr-lang/functions-array-groupby.yml") }}
```

### count(array[, predicate])

[述語](#述語)を満たす要素の数を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-count.yml") }}
```

述語が指定されない場合、配列内の`true`要素の数を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-count-boolean.yml") }}
```

### concat(array1, array2[, ...])

2つ以上の配列を連結します。

```yaml
{{ includex("examples/expr-lang/functions-array-concat.yml") }}
```

### flatten(array)

指定された配列を1次元配列に平坦化します。

```yaml
{{ includex("examples/expr-lang/functions-array-flatten.yml") }}
```

### uniq(array)

配列から重複を削除します。

```yaml
{{ includex("examples/expr-lang/functions-array-uniq.yml") }}
```

### join(array[, delimiter])

文字列の配列を指定された区切り文字で単一の文字列に結合します。
区切り文字が指定されない場合、空文字列が使用されます。

```yaml
{{ includex("examples/expr-lang/functions-array-join.yml") }}
```

### reduce(array, predicate[, initialValue])

配列の各要素に述語を適用し、配列を単一の値に縮約します。
オプションの`initialValue`引数を使用してアキュムレータの初期値を指定できます。
`initialValue`が指定されない場合、配列の最初の要素が初期値として使用されます。

述語内で使用できる変数：

- `#` - 現在の要素
- `#acc` - アキュムレータ
- `#index` - 現在の要素のインデックス

```yaml
{{ includex("examples/expr-lang/functions-array-reduce.yml") }}
```

### sum(array[, predicate])

配列内のすべての数値の合計を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-sum.yml") }}
```

オプションの`predicate`引数が指定された場合、合計する前に配列の各要素に適用される述語です。

```yaml
{{ includex("examples/expr-lang/functions-array-sum-predicate.yml") }}
```

### mean(array)

配列内のすべての数値の平均を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-mean.yml") }}
```

### median(array)

配列内のすべての数値の中央値を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-median.yml") }}
```

### first(array)

配列から最初の要素を返します。配列が空の場合、`nil`を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-first.yml") }}
```

### last(array)

配列から最後の要素を返します。配列が空の場合、`nil`を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-last.yml") }}
```

### take(array, n)

配列から最初の`n`個の要素を返します。配列の要素が`n`個未満の場合、配列全体を返します。

```yaml
{{ includex("examples/expr-lang/functions-array-take.yml") }}
```

### reverse(array)

配列の新しい逆順コピーを返します。

```yaml
{{ includex("examples/expr-lang/functions-array-reverse.yml") }}
```

### sort(array[, order])

配列を昇順でソートします。オプションの`order`引数を使用してソート順を指定できます：`asc`または`desc`。

```yaml
{{ includex("examples/expr-lang/functions-array-sort.yml") }}
```

### sortBy(array[, predicate, order])

[述語](#述語)の結果で配列をソートします。オプションの`order`引数を使用してソート順を指定できます：`asc`または`desc`。

```yaml
{{ includex("examples/expr-lang/functions-array-sortby.yml") }}
```

## マップ関数

### keys(map)

マップのキーを含む配列を返します。

```yaml
{{ includex("examples/expr-lang/functions-map-keys.yml") }}
```

### values(map)

マップの値を含む配列を返します。

```yaml
{{ includex("examples/expr-lang/functions-map-values.yml") }}
```

## 型変換関数

### type(v)

指定された値`v`の型を返します。

次のいずれかの型を返します：

- `nil`
- `bool`
- `int`
- `uint`
- `float`
- `string`
- `array`
- `map`

名前付き型と構造体の場合、型名が返されます。

```yaml
{{ includex("examples/expr-lang/functions-type-type.yml") }}
```

### int(v)

数値または文字列の整数値を返します。

```yaml
{{ includex("examples/expr-lang/functions-type-int.yml") }}
```

### float(v)

数値または文字列の浮動小数点値を返します。

```yaml
{{ includex("examples/expr-lang/functions-type-float.yml") }}
```

### string(v)

指定された値`v`を文字列表現に変換します。

```yaml
{{ includex("examples/expr-lang/functions-type-string.yml") }}
```

### toJSON(v)

指定された値`v`をJSON文字列表現に変換します。

```yaml
{{ includex("examples/expr-lang/functions-type-tojson.yml") }}
```

### fromJSON(v)

指定されたJSON文字列`v`を解析し、対応する値を返します。

```yaml
{{ includex("examples/expr-lang/functions-type-fromjson.yml") }}
```

### toBase64(v)

文字列`v`をBase64形式にエンコードします。

```yaml
{{ includex("examples/expr-lang/functions-type-tobase64.yml") }}
```

### fromBase64(v)

Base64エンコードされた文字列`v`を元の形式にデコードします。

```yaml
{{ includex("examples/expr-lang/functions-type-frombase64.yml") }}
```

### toPairs(map)

マップをキーと値のペアの配列に変換します。

```yaml
{{ includex("examples/expr-lang/functions-type-topairs.yml") }}
```

### fromPairs(array)

キーと値のペアの配列をマップに変換します。

```yaml
{{ includex("examples/expr-lang/functions-type-frompairs.yml") }}
```

## その他の関数

### len(v)

配列、マップ、または文字列の長さを返します。

```yaml
{{ includex("examples/expr-lang/functions-misc-len.yml") }}
```

### get(v, index)

配列またはマップ`v`から指定されたインデックスの要素を取得します。インデックスが範囲外の場合、`nil`を返します。
またはキーが存在しない場合、`nil`を返します。

```yaml
{{ includex("examples/expr-lang/functions-misc-get.yml") }}
```

## ビット演算関数

### bitand(int, int)

ビットAND演算の結果を返します。

```yaml
{{ includex("examples/expr-lang/functions-bitwise-bitand.yml") }}
```

### bitor(int, int)

ビットOR演算の結果を返します。

```yaml
{{ includex("examples/expr-lang/functions-bitwise-bitor.yml") }}
```

### bitxor(int, int)

ビットXOR演算の結果を返します。

```yaml
{{ includex("examples/expr-lang/functions-bitwise-bitxor.yml") }}
```

### bitnand(int, int)

ビットAND NOT演算の結果を返します。

```yaml
{{ includex("examples/expr-lang/functions-bitwise-bitnand.yml") }}
```

### bitnot(int)

ビットNOT演算の結果を返します。

```yaml
{{ includex("examples/expr-lang/functions-bitwise-bitnot.yml") }}
```

### bitshl(int, int)

左シフト演算の結果を返します。

```yaml
{{ includex("examples/expr-lang/functions-bitwise-bitshl.yml") }}
```

### bitshr(int, int)

右シフト演算の結果を返します。

```yaml
{{ includex("examples/expr-lang/functions-bitwise-bitshr.yml") }}
```

### bitushr(int, int)

符号なし右シフト演算の結果を返します。

```yaml
{{ includex("examples/expr-lang/functions-bitwise-bitushr.yml") }}
```