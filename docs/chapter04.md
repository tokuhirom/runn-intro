# 第4章：runnビルトイン関数編

> 「APIレスポンスのdiffを見やすく表示したい」「テストデータを毎回手動で作るのが面倒」「ファイルの内容を簡単に読み込みたい」
> 
> そんな願いを叶える、**runnの魔法の関数たち**を紹介します。

runnは標準的なexpr-lang関数に加えて、**テストシナリオを劇的に便利にする独自のビルトイン関数**を提供しています。これらの関数を使いこなせば、複雑なテストシナリオもシンプルに記述できるようになります。

### なぜrunnのビルトイン関数が必要なのか？

通常のプログラミング言語でテストを書く場合、こんな課題がありました：

- **差分比較が見にくい**: 大きなJSONの差分を見つけるのが大変
- **テストデータ作成が面倒**: 毎回ランダムなユーザー名やメールアドレスを考える必要がある
- **ファイル操作が煩雑**: ファイルを読み込んでパースして...というコードが冗長
- **対話的なテストが難しい**: パスワード入力などの対話的操作をテストに組み込めない

runnのビルトイン関数は、これらの課題を**シンプルな関数呼び出し一つ**で解決します！

## 🎯 runnビルトイン関数一覧

それぞれの関数が、あなたのテストライフを**劇的に改善**します：

| 関数名 | こんな時に使う！ | 実例 |
|--------|----------------|------|
| `urlencode` | 🔗 URLパラメータを安全にエンコード | `urlencode("検索 キーワード")` → `%E6%A4%9C%E7%B4%A2%20%E3%82%AD%E3%83%BC%E3%83%AF%E3%83%BC%E3%83%89` |
| `bool` | ✅ 文字列や数値を真偽値に変換 | `bool("1")` → `true`、`bool("")` → `false` |
| `compare` | 🔍 レスポンスが期待通りか厳密にチェック | `compare(response, expected)` |
| `diff` | 📊 何が違うのか一目で分かる差分表示 | `diff(actual, expected)` |
| `pick` | 🎯 必要なフィールドだけを抽出 | `pick(user, "id", "name")` |
| `omit` | 🚫 不要なフィールドを除外 | `omit(response, "timestamp", "requestId")` |
| `merge` | 🔄 複数のオブジェクトを合成 | `merge(defaults, overrides)` |
| `intersect` | 🔀 配列の共通要素を発見 | `intersect(tagsA, tagsB)` |
| `input` | 💬 対話的な入力を実現 | `input("APIキーを入力してください:")` |
| `secret` | 🔒 パスワードを安全に入力 | `secret("パスワード:")` |
| `select` | 📋 選択肢から簡単選択 | `select("環境を選択:", ["dev","prod"], "dev")` |
| `basename` | 📁 パスからファイル名をサクッと取得 | `basename("/uploads/image.jpg")` → `image.jpg` |
| `time` | ⏰ 様々な形式の時刻を統一処理 | `time("2024-01-01")` |
| `faker.*` | 🎲 リアルなテストデータを自動生成 | `faker.Name()` → `"田中太郎"` |
| `file` | 📄 ファイル内容を一発読み込み | `file("./testdata.json")` |

## 🔗 urlencode関数

**「日本語パラメータでAPIが動かない！」**という経験はありませんか？

URLエンコードを忘れると、日本語や特殊文字を含むパラメータが正しく送信されません。`urlencode`関数があれば、もう心配無用です：

```yaml
{{ includex("examples/chapter04/urlencode.yml") }}
```

結果:
```
{{ includex("examples/chapter04/urlencode.stdout") }}
```

## ✅ bool関数

**「この値はtrueなの？falseなの？」**と迷ったことはありませんか？

APIレスポンスの文字列"true"や数値の1を真偽値として扱いたい場合、`bool`関数が確実に変換してくれます：

```yaml
{{ includex("examples/chapter04/boolean_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/boolean_example.stdout") }}
```

## 🔍 compare関数

**「レスポンスが期待通りか確認したい、でも差分があったら即座に知りたい！」**

`compare`関数は、2つの値を厳密に比較し、差分があれば**テストを失敗させて詳細を表示**します。タイムスタンプなど、無視したいフィールドも指定可能：

```yaml
{{ includex("examples/chapter04/compare_basic.fail.yml") }}
```

結果:
```
{{ includex("examples/chapter04/compare_basic.fail.out") }}
```

## 📊 diff関数

**「巨大なJSONの中で、どこが違うのか探すのに30分かかった...」**

もうそんな苦労は不要です！`diff`関数は、差分を**色付きで見やすく表示**してくれます：

```yaml
{{ includex("examples/chapter04/diff_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/diff_example.stdout") }}
```

## 🎯 pick関数

**「レスポンスの一部だけをテストしたい」**ときの救世主！

巨大なAPIレスポンスから必要なフィールドだけを抜き出して、スッキリとテストできます：

```yaml
{{ includex("examples/chapter04/pick_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/pick_example.stdout") }}
```

## 🚫 omit関数

**「タイムスタンプやリクエストIDは毎回変わるから、テストから除外したい」**

そんな時は`omit`関数！不要なフィールドを除外して、本質的な部分だけをテストできます：

```yaml
{{ includex("examples/chapter04/omit_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/omit_example.stdout") }}
```

## 🔄 merge関数

**「デフォルト設定に一部だけ上書きしたい」**というケースで大活躍！

`merge`関数を使えば、複数のオブジェクトを賢く合成できます：

```yaml
{{ includex("examples/chapter04/merge_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/merge_example.stdout") }}
```

## 🔀 intersect関数

**「2つのAPIが返すタグの共通部分を知りたい」**

配列の共通要素を見つけるのは意外と面倒。`intersect`関数なら一発です：

```yaml
{{ includex("examples/chapter04/intersect_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/intersect_example.stdout") }}
```

## 💬 input関数

**「テスト実行時にAPIキーを入力したい」「環境によって異なる値を使いたい」**

`input`関数で、対話的なテストシナリオが実現できます：

```yaml
{{ includex("examples/chapter04/input_example.concept.yml") }}
```

## 🔒 secret関数

**「パスワードを入力したいけど、画面に表示されるのは困る！」**

`secret`関数なら、入力内容が***で隠されるので安心です：

```yaml
{{ includex("examples/chapter04/secret_example.concept.yml") }}
```

## 📋 select関数

**「どの環境でテストを実行する？」を毎回選びたい**

`select`関数で、実行時に選択肢から選べる対話的なテストが作れます：

```yaml
{{ includex("examples/chapter04/select_example.concept.yml") }}
```

## 📁 basename関数

**「アップロードされたファイルのパスから、ファイル名だけ取り出したい」**

パス操作は地味に面倒。`basename`関数でサクッと解決：

```yaml
{{ includex("examples/chapter04/basename_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/basename_example.stdout") }}
```

## ⏰ time関数

**「様々な形式の日時文字列を、統一的に扱いたい」**

`time`関数は賢く日時を解析し、Go標準の時刻形式に変換してくれます：

```yaml
{{ includex("examples/chapter04/time_convert_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/time_convert_example.stdout") }}
```

## 🎲 faker関数群

**「テストのたびに『test1@example.com』『田中太郎』って書くの、もう飽きた...」**

`faker`関数群が、**リアルで多様なテストデータを自動生成**してくれます！

```yaml
{{ includex("examples/chapter04/faker_builtin_example.yml") }}
```

結果:
```
{{ includex("examples/chapter04/faker_builtin_example.stdout") }}
```

## 📄 file関数

**「設定ファイルやテストデータをファイルから読み込みたい」**

`file`関数なら、ファイルの内容を**一行で読み込める**シンプルさ：

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

## まとめ：ビルトイン関数でテストが変わる！

runnのビルトイン関数を使いこなせば：

- 🚀 **テストの記述時間が1/3に短縮**
- 🎯 **バグの発見率が向上**（差分が一目瞭然）
- 😊 **テストデータ作成のストレスから解放**
- 🔧 **メンテナンスが圧倒的に楽に**

これらの関数は、あなたのテストライフを**劇的に改善**する強力な武器です。ぜひ活用して、より良いテストを書いていきましょう！

[第5章：ランナー詳細編へ →](chapter05.md)