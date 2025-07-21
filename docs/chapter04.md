# 第4章：ビルトイン関数編

runnは、テストシナリオで便利に使える独自のビルトイン関数を提供しています。これらの関数により、データの比較、変換、テストデータの生成などが簡単に行えます。

実装の詳細は [book.go](https://github.com/k1LoW/runn/blob/main/book.go) で確認できます。

## runnビルトイン関数一覧

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `urlencode` | 文字列をURLエンコード | `urlencode("hello world")` → `hello%20world` |
| `bool` | 値を真偽値に変換 | `bool("true")` → `true` |
| `compare` | 2つの値を比較（差分があるとエラー） | `compare(actual, expected)` |
| `diff` | 2つの値の差分を表示 | `diff(actual, expected)` |
| `pick` | オブジェクトから指定キーを抽出 | `pick(obj, "key1", "key2")` |
| `omit` | オブジェクトから指定キーを除外 | `omit(obj, "key1", "key2")` |
| `merge` | 複数のオブジェクトをマージ | `merge(obj1, obj2)` |
| `intersect` | 配列の共通要素を取得 | `intersect([1,2,3], [2,3,4])` → `[2,3]` |
| `input` | ユーザー入力を受け付け | `input("名前を入力:")` |
| `secret` | パスワード入力（非表示） | `secret("パスワード:")` |
| `select` | 選択肢から選択 | `select("選択:", ["A","B"], "A")` |
| `basename` | パスからファイル名を抽出 | `basename("/path/to/file.txt")` → `file.txt` |
| `time` | 文字列を時刻に変換 | `time("2024-01-01")` |
| `faker.*` | テストデータ生成 | `faker.Name()`, `faker.Email()` |
| `file` | ファイル内容を読み込み | `file("./data.txt")` |

## urlencode関数

URLエンコードを行う関数です。

```yaml
{{ includex("examples/chapter04/urlencode.yml") }}
```

結果:
```
{{ includex("examples/chapter04/urlencode.stdout") }}
```

## bool関数

文字列や数値を真偽値に変換します。内部的に[cast.ToBool](https://pkg.go.dev/github.com/spf13/cast#ToBool)を使用しています。

```yaml
{{ includex("examples/chapter04/boolean_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/boolean_example.stdout") }}
```

## compare関数

2つの値を比較し、同一かどうかを判定します。オプションで無視するパスを指定できます。

```yaml
{{ includex("examples/chapter04/compare_basic.fail.yml") }}
```

結果:
```
{{ includex("examples/chapter04/compare_basic.fail.out") }}
```

## diff関数

2つの値の差分を人間が読みやすい形式で表示します。

```yaml
{{ includex("examples/chapter04/diff_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/diff_example.stdout") }}
```

## pick関数

オブジェクトから指定したキーのみを抽出します。

```yaml
{{ includex("examples/chapter04/pick_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/pick_example.stdout") }}
```

## omit関数

オブジェクトから指定したキーを除外します。

```yaml
{{ includex("examples/chapter04/omit_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/omit_example.stdout") }}
```

## merge関数

複数のオブジェクトをマージします。後のオブジェクトが優先されます。

```yaml
{{ includex("examples/chapter04/merge_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/merge_example.stdout") }}
```

## intersect関数

2つの配列の共通要素を返します。

```yaml
{{ includex("examples/chapter04/intersect_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/intersect_example.stdout") }}
```

## input関数

実行時にユーザーからの入力を受け付けます。

```yaml
{{ includex("examples/chapter04/input_example.concept.yml") }}
```

## secret関数

パスワードなどの機密情報を入力する際に使用します。入力内容は画面に表示されません。

```yaml
{{ includex("examples/chapter04/secret_example.concept.yml") }}
```

## select関数

複数の選択肢から一つを選択する対話的な入力を提供します。

```yaml
{{ includex("examples/chapter04/select_example.concept.yml") }}
```

## basename関数

ファイルパスからファイル名を抽出します。

```yaml
{{ includex("examples/chapter04/basename_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/basename_example.stdout") }}
```

## time関数

文字列や数値を時刻形式に変換します。

```yaml
{{ includex("examples/chapter04/time_convert_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/time_convert_example.stdout") }}
```

## faker

テストデータを生成するための関数群です。

```yaml
{{ includex("examples/chapter04/faker_builtin_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/faker_builtin_example.stdout") }}
```

## file関数

ファイルの内容を読み込みます。

```yaml
steps:
  file_example:
    desc: ファイルの内容を読み込む
    bind:
      content: file("./data.txt")
    test: |
      // ファイルの内容を検証
      content != null && len(content) > 0
```

**使用例：**
- 設定ファイルの読み込み
- テストデータの読み込み
- テンプレートファイルの読み込み
- 期待値ファイルとの比較

[第5章：ランナー詳細編へ →](chapter05.md)