# 第4章：runnビルトイン関数編

runnは標準的なexpr-lang関数に加えて、テストシナリオ用の独自ビルトイン関数を提供しています。これらの関数により、テストシナリオの記述が容易になります。

### ビルトイン関数の利点

通常のプログラミング言語でテストを書く場合の課題：

- 差分比較が見にくい
- テストデータ作成に手間がかかる
- ファイル操作が複雑
- 対話的なテストが難しい

runnのビルトイン関数は、これらの課題を解決します。

## runnビルトイン関数一覧

以下に主要なビルトイン関数を紹介します。

| 関数名 | 用途 | 例 |
|--------|----------------|------|
| `urlencode` | URLパラメータをエンコード | `urlencode("検索 キーワード")` → `%E6%A4%9C%E7%B4%A2%20%E3%82%AD%E3%83%BC%E3%83%AF%E3%83%BC%E3%83%89` |
| `bool` | 文字列や数値を真偽値に変換 | `bool("1")` → `true`、`bool("")` → `false` |
| `compare` | レスポンスの厳密な比較 | `compare(response, expected)` |
| `diff` | 差分の表示 | `diff(actual, expected)` |
| `pick` | 必要なフィールドの抽出 | `pick(user, "id", "name")` |
| `omit` | 不要なフィールドの除外 | `omit(response, "timestamp", "requestId")` |
| `merge` | 複数のオブジェクトを合成 | `merge(defaults, overrides)` |
| `intersect` | 配列の共通要素を取得 | `intersect(tagsA, tagsB)` |
| `input` | 対話的な入力 | `input("APIキーを入力してください:")` |
| `secret` | パスワードの安全な入力 | `secret("パスワード:")` |
| `select` | 選択肢からの選択 | `select("環境を選択:", ["dev","prod"], "dev")` |
| `basename` | パスからファイル名を取得 | `basename("/uploads/image.jpg")` → `image.jpg` |
| `time` | 時刻の統一処理 | `time("2024-01-01")` |
| `faker.*` | テストデータの自動生成 | `faker.Name()` → `"田中太郎"` |
| `file` | ファイル内容の読み込み | `file("./testdata.json")` |

## urlencode関数

日本語や特殊文字を含むパラメータを正しくエンコードするための関数です。

```yaml
{{ includex("examples/runn-builtins/urlencode.yml") }}
```

結果:
```
{{ includex("examples/runn-builtins/urlencode.stdout") }}
```

## bool関数

APIレスポンスの文字列や数値を真偽値として扱いたい場合に使用します。

```yaml
desc: bool関数の使用例
vars:
  string_true: "true"
  string_false: "false"
  number_one: 1
  number_zero: 0
  empty_string: ""
steps:
  bool_example:
    desc: 様々な値を真偽値に変換
    bind:
      results:
        string_true: bool(vars.string_true)     # true
        string_false: bool(vars.string_false)   # false
        number_one: bool(vars.number_one)       # true
        number_zero: bool(vars.number_zero)     # false
        empty_string: bool(vars.empty_string)   # false
    test: |
      current.results.string_true == true &&
      current.results.string_false == false &&
      current.results.number_one == true &&
      current.results.number_zero == false &&
      current.results.empty_string == false
```

変換ルール：

- 文字列: `"true"` → `true`、`"false"` や空文字 → `false`
- 数値: `1` → `true`、`0` → `false`
- その他: 値が存在すれば `true`、null や空なら `false`

## compare関数

2つの値を厳密に比較し、差分があればテストを失敗させて詳細を表示します。

```yaml
{{ includex("examples/runn-builtins/compare_basic.fail.yml") }}
```

結果:
```
{{ includex("examples/runn-builtins/compare_basic.fail.out") }}
```

## diff関数

差分を色付きで見やすく表示する関数です。

```yaml
{{ includex("examples/runn-builtins/diff_example.yml") }}
```

結果:
```
{{ includex("examples/runn-builtins/diff_example.stdout") }}
```

## pick関数

オブジェクトから必要なフィールドだけを抽出する関数です。

```yaml
{{ includex("examples/runn-builtins/pick_example.yml") }}
```

結果:
```
{{ includex("examples/runn-builtins/pick_example.stdout") }}
```

## omit関数

オブジェクトから不要なフィールドを除外する関数です。タイムスタンプやリクエストIDなど、テストで無視したいフィールドがある場合に便利です。

```yaml
{{ includex("examples/runn-builtins/omit_example.yml") }}
```

結果:
```
{{ includex("examples/runn-builtins/omit_example.stdout") }}
```

## merge関数

複数のオブジェクトを合成する関数です。デフォルト設定に一部だけ上書きしたい場合などに使用します。

```yaml
{{ includex("examples/runn-builtins/merge_example.yml") }}
```

結果:
```
{{ includex("examples/runn-builtins/merge_example.stdout") }}
```

## intersect関数

配列の共通要素を見つける関数です。

```yaml
{{ includex("examples/runn-builtins/intersect_example.yml") }}
```

結果:
```
{{ includex("examples/runn-builtins/intersect_example.stdout") }}
```

## input関数

テスト実行時に対話的に値を入力できる関数です。

```yaml
{{ includex("examples/runn-builtins/input_example.concept.yml") }}
```

## secret関数

パスワードなどの機密情報を安全に入力するための関数です。入力内容は画面に表示されません。

```yaml
{{ includex("examples/runn-builtins/secret_example.concept.yml") }}
```

## select関数

実行時に選択肢から値を選択できる関数です。

```yaml
{{ includex("examples/runn-builtins/select_example.concept.yml") }}
```

## basename関数

ファイルパスからファイル名を取得する関数です。

```yaml
{{ includex("examples/runn-builtins/basename_example.yml") }}
```

結果:
```
{{ includex("examples/runn-builtins/basename_example.stdout") }}
```

## time関数

様々な形式の日時文字列を統一的に扱うための関数です。

```yaml
{{ includex("examples/runn-builtins/time_convert_example.yml") }}
```

結果:
```
{{ includex("examples/runn-builtins/time_convert_example.stdout") }}
```

## faker関数群

リアルで多様なテストデータを自動生成する関数群です。

```yaml
{{ includex("examples/runn-builtins/faker_builtin_example.yml") }}
```

結果:
```
{{ includex("examples/runn-builtins/faker_builtin_example.stdout") }}
```

## file関数

ファイルの内容を読み込む関数です。

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

使用例：

- 設定ファイルの読み込み
- テストデータの読み込み
- テンプレートファイルの読み込み
- 期待値ファイルとの比較

