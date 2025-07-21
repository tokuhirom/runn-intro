# 第9章：リファレンス - 究極の参考書！

**ついに最終章だ！** これまでの旅で身につけた**すべての知識を完璧に整理**しよう！この章は、**runnマスターの聖典**として、いつでも参照できる**最強のリファレンス**だ。困ったときは、ここに戻ってこい！

## 📜 YAMLスキーマ - 設計図の完全マスター！

### 📚 Runbookの基本構造 - すべての始まり！

```yaml
{{ includex("examples/chapter09/runbook_structure.yml") }}
```

### 🎯 Step（ステップ）の構造 - アクションの設計図！

```yaml
{{ includex("examples/chapter09/step_structure.yml") }}
```

### 🏃 Runner（ランナー）の定義 - プロトコルの司令塔！

```yaml
{{ includex("examples/chapter09/runner_definition.yml") }}
```

### 🔁 Loop（ループ）の設定 - 繰り返しの魔法！

```yaml
{{ includex("examples/chapter09/loop_settings.yml") }}
```

### 🌐 HTTPRequest（HTTPリクエスト）の構造 - Web APIの設計図！

```yaml
{{ includex("examples/chapter09/http_request_structure.yml") }}
```

### 🗄️ DBQuery（データベースクエリ）の構造 - SQLの秘伝書！

```yaml
{{ includex("examples/chapter09/db_query_structure.yml") }}
```

### 🎮 CDPAction（ブラウザ操作）の構造 - ブラウザ制御の極意！

```yaml
{{ includex("examples/chapter09/cdp_action_structure.yml") }}
```

## 🎆 全ビルトイン関数一覧 - 最強の武器庫！

**これが全武器のカタログだ！** どんな状況でも、**必要な関数が必ず見つかる**！

### 🔍 比較・差分関数 - 違いを見逃さない！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `compare(x, y, ignorePaths...)` | 2つの値を比較 | `compare(actual, expected)` |
| `diff(x, y, ignorePaths...)` | 2つの値の差分を表示 | `diff(actual, expected, [".timestamp"])` |

### 🎭 データ操作関数 - データの魔術師！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `pick(obj, keys...)` | 指定キーのみ抽出 | `pick(user, "id", "name")` |
| `omit(obj, keys...)` | 指定キーを除外 | `omit(user, "password")` |
| `merge(obj1, obj2...)` | オブジェクトをマージ | `merge(defaults, config)` |
| `intersect(arr1, arr2)` | 配列の積集合 | `intersect([1,2,3], [2,3,4])` |
| `union(arr1, arr2)` | 配列の和集合 | `union([1,2], [2,3])` |
| `unique(arr)` | 重複を除去 | `unique([1,2,2,3])` |
| `groupBy(arr, expr)` | 配列をグループ化 | `groupBy(users, {.department})` |

### 📝 配列・文字列関数 - テキストの支配者！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `len(x)` | 長さを取得 | `len(array)`, `len(string)` |
| `map(arr, expr)` | 配列の各要素を変換 | `map(users, {.name})` |
| `filter(arr, expr)` | 配列をフィルタリング | `filter(users, {.active})` |
| `sort(arr, expr)` | 配列をソート | `sort(users, {.name})` |
| `reverse(arr)` | 配列を逆順 | `reverse([1,2,3])` |
| `join(arr, sep)` | 配列を文字列に結合 | `join(names, ", ")` |
| `split(str, sep)` | 文字列を分割 | `split("a,b,c", ",")` |
| `contains(str, substr)` | 文字列が含まれるか | `contains(text, "hello")` |
| `startsWith(str, prefix)` | 文字列が指定文字で始まるか | `startsWith(text, "http")` |
| `endsWith(str, suffix)` | 文字列が指定文字で終わるか | `endsWith(text, ".com")` |
| `upper(str)` | 大文字に変換 | `upper("hello")` |
| `lower(str)` | 小文字に変換 | `lower("HELLO")` |
| `trim(str)` | 前後の空白を除去 | `trim(" hello ")` |

### 🔢 数値・計算関数 - 計算の達人！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `sum(arr)` | 配列の合計 | `sum([1,2,3])` |
| `min(arr)` | 配列の最小値 | `min([1,2,3])` |
| `max(arr)` | 配列の最大値 | `max([1,2,3])` |
| `avg(arr)` | 配列の平均値 | `avg([1,2,3])` |
| `abs(num)` | 絶対値 | `abs(-5)` |
| `ceil(num)` | 切り上げ | `ceil(3.14)` |
| `floor(num)` | 切り下げ | `floor(3.14)` |
| `round(num)` | 四捨五入 | `round(3.14)` |

### 🔄 型変換関数 - 変身の魔法！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `string(x)` | 文字列に変換 | `string(123)` |
| `int(x)` | 整数に変換 | `int("123")` |
| `float(x)` | 浮動小数点数に変換 | `float("3.14")` |
| `bool(x)` | 真偽値に変換 | `bool("true")` |
| `toJSON(x)` | JSON文字列に変換 | `toJSON(object)` |
| `fromJSON(str)` | JSON文字列から変換 | `fromJSON(jsonStr)` |

### 🔐 エンコーディング関数 - 暗号化の極意！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `urlencode(str)` | URLエンコード | `urlencode("hello world")` |
| `urldecode(str)` | URLデコード | `urldecode("hello%20world")` |
| `toBase64(str)` | Base64エンコード | `toBase64("hello")` |
| `fromBase64(str)` | Base64デコード | `fromBase64("aGVsbG8=")` |
| `toHex(str)` | 16進数エンコード | `toHex("hello")` |
| `fromHex(str)` | 16進数デコード | `fromHex("68656c6c6f")` |

### ⏰ 時間関数 - 時を操る力！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `time(x)` | 時刻オブジェクトに変換 | `time("2024-01-01T00:00:00Z")` |
| `now()` | 現在時刻 | `now()` |
| `unix(t)` | Unix timestamp | `unix(time("2024-01-01"))` |
| `format(t, layout)` | 時刻をフォーマット | `format(now(), "2006-01-02")` |
| `parse(str, layout)` | 文字列を時刻に変換 | `parse("2024-01-01", "2006-01-02")` |

### 📁 ファイル・URL関数 - 外部リソースの架け橋！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `file(path)` | ファイルを読み込み | `file("data.json")` |
| `url(rawURL)` | URLを解析 | `url("https://example.com/path")` |

### 🎲 Faker関数（テストデータ生成） - 無限のデータ工場！

#### 👤 個人情報 - リアルな人物データ！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `faker.name()` | ランダムな名前 | `faker.name()` |
| `faker.firstName()` | 名前（名） | `faker.firstName()` |
| `faker.lastName()` | 名前（姓） | `faker.lastName()` |
| `faker.email()` | メールアドレス | `faker.email()` |
| `faker.phone()` | 電話番号 | `faker.phone()` |
| `faker.username()` | ユーザー名 | `faker.username()` |

#### 🏠 住所情報 - 世界中の住所を生成！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `faker.address()` | 住所 | `faker.address()` |
| `faker.city()` | 都市名 | `faker.city()` |
| `faker.state()` | 州・県名 | `faker.state()` |
| `faker.country()` | 国名 | `faker.country()` |
| `faker.zipCode()` | 郵便番号 | `faker.zipCode()` |

#### 🌐 インターネット関連 - ネットワークデータの宝庫！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `faker.url()` | URL | `faker.url()` |
| `faker.domainName()` | ドメイン名 | `faker.domainName()` |
| `faker.ipv4()` | IPv4アドレス | `faker.ipv4()` |
| `faker.ipv6()` | IPv6アドレス | `faker.ipv6()` |
| `faker.macAddress()` | MACアドレス | `faker.macAddress()` |

#### 🎯 数値・文字列 - ランダムデータの魔法！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `faker.number(min, max)` | ランダムな数値 | `faker.number(1, 100)` |
| `faker.float(min, max, precision)` | ランダムな浮動小数点数 | `faker.float(0.0, 1.0, 2)` |
| `faker.uuid()` | UUID | `faker.uuid()` |
| `faker.randomString(length)` | ランダムな文字列 | `faker.randomString(10)` |
| `faker.password(lower, upper, numeric, special, space, length)` | パスワード | `faker.password(true, true, true, false, false, 12)` |

#### 📅 日時 - 時間データの生成器！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `faker.date()` | ランダムな日付 | `faker.date()` |
| `faker.dateRange(start, end)` | 期間内のランダムな日付 | `faker.dateRange("2024-01-01", "2024-12-31")` |
| `faker.time()` | ランダムな時刻 | `faker.time()` |

#### 🎰 選択・真偽値 - 運命の選択！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `faker.randomChoice(arr)` | 配列からランダム選択 | `faker.randomChoice(["A", "B", "C"])` |
| `faker.randomBool()` | ランダムな真偽値 | `faker.randomBool()` |
| `faker.randomInt(min, max)` | ランダムな整数 | `faker.randomInt(1, 100)` |

### ✨ その他の関数 - 特殊部隊！

| 関数名 | 説明 | 使用例 |
|--------|------|--------|
| `input(prompt, default)` | ユーザー入力を取得 | `input("Enter name", "default")` |
| `env(key, default)` | 環境変数を取得 | `env("HOME", "/tmp")` |
| `range(start, end, step)` | 数値の範囲を生成 | `range(1, 10, 2)` |
| `keys(obj)` | オブジェクトのキー一覧 | `keys({"a": 1, "b": 2})` |
| `values(obj)` | オブジェクトの値一覧 | `values({"a": 1, "b": 2})` |

## 🚨 エラーメッセージ一覧 - トラブルシューティングの極意！

**エラーは敵じゃない、成長のチャンスだ！** この一覧で、**どんなエラーも瞬時に解決**！

### ⚠️ 一般的なエラー - 基本の対処法！

| エラーメッセージ | 原因 | 対処法 |
|------------------|------|--------|
| `failed to load runbook` | YAMLファイルの読み込みエラー | ファイルパスとYAML構文を確認 |
| `invalid YAML format` | YAML構文エラー | YAMLの構文を確認（インデント、引用符など） |
| `step not found` | 存在しないステップを参照 | ステップ名のスペルミスを確認 |
| `runner not found` | 存在しないランナーを参照 | ランナー名のスペルミスを確認 |

### 🌐 HTTP関連エラー - Web通信のトラブル解決！

| エラーメッセージ | 原因 | 対処法 |
|------------------|------|--------|
| `connection refused` | サーバーに接続できない | サーバーが起動しているか確認 |
| `timeout` | リクエストがタイムアウト | タイムアウト値を調整またはサーバー応答を確認 |
| `invalid URL` | 不正なURL形式 | URL形式を確認 |
| `unsupported content type` | サポートされていないContent-Type | 適切なContent-Typeを指定 |

### 🗄️ データベース関連エラー - DB問題の即効薬！

| エラーメッセージ | 原因 | 対処法 |
|------------------|------|--------|
| `database connection failed` | データベース接続エラー | 接続文字列とデータベース状態を確認 |
| `SQL syntax error` | SQL構文エラー | SQLクエリの構文を確認 |
| `table not found` | テーブルが存在しない | テーブル名とスキーマを確認 |
| `permission denied` | 権限不足 | データベースユーザーの権限を確認 |

### 🧮 式評価エラー - 計算ミスを撲滅！

| エラーメッセージ | 原因 | 対処法 |
|------------------|------|--------|
| `expression evaluation failed` | 式の評価エラー | 式の構文と変数の存在を確認 |
| `variable not found` | 変数が存在しない | 変数名のスペルミスと定義を確認 |
| `type mismatch` | 型の不一致 | 期待される型と実際の型を確認 |
| `division by zero` | ゼロ除算エラー | 除数がゼロでないことを確認 |

### 🎮 ブラウザ操作エラー - UI操作の問題解決！

| エラーメッセージ | 原因 | 対処法 |
|------------------|------|--------|
| `element not found` | 要素が見つからない | セレクターと要素の存在を確認 |
| `element not visible` | 要素が表示されていない | 要素の表示状態を確認 |
| `browser not responding` | ブラウザが応答しない | ブラウザの状態とリソース使用量を確認 |
| `navigation failed` | ページ遷移に失敗 | URLとネットワーク接続を確認 |

## 💡 FAQ（よくある質問） - みんなの疑問を完全解決！

**先輩たちの知恵が詰まった宝箱！** ここには、**実践で役立つ回答**が満載だ！

### 🎯 基本的な使い方 - 初心者の疑問を解消！

**Q: runnとは何ですか？**
A: runnは、シナリオベースのテスト・自動化ツールです。YAMLでテストシナリオを記述し、HTTP、gRPC、データベース、ブラウザ操作などを統一的に扱えます。

**Q: CLIツールとGoテストヘルパーの違いは何ですか？**
A: CLIツールは単独でシナリオを実行しますが、Goテストヘルパーは`go test`と統合してより柔軟なテスト環境を構築できます。本書では特にGoテストヘルパーとしての利用を推奨しています。

**Q: どのようなプロトコルをサポートしていますか？**
A: HTTP/HTTPS、gRPC、データベース（PostgreSQL、MySQL、SQLite、Cloud Spanner）、ブラウザ操作（CDP）、SSH、ローカルコマンド実行をサポートしています。

### 🔧 インストールと設定 - セットアップの悩みを解決！

**Q: インストール方法を教えてください。**
A: 複数の方法があります：
- Homebrew: `brew install k1LoW/tap/runn`
- Go install: `go install github.com/k1LoW/runn/cmd/runn@latest`
- バイナリダウンロード: GitHub Releasesから取得

**Q: Goプロジェクトに統合するにはどうすればよいですか？**
A: `go.mod`に`github.com/k1LoW/runn`を追加し、テストファイルで`runn.Load()`を使用してシナリオを実行します。

### 📝 シナリオ記述 - 書き方のコツを伝授！

**Q: リスト形式とマップ形式のどちらを使うべきですか？**
A: シンプルなシナリオはリスト形式、複雑なシナリオや可読性を重視する場合はマップ形式を推奨します。

**Q: 変数はどのように定義・参照しますか？**
A: `vars:`セクションで定義し、`{% raw %}{{ vars.変数名 }}{% endraw %}`で参照します。環境変数は`{% raw %}{{ env.環境変数名 }}{% endraw %}`で参照できます。

**Q: 前のステップの結果を次のステップで使用するには？**
A: `steps.ステップ名.res.body`（マップ形式）または`steps[インデックス].res.body`（リスト形式）で参照します。

### 🛡️ エラーハンドリング - エラーを味方に！

**Q: テストが失敗した時の詳細情報を取得するには？**
A: `dump:`セクションを使用してデバッグ情報を出力し、`test: true`でエラーでも続行させることができます。

**Q: リトライ機能はありますか？**
A: `loop:`セクションで`until:`条件と`minInterval:`、`maxInterval:`を指定することでリトライ機能を実装できます。

**Q: 条件付きでステップを実行するには？**
A: `if:`フィールドに条件式を記述することで、条件付き実行が可能です。

### ⚡ パフォーマンス - 速度の限界突破！

**Q: 並列実行は可能ですか？**
A: `concurrency:`フィールドで並列実行数を制御できます。共有リソースを使用する場合は同じキーを指定して順次実行も可能です。

**Q: 大量のテストデータを効率的に処理するには？**
A: `loop:`を使用した繰り返し処理、Faker関数による動的データ生成、外部ファイルからのデータ読み込みを組み合わせます。

**Q: パフォーマンステストは実行できますか？**
A: 可能です。`loop:`で大量のリクエストを生成し、レスポンス時間やスループットを測定できます。

### 🔗 統合・連携 - システムとの完璧な融合！

**Q: CI/CDパイプラインで実行するには？**
A: GitHub ActionsやJenkinsなどで`go test`コマンドを実行するだけです。環境変数でテスト環境を切り替えることも可能です。

**Q: 既存のテストフレームワークと併用できますか？**
A: はい。runnはGoの標準的なテストフレームワークと統合されているため、既存のテストと併用できます。

**Q: モックサーバーとの連携は可能ですか？**
A: 可能です。`httptest.NewServer()`で作成したモックサーバーのURLをランナーに設定することで連携できます。

### 🔧 トラブルシューティング - 問題解決の達人技！

**Q: 「runner not found」エラーが発生します。**
A: `runners:`セクションでランナーが正しく定義されているか、ステップでのランナー名にスペルミスがないか確認してください。

**Q: データベース接続エラーが発生します。**
A: 接続文字列の形式、データベースサーバーの起動状態、ネットワーク接続、認証情報を確認してください。

**Q: ブラウザ操作で要素が見つからないエラーが発生します。**
A: セレクターが正しいか、要素が表示されるまで適切に待機しているか、ページの読み込みが完了しているかを確認してください。

**Q: 式評価でエラーが発生します。**
A: 変数名のスペルミス、変数の存在、型の不一致、構文エラーを確認してください。`dump:`を使用してデバッグ情報を出力することも有効です。

### 🏆 ベストプラクティス - プロの流儀！

**Q: テストデータはどのように管理すべきですか？**
A: 固定データは再現性のために、ランダムデータは多様性のために使い分けます。外部ファイルからの読み込みも活用しましょう。

**Q: 大規模なプロジェクトでの構成はどうすべきですか？**
A: 機能別にディレクトリを分け、共通処理は`include:`で再利用し、環境別設定を適切に管理することを推奨します。

**Q: セキュリティ面で注意すべき点はありますか？**
A: 認証情報は環境変数で管理し、ログに機密情報が出力されないよう注意し、テスト用データベースは本番から分離してください。

## 📊 バージョン情報と互換性 - 環境の完全ガイド！

### ✅ サポートされるバージョン - 動作保証環境！

- **runn**: v0.100.0以降
- **Go**: 1.21以降
- **PostgreSQL**: 12以降
- **MySQL**: 8.0以降
- **Chrome/Chromium**: 90以降

### 📈 変更履歴の重要なポイント - 進化の軌跡！

- v0.100.0: 安定版リリース、APIの大幅な変更
- v0.90.0: Goテストヘルパー機能の強化
- v0.80.0: CDP（ブラウザ操作）サポート追加
- v0.70.0: 並行実行制御機能追加

## 🔗 参考リンク - さらなる探求への道！

- **公式リポジトリ**: https://github.com/k1LoW/runn
- **ドキュメント**: https://github.com/k1LoW/runn/tree/main/docs
- **サンプル**: https://github.com/k1LoW/runn/tree/main/testdata
- **Issue報告**: https://github.com/k1LoW/runn/issues
- **expr-lang/expr**: https://github.com/expr-lang/expr

## 🎆 まとめ - リファレンスマスター誕生！

**やったぞ！** あなたは今、**runnの完全なリファレンスを手に入れた**！

### 🏆 この章で獲得した4つの最強武器：

1. **📜 YAMLスキーマ**: Runbook、Step、Runnerの**設計図を完全理解**！
2. **🎆 全ビルトイン関数一覧**: 100以上の関数を**いつでも引き出せる**！
3. **🚨 エラーメッセージ一覧**: どんなエラーも**瞬時に解決**！
4. **💡 FAQ**: 実践的な疑問を**すべて解消**！

このリファレンスは、**あなたの最強の相棒**だ。困ったときは、いつでもここに戻ってこい！runnは**日々進化している**から、公式リポジトリも要チェックだ！

---

# 🎉 本書の完了 - あなたはrunnマスターだ！

**おめでとう！** ついに「runn入門」**全9章を完全制覇**した！

## 🚀 あなたが手に入れた力

1. **基礎編**: runnの基本を**完璧にマスター**！
2. **シナリオ記述編**: YAMLで**美しいテストを設計**！
3. **Expression文法編**: 式の力で**データを自在に操作**！
4. **ビルトイン関数編**: 100以上の関数を**使いこなす**！
5. **ランナー詳細編**: 6大プロトコルを**完全制圧**！
6. **高度な機能編**: ループ、条件分岐で**複雑なテストも余裕**！
7. **Goテストヘルパー編**: **最強の統合テスト環境**を構築！
8. **実践編**: 現場で**即戦力のテクニック**を習得！
9. **リファレンス**: いつでも参照できる**最強の辞書**！

特に第7章で学んだ**Goテストヘルパー**こそが、**runnの真の姿**だ。YAMLの**美しさ**とGoの**パワー**が融合し、**史上最強のテストスイート**が誕生する！

## 🎆 次のステップへ

あなたはもう、**ただのテスターではない**。**テストアーキテクト**として、**品質の守護者**として、プロジェクトを成功に導く**ヒーロー**だ！

**さあ、実践の舞台へ飛び出そう！** runnという**最強の武器**を手に、**テスト自動化の新時代**を切り拓け！

**Happy Testing with runn! 🚀**
