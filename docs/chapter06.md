# 第6章：高度な機能編

この章では、runnの高度な機能について詳しく解説します。これらの機能を使いこなすことで、より複雑で実用的なテストシナリオを構築できます。

## ループ処理

### 基本的なループ

```yaml
{{ includex("examples/chapter06/loop_basic.yml") }}
```

### 条件付きループ（リトライ機能）

```yaml
{{ includex("examples/chapter06/loop_retry.yml") }}
```

### 動的なループ制御

```yaml
{{ includex("examples/chapter06/loop_dynamic.yml") }}
```

## 条件付き実行

### 基本的な条件分岐

```yaml
{{ includex("examples/chapter06/conditional_basic.yml") }}
```

### 複雑な条件式

```yaml
{{ includex("examples/chapter06/conditional_complex.yml") }}
```

### エラーハンドリングと条件分岐

```yaml
{{ includex("examples/chapter06/conditional_error_handling.yml") }}
```

## シナリオのインクルード

### 基本的なインクルード

```yaml
{{ includex("examples/chapter06/include_basic.yml") }}
```

```yaml
{{ includex("examples/chapter06/common/auth.yml") }}
```

### 動的なインクルード

```yaml
{{ includex("examples/chapter06/include_dynamic.yml") }}
```

### ネストしたインクルード

```yaml
{{ includex("examples/chapter06/include_nested.yml") }}
```

```yaml
{{ includex("examples/chapter06/level2.yml") }}
```

## 並行実行制御

### 基本的な並行実行

```yaml
{{ includex("examples/chapter06/concurrency_basic.yml") }}
```

### 共有リソースの制御

```yaml
{{ includex("examples/chapter06/concurrency_shared_resource.yml") }}
```

### 複雑な並行制御

```yaml
{{ includex("examples/chapter06/concurrency_complex.yml") }}
```

## 依存関係の定義

### 基本的な依存関係

```yaml
{{ includex("examples/chapter06/dependency_basic.yml") }}
```

### 複雑な依存関係

```yaml
{{ includex("examples/chapter06/dependency_complex.yml") }}
```

## カスタムランナーの作成

### プラグインランナーの定義

```yaml
{{ includex("examples/chapter06/custom_runner_plugin.yml") }}
```

### 外部コマンドランナー

```yaml
{{ includex("examples/chapter06/custom_runner_external.yml") }}
```

## 高度なデータ処理

### 動的なテストデータ生成

```yaml
{{ includex("examples/chapter06/data_generation.yml") }}
```

### 複雑なデータ変換

```yaml
{{ includex("examples/chapter06/data_transformation.yml") }}
```

## エラーハンドリングとデバッグ

### 包括的なエラーハンドリング

```yaml
{{ includex("examples/chapter06/error_handling.yml") }}
```

### デバッグ情報の出力

```yaml
{{ includex("examples/chapter06/debug_output.yml") }}
```

## まとめ

この章では、runnの高度な機能について学びました：

1. **ループ処理**: 基本的なループから条件付きリトライまで
2. **条件付き実行**: 複雑な条件分岐とエラーハンドリング
3. **シナリオのインクルード**: 再利用可能なシナリオの構築
4. **並行実行制御**: パフォーマンステストと共有リソース管理
5. **依存関係の定義**: 複雑なテストシナリオの順序制御
6. **カスタムランナー**: 独自の実行環境の構築
7. **高度なデータ処理**: 動的データ生成と複雑な変換
8. **エラーハンドリング**: 堅牢なテストシナリオの作成

これらの機能を組み合わせることで、実際のプロダクション環境で使用できる高品質なテストスイートを構築できます。次章では、runnの真の力を発揮するGoテストヘルパーとしての活用方法について詳しく解説します。

[第7章：Goテストヘルパー編へ →](chapter07.md)