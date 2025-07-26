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

[第7章：Goテストヘルパー編へ →](test-helper.md)