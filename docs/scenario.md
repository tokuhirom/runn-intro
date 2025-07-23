# 第2章：シナリオ記述編

runnでは、YAMLを使ってテストシナリオを記述します。プログラミング知識がなくても、チーム全員が読み書きできるテストが作成できます。

## Runbook - あなたのテストの台本

runnでは、テストシナリオをRunbookと呼びます。何をどの順番で実行するかを記述したものです。

### Runbookの利点

従来のテストコード：
```javascript
// 読むのも書くのも大変...
const response = await fetch('/users', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ name: 'Alice' })
});
assert(response.status === 201);
```

**runnのRunbook：**
```yaml
{{ includex("examples/scenario/basic-structure.yml") }}
```

YAMLで記述することで、より直感的にテストを表現できます。

## Runbookの5つの要素

Runbookは5つの主要セクションで構成されています。

### 1. desc - シナリオの目的
テストの目的を明確に記述します。

### 2. labels - テストの分類
`--label`オプションで特定のテストのみを実行できます。

### 3. runners - 接続先の定義
HTTP、gRPC、データベースなどの接続先を定義します。

### 4. vars - 変数定義
繰り返し使う値を変数として定義できます。

### 5. steps - 実行ステップ
テストの本体となる、順番に実行されるアクションのリストです。

## 2つの記述スタイル

runnは2つの記述スタイルを提供しています。

### スタイル1: リスト形式

インデックスでステップを参照する、シンプルな方法です。

```yaml
{{ includex("examples/scenario/list-format.yml") }}
```

メリット：
- 学習が簡単
- 短いシナリオに適している
- `steps[0]`、`steps[1]`と直感的に参照できる

### スタイル2: マップ形式

名前付きステップで、より読みやすくなります。

```yaml
{{ includex("examples/scenario/map-format.yml") }}
```

メリット：
- ステップの意図が明確
- 長いシナリオでも管理しやすい
- `steps.create_user`のように意味のある名前で参照できる

10ステップ以上のシナリオでは、マップ形式の使用をお勧めします。

## 変数の活用

### 変数定義

runnの変数システムを使用すると、柔軟な設定が可能です。

```yaml
{{ includex("examples/scenario/variable-definition.yml") }}
```

ポイント：
- 環境変数から自動取得（`${API_KEY}`）
- デフォルト値の設定（`${ENV:-development}`）
- 複雑なデータ構造も定義可能

### 変数参照

定義した変数は以下のように参照できます。

```yaml
{{ includex("examples/scenario/variable-reference.yml") }}
```

`{% raw %}{{ vars.変数名 }}{% endraw %}`の記法で変数を参照できます。

## ステップ記述の詳細

### HTTPリクエストの例

HTTPリクエストの各種機能を使用した例です。

```yaml
{{ includex("examples/scenario/http-request-complete.yml") }}
```

標準出力：

```
{{ includex("examples/scenario/http-request-complete.stdout") }}
```

利用可能な機能：
- 条件付き実行（`if`）
- タイムアウト設定
- テストアサーション
- デバッグ用のダンプ機能

### データベースクエリ

SQLクエリもYAMLで記述できます。

```yaml
{{ includex("examples/scenario/database-query.concept.yml") }}
```

HTTPもデータベースも、同じ`test`構文でアサーションを記述できます。

## 実践的なシナリオ例

### CRUD操作の実装

完全なCRUD操作を実装した例です。

```yaml
{{ includex("examples/scenario/crud-operations.yml") }}
```

```
{{ includex("examples/scenario/crud-operations.out") }}
```

この例では以下を実装しています：
- Create → Read → Update → Delete のフロー
- ステップ間でのID受け渡し
- 削除後の404確認

## YAML記述のテクニック

### テクニック1: 複数行文字列

長いテキストの記述方法です。

```yaml
{{ includex("examples/scenario/multiline-strings.yml") }}
```

### テクニック2: アンカー＆エイリアス

共通設定を再利用する方法です。

```yaml
{{ includex("examples/scenario/anchors-aliases.yml") }}
```

共通設定を一箇所で管理することで、メンテナンスが容易になります。

### テクニック3: 環境別設定

開発・本番環境の設定を切り替える方法です。

```yaml
{{ includex("examples/scenario/environment-config.yml") }}
```

## まとめ

この章では、以下の内容を学びました。

- Runbookの5つの要素
- リスト形式とマップ形式の使い分け
- 変数を活用したシナリオ作成
- 実践的なシナリオの作成方法
- YAMLの便利なテクニック

次章では、runnの式評価エンジンについて学びます。前のステップの結果を参照したり、条件分岐やフィルタリングを行う方法を説明します。