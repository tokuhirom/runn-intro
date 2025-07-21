# 第6章：高度な機能編 - runnの真の力を解放せよ！

**ここからが本番だ！** 今まで学んできた機能はほんの**序の口**に過ぎない。この章では、runnの**本当にすごい機能**を伝授しよう！これらの機能をマスターすれば、**どんなに複雑なテストシナリオも余裕**だ！

## 🔁 ループ処理 - 繰り返しの魔法！

### 🎯 基本的なループ - 単純な繰り返しを極めろ！

```yaml
{{ includex("examples/advanced/loop_basic.yml") }}
```

### 🔄 条件付きループ（リトライ機能） - 諸めないテストの極意！

```yaml
{{ includex("examples/advanced/loop_retry.yml") }}
```

### 🎮 動的なループ制御 - ループを思いのままに操れ！

```yaml
{{ includex("examples/advanced/loop_dynamic.yml") }}
```

## 🔀 条件付き実行 - 賢いテストの秘訣！

### 🎆 基本的な条件分岐 - if文でテストを制御！

```yaml
{{ includex("examples/advanced/conditional_basic.yml") }}
```

### 🧠 複雑な条件式 - どんな条件も表現できる！

```yaml
{{ includex("examples/advanced/conditional_complex.yml") }}
```

### 🛡️ エラーハンドリングと条件分岐 - 失敗を成功に変えろ！

```yaml
{{ includex("examples/advanced/conditional_error_handling.yml") }}
```

## 📦 シナリオのインクルード - DRYの極意！

### 🔗 基本的なインクルード - シナリオを再利用せよ！

```yaml
{{ includex("examples/advanced/include_basic.yml") }}
```

```yaml
{{ includex("examples/advanced/common/auth.yml") }}
```

### 🎭 動的なインクルード - 実行時にシナリオを選択！

```yaml
{{ includex("examples/advanced/include_dynamic.yml") }}
```

### 🏢 ネストしたインクルード - 階層的なシナリオ構築！

```yaml
{{ includex("examples/advanced/include_nested.yml") }}
```

```yaml
{{ includex("examples/advanced/level2.yml") }}
```

## ⚡ 並行実行制御 - スピードの限界を突破！

### 🚀 基本的な並行実行 - 複数テストを同時実行！

```yaml
{{ includex("examples/advanced/concurrency_basic.yml") }}
```

### 🔒 共有リソースの制御 - 競合を防ぐ鉄壁の守り！

```yaml
{{ includex("examples/advanced/concurrency_shared_resource.yml") }}
```

### 🎆 複雑な並行制御 - プロフェッショナルの技！

```yaml
{{ includex("examples/advanced/concurrency_complex.yml") }}
```

## 🌐 依存関係の定義 - テストの流れを完全制御！

### 🔗 基本的な依存関係 - 順番を守る賢いテスト！

```yaml
{{ includex("examples/advanced/dependency_basic.yml") }}
```

### 🕸️ 複雑な依存関係 - ネットワーク状の依存も余裕！

```yaml
{{ includex("examples/advanced/dependency_complex.yml") }}
```

## 🔧 カスタムランナーの作成 - 独自のランナーを生み出せ！

### 🔌 プラグインランナーの定義 - runnを拡張せよ！

```yaml
{{ includex("examples/advanced/custom_runner_plugin.yml") }}
```

### 🌍 外部コマンドランナー - どんなツールも統合！

```yaml
{{ includex("examples/advanced/custom_runner_external.yml") }}
```

## 🎭 高度なデータ処理 - データを魔法のように操れ！

### 🎲 動的なテストデータ生成 - 無限のテストデータ！

```yaml
{{ includex("examples/advanced/data_generation.yml") }}
```

### 🌀 複雑なデータ変換 - データの形を自在に変えろ！

```yaml
{{ includex("examples/advanced/data_transformation.yml") }}
```

## 🛡️ エラーハンドリングとデバッグ - トラブルシューティングの達人！

### 💜 包括的なエラーハンドリング - どんなエラーも恐れない！

```yaml
{{ includex("examples/advanced/error_handling.yml") }}
```

### 🔍 デバッグ情報の出力 - 問題を瞬時に特定！

```yaml
{{ includex("examples/advanced/debug_output.yml") }}
```

## 🎆 まとめ - 高度な機能のマスター誕生！

**やったぞ！** あなたは今、**runnの真の力を解放**した！

### 🏆 この章で手に入れた8つの超能力：

1. **🔁 ループ処理**: 繰り返しを**完全に支配**！リトライも思いのまま！
2. **🔀 条件付き実行**: 賢いテストが**自分で判断**！
3. **📦 シナリオのインクルード**: DRY原則を**極限まで追求**！
4. **⚡ 並行実行制御**: スピードの**限界を突破**！
5. **🌐 依存関係の定義**: 複雑なフローも**美しく制御**！
6. **🔧 カスタムランナー**: あなただけの**特別なランナー**を作れ！
7. **🎭 高度なデータ処理**: データを**魔法のように変形**！
8. **🛡️ エラーハンドリング**: どんなエラーも**恐れない**！

これらの機能を**絶妙に組み合わせれば**、あなたのテストは**芸術作品**に昇華する。もう、どんなに複雑なプロダクション環境も**完璧にテスト**できる！

**次章では、runnの最終兵器、Goテストヘルパーとしての使い方を伝授しよう！** これであなたも**真のrunnマスター**だ！

[第7章：Goテストヘルパー編へ →](test-helper.md)