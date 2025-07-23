# 第6章：高度な機能編

この章では、runnの高度な機能について説明します。これらの機能を使うことで、より複雑なテストシナリオを作成できます。

## ループ処理

### 基本的なループ

```yaml
{{ includex("examples/advanced/loop_basic.yml") }}
```

### 条件付きループ（リトライ機能）

```yaml
{{ includex("examples/advanced/loop_retry.yml") }}
```

### 動的なループ制御

```yaml
{{ includex("examples/advanced/loop_dynamic.yml") }}
```

## 条件付き実行

### 基本的な条件分岐

```yaml
{{ includex("examples/advanced/conditional_basic.yml") }}
```

### 複雑な条件式

```yaml
{{ includex("examples/advanced/conditional_complex.yml") }}
```

### エラーハンドリングと条件分岐

```yaml
{{ includex("examples/advanced/conditional_error_handling.yml") }}
```

## シナリオのインクルード

### 基本的なインクルード

```yaml
{{ includex("examples/advanced/include_basic.yml") }}
```

```yaml
{{ includex("examples/advanced/common/auth.yml") }}
```

### 動的なインクルード

```yaml
{{ includex("examples/advanced/include_dynamic.yml") }}
```

### ネストしたインクルード

```yaml
{{ includex("examples/advanced/include_nested.yml") }}
```

```yaml
{{ includex("examples/advanced/level2.yml") }}
```

## 並行実行制御

### 基本的な並行実行

```yaml
{{ includex("examples/advanced/concurrency_basic.yml") }}
```

### 共有リソースの制御

```yaml
{{ includex("examples/advanced/concurrency_shared_resource.yml") }}
```

### 複雑な並行制御

```yaml
{{ includex("examples/advanced/concurrency_complex.yml") }}
```

## 依存関係の定義

### 基本的な依存関係

```yaml
{{ includex("examples/advanced/dependency_basic.yml") }}
```

### 複雑な依存関係

```yaml
{{ includex("examples/advanced/dependency_complex.yml") }}
```

## カスタムランナーの作成

### プラグインランナーの定義

```yaml
{{ includex("examples/advanced/custom_runner_plugin.yml") }}
```

### 外部コマンドランナー

```yaml
{{ includex("examples/advanced/custom_runner_external.yml") }}
```

## 高度なデータ処理

### 動的なテストデータ生成

```yaml
{{ includex("examples/advanced/data_generation.yml") }}
```

### 複雑なデータ変換

```yaml
{{ includex("examples/advanced/data_transformation.yml") }}
```

## エラーハンドリングとデバッグ

### 包括的なエラーハンドリング

```yaml
{{ includex("examples/advanced/error_handling.yml") }}
```

### デバッグ情報の出力

```yaml
{{ includex("examples/advanced/debug_output.yml") }}
```

## まとめ

この章では、runnの高度な機能について学びました：

1. ループ処理: 繰り返し処理とリトライ機能
2. 条件付き実行: if文を使った制御
3. シナリオのインクルード: シナリオの再利用
4. 並行実行制御: 複数テストの同時実行
5. 依存関係の定義: テストの実行順序制御
6. カスタムランナー: 独自ランナーの作成
7. 高度なデータ処理: 動的データ生成と変換
8. エラーハンドリング: エラー処理とデバッグ

これらの機能を組み合わせることで、複雑なテストシナリオも効率的に作成できます。

[第7章：Goテストヘルパー編へ →](test-helper.md)