# 第4章：ビルトイン関数編 - 魔法の関数カタログ！

**ようこそ、runnの関数パラダイスへ！** expr-lang/exprの基本機能だけでも十分強力だが、runnは**さらにその上をいく**！テストやAPIシナリオで**魔法のように使える豊富なビルトイン関数**を用意している。
これらの関数を使いこなせば、**どんなに複雑なデータ操作もエレガントに決まる**！

実装が気になるスーパーハカーな君は [book.go](https://github.com/k1LoW/runn/blob/main/book.go) を見れば辿れるぞ！

## urlencode関数 - URLエンコードの達人！

urlencode関数は、**URLエンコードを一瞬で行う**魔法の関数だ。

```yaml
{{ includex("examples/chapter04/urlencode.yml") }}
```

結果は以下のように出力される。

```
{{ includex("examples/chapter04/urlencode.stdout") }}
```

## bool 関数 - 真偽値のマスター！

bool　関数は、**真偽値を簡単に取得**できる便利な関数だ。文字列や数値を**真偽値に変換**する。
[cast.ToBool](https://pkg.go.dev/github.com/spf13/cast#ToBool) を内部的に利用しているぞ。

これをどんな時に使ったらいいかはよくわからない！

```yaml
{{ includex("examples/chapter04/boolean_example.yml") }}
```

出力結果は以下のようになる。

```
{{ includex("examples/chapter04/boolean_example.stdout") }}
```

## 🎯 compare関数 - 差分を瞬時に検出！

**2つの値の違いを一瞬で見抜く**最強の比較関数！どんな小さな差分も**逃さない**！

```yaml
{{ includex("examples/chapter04/compare_basic.fail.yml") }}
```

出力結果はこんな感じだ。

```
{{ includex("examples/chapter04/compare_basic.fail.out") }}
```

<!-- TODO: compare の第３引数以後には jq の ignore list がかけるっぽいので
     それを使った例も追加する -->

## 📋 diff関数 - 人間に優しい差分表示！

データ構造の差分を**誰もが一目で理解できる形式**で出力！**デバッグの強い味方**だ！

```yaml
{{ includex("examples/chapter04/diff_example.yml") }}
```

出力は以下のようになる。 

```
{{ includex("examples/chapter04/diff_example.stdout") }}
```

文字列が Unified diff で出るのかと思いきやそうではないので注意が必要だ！ 

## ✨ pick関数 - 必要なものだけを優雅に抽出！

[lo.PickByKeys](https://github.com/samber/lo?tab=readme-ov-file#pickbykeys) 関数を使って、**オブジェクトから必要なキーだけをピックアップ**！

golang でいうと以下のような感じだ。

```go
m := lo.PickByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})
// map[string]int{"foo": 1, "baz": 3}
```

runn のシナリオでは以下のように書けばいいね！

```yaml
{{ includex("examples/chapter04/pick_example.yml") }}
```

出力は以下のようになるね！

```
{{ includex("examples/chapter04/pick_example.stdout") }}
```

## 🚫 omit関数 - 不要なものをバッサリ捨てろ！

[lo.OmitByKeys](https://github.com/samber/lo?tab=readme-ov-file#omitbykeys) を使って、**オブジェクトから不要なキーを一気に削除**！pickの逆バージョンだ！

golang でいうと以下のような感じだ。

```go
m := lo.OmitByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})
// map[string]int{"bar": 2}
```

```yaml
{{ includex("examples/chapter04/omit_example.yml") }}
```

結果はこーなるね！

```
{{ includex("examples/chapter04/omit_example.stdout") }}
```

<!-- TODO: 以後は見直し必要 -->

### 🤝 merge関数 - オブジェクトを融合させろ！

複数のオブジェクトを**スマートに統合**！後のオブジェクトが優先されるから、**設定のオーバーライドに最適**！

[lo.Assign](https://github.com/samber/lo?tab=readme-ov-file#assign) 関数を使って実装されているぞ！

```yaml
{{ includex("examples/chapter04/merge_example.yml") }}
```

結果は以下のようになる：

```
{{ includex("examples/chapter04/merge_example.stdout") }}
```

### ∩ intersect関数 - 共通部分を見つけ出せ！

複数の配列やオブジェクトの**共通部分だけを抽出**！集合演算の**積集合を一瞬で計算**！

```yaml
{{ includex("examples/chapter04/intersect_example.yml") }}
```

### 📦 groupBy関数 - データを賢く分類！

配列の要素を**指定した条件で自動分類**！データ分析や**統計処理に必須**の機能！

```yaml
{{ includex("examples/chapter04/groupby_example.yml") }}
```

### 🎯 flatten関数 - ネストを解きほぐせ！

ネストした配列を**フラットな1次元配列に変換**！階層構造を**シンプルに整理**！

```yaml
{{ includex("examples/chapter04/flatten_example.yml") }}
```

### 💎 unique関数 - 唯一無二の値だけを残せ！

配列から**重複を完全排除**！ユニークな値だけを**スマートに抽出**！

```yaml
{{ includex("examples/chapter04/unique_example.yml") }}
```

## ✏️ 文字列処理関数 - テキストを思いのままに！

### 📝 基本的な文字列関数 - 日常使いの必須ツール！

```yaml
{{ includex("examples/chapter04/string_basics.yml") }}
```

### 🎯 高度な文字列関数 - プロフェッショナルの技！

```yaml
{{ includex("examples/chapter04/string_advanced.yml") }}
```

### 🔐 エンコーディング関数 - データ形式の変換術師！

```yaml
{{ includex("examples/chapter04/encoding_example.yml") }}
```

## 📁 ファイル操作関数 - 外部データとの架け橋！

### 📄 file関数 - ファイルを一瞬で読み込め！

**あらゆるファイルを瞬時に読み込む**魔法の関数！JSON、YAML、テキスト、画像まで**何でも来い**！

```yaml
steps:
  file_example:
    dump: |
      {
        # テキストファイルの読み込み
        "config": file("./config.json"),
        # YAMLファイルの読み込み（自動的にパース）
        "settings": file("./settings.yml"),
        # バイナリファイルはbase64エンコード
        "image": file("./logo.png")
      }
```

## ⏰ 時間処理関数 - 時間を支配するタイムマスター！

### ⌚ time関数 - 時間を自在に操れ！

**現在、過去、未来**、どんな時間も思いのまま！時間の**加算、減算、フォーマット**まで完璧にサポート！

```yaml
{{ includex("examples/chapter04/time_example.yml") }}
```

### 📏 時間の比較と計算 - 時間計算の達人に！

```yaml
{{ includex("examples/chapter04/time_calculation.yml") }}
```

## 🎲 テストデータ生成（faker関数） - リアルなダミーデータを一瞬で！

### 🎆 基本的なfaker関数 - テストデータの魔術師！

```yaml
{{ includex("examples/chapter04/faker_basic.yml") }}
```

### 🎰 数値とランダムデータ - 予測不可能なデータを生成！

```yaml
{{ includex("examples/chapter04/faker_numbers.yml") }}
```

### 📅 日付関連のfaker関数 - 時間を超えたデータ生成！

```yaml
{{ includex("examples/chapter04/faker_dates.yml") }}
```

### 🌍 ローカライズされたデータ - 世界各国のデータを再現！

```yaml
{{ includex("examples/chapter04/faker_localized.yml") }}
```

## 💼 実用的な使用例 - 現場で即戦力のテクニック！

### 🎯 APIテストでの活用 - プロのテストシナリオ！

```yaml
{{ includex("examples/chapter04/api_test_example.yml") }}
```

### 🏭 データ変換パイプライン - データを華麗に変身！

```yaml
{{ includex("examples/chapter04/data_transformation.yml") }}
```

### 🔍 複雑なデータ検証 - どんなデータも完璧にチェック！

```yaml
{{ includex("examples/chapter04/complex_validation.yml") }}
```

## ⚡ パフォーマンスとベストプラクティス - プロの流儀！

### 🏃‍♂️ 効率的な関数の使用 - パフォーマンスを最大化！

```yaml
{{ includex("examples/chapter04/performance_example.yml") }}
```

### 🛡️ エラーハンドリング - 安全第一の鉄壁ガード！

```yaml
{{ includex("examples/chapter04/error_handling.yml") }}
```

## 🎉 まとめ - ビルトイン関数マスター誕生！

**素晴らしい！** あなたは今、**runnのビルトイン関数の全てをマスター**した！

### 🏆 この章で習得した5つの必殺技：

1. **🔍 比較・差分関数**: compare、diffで**差異を瞬時に発見**！
2. **🎭 データ操作関数**: pick、omit、merge、groupByで**データを自在に変形**！
3. **✏️ 文字列処理関数**: 基本からエンコーディングまで**完全網羅**！
4. **📁⏰ ファイル・時間関数**: 外部リソースも時間も**思いのまま**！
5. **🎲 faker関数**: **リアルなテストデータ**を無限に生成！

これらの関数を**自由自在に組み合わせれば**、どんなに複雑なテストシナリオも**エレガントに表現**できる。あなたはもう、**テストデータの魔術師**だ！

**次章では、これらの関数をフル活用した各種ランナーの世界へ突入しよう！** 準備はいいか？

[第5章：ランナー詳細編へ →](chapter05.md)