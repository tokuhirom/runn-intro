# 第4章：ビルトイン関数編 - 魔法の関数カタログ！

**ようこそ、runnの関数パラダイスへ！** expr-lang/exprの基本機能だけでも十分強力だが、runnは**さらにその上をいく**！テストやAPIシナリオで**魔法のように使える豊富なビルトイン関数**を用意した。これらの関数を使いこなせば、**どんなに複雑なデータ操作もエレガントに決まる**！

## 🎆 6大関数カテゴリ - あなたの武器庫を充実させろ！

runnのビルトイン関数は**6つの強力なカテゴリ**に分かれている。それぞれが**特定の問題を解決するスペシャリスト**だ！

1. **🔍 比較・差分関数**: データの違いを**瞬時に見抜く**！
2. **🎭 データ操作関数**: オブジェクトを**思いのままに変形**！
3. **✏️ 文字列処理関数**: 文字列を**自在に操る**！
4. **📁 ファイル操作関数**: 外部データと**シームレスに連携**！
5. **⏰ 時間処理関数**: 時間を**完全に支配**！
6. **🎲 テストデータ生成**: **リアルなダミーデータ**を一瞬で！

## 🔍 比較・差分関数 - 差異を見逃さない鷹の目！

### 🎯 compare関数 - 差分を瞬時に検出！

**2つの値の違いを一瞬で見抜く**最強の比較関数！どんな小さな差分も**逃さない**！

```yaml
{{ includex("examples/chapter04/compare_basic.yml") }}
```

### 📋 diff関数 - 人間に優しい差分表示！

文字列やJSONの差分を**誰もが一目で理解できる形式**で出力！**デバッグの強い味方**だ！

```yaml
{{ includex("examples/chapter04/diff_example.yml") }}
```

## 🎭 データ操作関数 - データを自在に変形する魔法！

### ✨ pick関数 - 必要なものだけを優雅に抽出！

オブジェクトから**必要なキーだけをピンポイントで抽出**！センシティブな情報を**スマートに除外**できる！

```yaml
{{ includex("examples/chapter04/pick_example.yml") }}
```

### 🚫 omit関数 - 不要なものをバッサリ捨てろ！

オブジェクトから**不要なキーを一気に削除**！pickの逆バージョンで、**より柔軟なデータ整形**が可能！

```yaml
{{ includex("examples/chapter04/omit_example.yml") }}
```

### 🤝 merge関数 - オブジェクトを融合させろ！

複数のオブジェクトを**スマートに統合**！後のオブジェクトが優先されるから、**設定のオーバーライドに最適**！

```yaml
{{ includex("examples/chapter04/merge_example.yml") }}
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