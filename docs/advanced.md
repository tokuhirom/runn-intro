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

`if`フィールドを使用すると、特定の条件が満たされた場合にのみステップを実行できます。

```yaml
{{ includex("examples/advanced/conditional_basic.yml") }}
```

## シナリオのインクルード

他のYAMLファイルをインクルードすることで、共通処理を再利用できます。

```yaml
{{ includex("examples/advanced/include_basic.yml") }}
```

```yaml
{{ includex("examples/advanced/common/auth.include.yml") }}
```

実行結果:
```
{{ includex("examples/advanced/include_basic.stdout") }}
```

## Runbook間の依存関係（needs）

`needs`フィールドを使用すると、他の Runbook を事前実行できます。

```yaml
{{ includex("examples/advanced/needs_basic.yml") }}
```

```yaml
{{ includex("examples/advanced/common/setup.include.yml") }}
```

実行結果:
```
{{ includex("examples/advanced/needs_basic.stdout") }}
```
