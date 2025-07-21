# 第2章：シナリオ記述編

> **「YAMLでテストを書く？そんなことできるの？」**
> 
> **「複雑なAPIの連携テストをシンプルに表現したい」**
> 
> **「チーム全員が読み書きできるテストシナリオが欲しい」**

**できます！** runnのシナリオ記述を学べば、**プログラミング知識がなくても**複雑なE2Eテストが書けるようになります。

## Runbook - あなたのテストの台本

> runnでは、テストシナリオを**Runbook**と呼びます。
> 
> まるで映画の台本のように、**何をどの順番で実行するか**を記述する、それがRunbookです。

### なぜRunbookなのか？

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

**見てください！** これだけ直感的に書けるんです。

## Runbookの5つの要素をマスターしよう

Runbookは**5つの主要セクション**で構成されています。それぞれが重要な役割を持っています：

### 1. **desc** - シナリオの目的を宣言
何をテストするのか、一目で分かるように記述します。**チームメンバーへの最高のドキュメント**です。

### 2. **labels** - スマートな分類
`--label`オプションで特定のテストだけを実行。**大規模プロジェクトの必須テクニック**！

### 3. **runners** - 接続先の定義
HTTP、gRPC、データベース...すべての接続先をここで定義。**マルチプロトコル対応の心臓部**！

### 4. **vars** - 変数で効率化
繰り返し使う値を変数化。**DRY原則をYAMLでも実現**！

### 5. **steps** - 実行ステップの記述
テストの本体。順番に実行される**アクションのリスト**です。

## 2つの記述スタイル - あなたはどっち派？

runnは**2つの記述スタイル**を提供しています。プロジェクトに合わせて選べる柔軟性！

### スタイル1: リスト形式 - シンプル＆ストレート

**インデックスでステップを参照**する、最もシンプルな方法：

```yaml
{{ includex("examples/scenario/list-format.yml") }}
```

**メリット：**
- 学習コストゼロ
- 短いシナリオに最適
- `steps[0]`、`steps[1]`...と直感的

### スタイル2: マップ形式 - 可読性の極み

**名前付きステップ**で、まるで関数のように扱える：

```yaml
{{ includex("examples/scenario/map-format.yml") }}
```

**メリット：**
- ステップの意図が明確
- 長いシナリオでも迷わない
- `steps.create_user`のように意味のある名前で参照

**プロのアドバイス：** 10ステップ以上のシナリオなら、迷わずマップ形式を選ぼう！

## 変数マスターへの道

### 変数定義の極意

runnの変数システムは**驚くほどパワフル**です：

```yaml
{{ includex("examples/scenario/variable-definition.yml") }}
```

**注目ポイント：**
- 環境変数から自動取得（`${API_KEY}`）
- デフォルト値の設定（`${ENV:-development}`）
- 複雑なデータ構造もOK

### 変数参照の魔法

定義した変数を**自在に活用**：

```yaml
{{ includex("examples/scenario/variable-reference.yml") }}
```

**ワンポイント：** `{% raw %}{{ vars.変数名 }}{% endraw %}`の記法で、どこでも変数を展開できます！

## 実践！ステップ記述の完全ガイド

### HTTPリクエストの全機能

**これがrunnの真骨頂！** すべての機能を詰め込んだ例：

```yaml
{{ includex("examples/scenario/http-request-complete.yml") }}
```

標準出力には以下のように出る！

```
{{ includex("examples/scenario/http-request-complete.stdout") }}
```

**驚きの機能群：**
- 条件付き実行（`if`）
- タイムアウト設定
- 複雑なテストアサーション
- デバッグ用のダンプ機能

### データベースも同じ感覚で

SQLクエリも**YAMLで自然に記述**：

```yaml
{{ includex("examples/scenario/database-query.concept.yml") }}
```

**ポイント：** HTTPもDBも、同じ`test`構文でアサーション！統一感が素晴らしい。

## 現場で使える！実践シナリオ集

### マスターピース1: 完璧なCRUD操作

**実際のプロジェクトで使えるレベル**のシナリオ：

```yaml
{{ includex("examples/scenario/crud-operations.yml") }}
```

```
{{ includex("examples/scenario/crud-operations.out") }}
```

**学べること：**
- Create → Read → Update → Delete の完全なフロー
- ステップ間でのID受け渡し
- 削除後の404確認まで網羅

## YAML記述の秘伝テクニック

### テクニック1: 複数行文字列

長いテキストも**美しく記述**：

```yaml
{{ includex("examples/scenario/multiline-strings.yml") }}
```

### テクニック2: アンカー＆エイリアス

**DRYの極み**を実現：

```yaml
{{ includex("examples/scenario/anchors-aliases.yml") }}
```

**効果：** 共通設定を一箇所で管理！メンテナンスが劇的に楽に。

### テクニック3: 環境別設定の管理

**開発・本番環境を賢く切り替え**：

```yaml
{{ includex("examples/scenario/environment-config.yml") }}
```

## あなたは今、YAMLマスター！

> **おめでとうございます！**
> 
> この章を読み終えたあなたは、もう立派な**Runbook作成者**です。

### 習得したスキル

✅ **Runbookの5つの要素**を完全理解

✅ **リスト形式とマップ形式**を使い分けられる

✅ **変数を活用**してDRYなシナリオが書ける

✅ **実践的なシナリオ**が作成できる

✅ **YAMLの高度なテクニック**を身につけた

### 次なる高みへ

でも、これはまだ**序章**です！

次章では、runnの**最強の武器**である式評価エンジンを学びます。前のステップの結果を自在に操り、条件分岐やフィルタリングを駆使する...

**もっとパワフルなテストが書けるようになります！**

---

*ヒント: この章で学んだ技術を組み合わせれば、どんな複雑なシナリオも表現できます。まずは小さなシナリオから始めて、徐々に大きくしていきましょう！*