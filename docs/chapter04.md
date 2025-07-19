# 第4章：ビルトイン関数編

runnは、expr-lang/exprの基本機能に加えて、テストやAPIシナリオで役立つ豊富なビルトイン関数を提供しています。これらの関数により、複雑なデータ操作や検証を簡潔に記述できます。

## 関数のカテゴリ

runnのビルトイン関数は以下のカテゴリに分類されます：

1. **比較・差分関数**: データの比較と差分抽出
2. **データ操作関数**: オブジェクトや配列の加工
3. **文字列処理関数**: 文字列の変換と操作
4. **ファイル操作関数**: ファイルの読み書き
5. **時間処理関数**: 日時の操作とフォーマット
6. **テストデータ生成**: faker関数によるダミーデータ生成

## 比較・差分関数

### compare関数

2つの値を比較し、差分情報を返します。

```yaml
{{ includex("examples/chapter04/compare_basic.yml") }}
```

### diff関数

文字列やJSONの差分を人間が読みやすい形式で出力します。

```yaml
{{ includex("examples/chapter04/diff_example.yml") }}
```

## データ操作関数

### pick関数

オブジェクトから指定したキーのみを抽出します。

```yaml
{{ includex("examples/chapter04/pick_example.yml") }}
```

### omit関数

オブジェクトから指定したキーを除外します。

```yaml
{{ includex("examples/chapter04/omit_example.yml") }}
```

### merge関数

複数のオブジェクトをマージします。後のオブジェクトが優先されます。

```yaml
{{ includex("examples/chapter04/merge_example.yml") }}
```

### intersect関数

複数の配列やオブジェクトの共通部分を返します。

```yaml
{{ includex("examples/chapter04/intersect_example.yml") }}
```

### groupBy関数

配列の要素を指定した条件でグループ化します。

```yaml
{{ includex("examples/chapter04/groupby_example.yml") }}
```

### flatten関数

ネストした配列を平坦化します。

```yaml
{{ includex("examples/chapter04/flatten_example.yml") }}
```

### unique関数

配列から重複を除去します。

```yaml
{{ includex("examples/chapter04/unique_example.yml") }}
```

## 文字列処理関数

### 基本的な文字列関数

```yaml
{{ includex("examples/chapter04/string_basics.yml") }}
```

### 高度な文字列関数

```yaml
{{ includex("examples/chapter04/string_advanced.yml") }}
```

### エンコーディング関数

```yaml
{{ includex("examples/chapter04/encoding_example.yml") }}
```

## ファイル操作関数

### file関数

ファイルの内容を読み込みます。

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

## 時間処理関数

### time関数

現在時刻や時間の操作を行います。

```yaml
{{ includex("examples/chapter04/time_example.yml") }}
```

### 時間の比較と計算

```yaml
{{ includex("examples/chapter04/time_calculation.yml") }}
```

## テストデータ生成（faker関数）

### 基本的なfaker関数

```yaml
{{ includex("examples/chapter04/faker_basic.yml") }}
```

### 数値とランダムデータ

```yaml
{{ includex("examples/chapter04/faker_numbers.yml") }}
```

### 日付関連のfaker関数

```yaml
{{ includex("examples/chapter04/faker_dates.yml") }}
```

### ローカライズされたデータ

```yaml
{{ includex("examples/chapter04/faker_localized.yml") }}
```

## 実用的な使用例

### APIテストでの活用

```yaml
{{ includex("examples/chapter04/api_test_example.yml") }}
```

### データ変換パイプライン

```yaml
{{ includex("examples/chapter04/data_transformation.yml") }}
```

### 複雑なデータ検証

```yaml
{{ includex("examples/chapter04/complex_validation.yml") }}
```

## パフォーマンスとベストプラクティス

### 効率的な関数の使用

```yaml
{{ includex("examples/chapter04/performance_example.yml") }}
```

### エラーハンドリング

```yaml
{{ includex("examples/chapter04/error_handling.yml") }}
```

## まとめ

この章では、runnの豊富なビルトイン関数について学びました：

1. **比較・差分関数**: compare、diffによる詳細な比較
2. **データ操作関数**: pick、omit、merge、groupByなどの強力なデータ加工
3. **文字列処理関数**: 基本的な操作からエンコーディングまで
4. **ファイル・時間関数**: 外部リソースへのアクセスと時間処理
5. **faker関数**: リアルなテストデータの自動生成

これらの関数を組み合わせることで、複雑なテストシナリオも簡潔かつ表現力豊かに記述できます。次章では、これらの関数を活用した各種ランナーの詳細について見ていきます。

[第5章：ランナー詳細編へ →](chapter05.md)