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
steps:
  compare_example:
    dump: |
      {
        "expected": {"name": "Alice", "age": 30, "city": "Tokyo"},
        "actual": {"name": "Alice", "age": 31, "country": "Japan"}
      }
    test: |
      # compare関数で差分を検出
      compare(current.expected, current.actual) != null
    
    dump_diff:
      # 差分の詳細を表示
      diff: compare(current.expected, current.actual)
```

### diff関数

文字列やJSONの差分を人間が読みやすい形式で出力します。

```yaml
steps:
  diff_example:
    dump: |
      {
        "old": "Hello\nWorld\nTest",
        "new": "Hello\nPlanet\nTest"
      }
    dump_diff:
      # テキストの差分
      text_diff: diff(current.old, current.new)
      # JSONの差分
      json_diff: diff(
        {"users": ["Alice", "Bob"]},
        {"users": ["Alice", "Charlie"], "count": 2}
      )
```

## データ操作関数

### pick関数

オブジェクトから指定したキーのみを抽出します。

```yaml
vars:
  user:
    id: 1
    name: "Alice"
    email: "alice@example.com"
    password: "secret"
    created_at: "2024-01-01"
steps:
  pick_example:
    dump: |
      # パスワードを除外してユーザー情報を抽出
      pick(vars.user, ["id", "name", "email"])
    test: |
      !("password" in current) &&
      current.name == "Alice"
```

### omit関数

オブジェクトから指定したキーを除外します。

```yaml
steps:
  omit_example:
    dump: |
      # センシティブな情報を除外
      omit(vars.user, ["password", "created_at"])
    test: |
      !("password" in current) &&
      !("created_at" in current) &&
      current.name == "Alice"
```

### merge関数

複数のオブジェクトをマージします。後のオブジェクトが優先されます。

```yaml
vars:
  defaults:
    timeout: 30
    retries: 3
    debug: false
  custom:
    timeout: 60
    verbose: true
steps:
  merge_example:
    dump: |
      # デフォルト設定とカスタム設定をマージ
      merge(vars.defaults, vars.custom)
    test: |
      current.timeout == 60 &&  # customの値で上書き
      current.retries == 3 &&   # defaultsの値を維持
      current.debug == false && # defaultsの値を維持
      current.verbose == true   # customの新しい値
```

### intersect関数

複数の配列やオブジェクトの共通部分を返します。

```yaml
vars:
  list1: ["apple", "banana", "orange", "grape"]
  list2: ["banana", "grape", "melon"]
  list3: ["grape", "apple", "banana"]
steps:
  intersect_example:
    dump: |
      {
        "two_arrays": intersect(vars.list1, vars.list2),
        "three_arrays": intersect(vars.list1, vars.list2, vars.list3),
        "objects": intersect(
          {"a": 1, "b": 2, "c": 3},
          {"b": 2, "c": 3, "d": 4}
        )
      }
    test: |
      current.two_arrays == ["banana", "grape"] &&
      current.three_arrays == ["banana", "grape"] &&
      current.objects == {"b": 2, "c": 3}
```

### groupBy関数

配列の要素を指定した条件でグループ化します。

```yaml
vars:
  users:
    - name: "Alice"
      role: "admin"
      department: "IT"
    - name: "Bob"
      role: "user"
      department: "Sales"
    - name: "Charlie"
      role: "admin"
      department: "IT"
    - name: "David"
      role: "user"
      department: "Sales"
steps:
  groupby_example:
    dump: |
      {
        "byRole": groupBy(vars.users, {.role}),
        "byDepartment": groupBy(vars.users, {.department})
      }
    test: |
      len(current.byRole.admin) == 2 &&
      len(current.byRole.user) == 2 &&
      len(current.byDepartment.IT) == 2 &&
      len(current.byDepartment.Sales) == 2
```

### flatten関数

ネストした配列を平坦化します。

```yaml
steps:
  flatten_example:
    dump: |
      {
        "nested": [[1, 2], [3, 4], [5]],
        "flattened": flatten([[1, 2], [3, 4], [5]]),
        "deep_nested": [[[1]], [[2, 3]], [4]],
        "deep_flattened": flatten([[[1]], [[2, 3]], [4]])
      }
    test: |
      current.flattened == [1, 2, 3, 4, 5] &&
      current.deep_flattened == [[1], [2, 3], 4]
```

### unique関数

配列から重複を除去します。

```yaml
steps:
  unique_example:
    dump: |
      {
        "numbers": unique([1, 2, 2, 3, 3, 3, 4]),
        "strings": unique(["apple", "banana", "apple", "orange", "banana"]),
        "objects": unique([
          {"id": 1, "name": "A"},
          {"id": 2, "name": "B"},
          {"id": 1, "name": "A"}
        ])
      }
    test: |
      current.numbers == [1, 2, 3, 4] &&
      current.strings == ["apple", "banana", "orange"] &&
      len(current.objects) == 2
```

## 文字列処理関数

### 基本的な文字列関数

```yaml
steps:
  string_basics:
    dump: |
      {
        "upper": upper("hello world"),
        "lower": lower("HELLO WORLD"),
        "trim": trim("  hello  "),
        "trimPrefix": trimPrefix("Mr. Smith", "Mr. "),
        "trimSuffix": trimSuffix("test.txt", ".txt"),
        "split": split("a,b,c", ","),
        "join": join(["a", "b", "c"], "-"),
        "replace": replace("hello world", "world", "runn"),
        "replaceAll": replaceAll("foo bar foo", "foo", "baz")
      }
    test: |
      current.upper == "HELLO WORLD" &&
      current.lower == "hello world" &&
      current.trim == "hello" &&
      current.trimPrefix == "Smith" &&
      current.trimSuffix == "test" &&
      current.split == ["a", "b", "c"] &&
      current.join == "a-b-c" &&
      current.replace == "hello runn" &&
      current.replaceAll == "baz bar baz"
```

### 高度な文字列関数

```yaml
steps:
  string_advanced:
    dump: |
      {
        "contains": contains("hello world", "world"),
        "startsWith": startsWith("hello world", "hello"),
        "endsWith": endsWith("hello.txt", ".txt"),
        "matches": matches("test@example.com", "^[a-z]+@[a-z]+\\.[a-z]+$"),
        "substr": substr("hello world", 6, 5),
        "indexOf": indexOf("hello world", "world"),
        "lastIndexOf": lastIndexOf("hello world world", "world"),
        "repeat": repeat("*", 5)
      }
    test: |
      current.contains == true &&
      current.startsWith == true &&
      current.endsWith == true &&
      current.matches == true &&
      current.substr == "world" &&
      current.indexOf == 6 &&
      current.lastIndexOf == 12 &&
      current.repeat == "*****"
```

### エンコーディング関数

```yaml
steps:
  encoding_example:
    dump: |
      {
        "urlEncode": urlencode("hello world&test=1"),
        "urlDecode": urldecode("hello%20world%26test%3D1"),
        "base64Encode": toBase64("hello world"),
        "base64Decode": fromBase64("aGVsbG8gd29ybGQ="),
        "jsonEncode": toJSON({"name": "Alice", "age": 30}),
        "jsonDecode": fromJSON('{"name":"Bob","age":25}')
      }
    test: |
      current.urlEncode == "hello+world%26test%3D1" &&
      current.urlDecode == "hello world&test=1" &&
      current.base64Encode == "aGVsbG8gd29ybGQ=" &&
      current.base64Decode == "hello world" &&
      current.jsonDecode.name == "Bob"
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
steps:
  time_example:
    dump: |
      {
        "now": time.now(),
        "today": time.format(time.now(), "2006-01-02"),
        "timestamp": time.unix(time.now()),
        "parsed": time.parse("2024-01-01 10:00:00", "2006-01-02 15:04:05"),
        "formatted": time.format(
          time.parse("2024-01-01T10:00:00Z", time.RFC3339),
          "January 2, 2006"
        ),
        "addHours": time.add(time.now(), time.hour * 2),
        "subDays": time.add(time.now(), -time.day * 7)
      }
```

### 時間の比較と計算

```yaml
steps:
  time_calculation:
    dump: |
      let start = time.parse("2024-01-01T00:00:00Z", time.RFC3339);
      let end = time.parse("2024-01-02T12:30:00Z", time.RFC3339);
      {
        "duration": time.sub(end, start),
        "hours": time.sub(end, start) / time.hour,
        "isAfter": time.after(end, start),
        "isBefore": time.before(start, end),
        "equal": time.equal(start, start)
      }
    test: |
      current.hours == 36.5 &&
      current.isAfter == true &&
      current.isBefore == true &&
      current.equal == true
```

## テストデータ生成（faker関数）

### 基本的なfaker関数

```yaml
steps:
  faker_basic:
    dump: |
      {
        "name": faker.name(),
        "email": faker.email(),
        "phone": faker.phoneNumber(),
        "address": faker.address(),
        "company": faker.company(),
        "jobTitle": faker.jobTitle(),
        "uuid": faker.uuid(),
        "url": faker.url()
      }
    test: |
      # 生成されたデータの基本的な検証
      contains(current.email, "@") &&
      len(current.uuid) == 36 &&
      startsWith(current.url, "http")
```

### 数値とランダムデータ

```yaml
steps:
  faker_numbers:
    dump: |
      {
        "randomInt": faker.randomInt(1, 100),
        "randomFloat": faker.randomFloat(2, 0, 1),
        "randomBool": faker.randomBool(),
        "randomString": faker.randomString(10),
        "randomChoice": faker.randomChoice(["apple", "banana", "orange"]),
        "shuffle": faker.shuffle([1, 2, 3, 4, 5])
      }
    test: |
      current.randomInt >= 1 && current.randomInt <= 100 &&
      current.randomFloat >= 0 && current.randomFloat <= 1 &&
      type(current.randomBool) == "bool" &&
      len(current.randomString) == 10 &&
      current.randomChoice in ["apple", "banana", "orange"] &&
      len(current.shuffle) == 5
```

### 日付関連のfaker関数

```yaml
steps:
  faker_dates:
    dump: |
      {
        "pastDate": faker.dateTime().past(1),  # 過去1年以内
        "futureDate": faker.dateTime().future(1),  # 未来1年以内
        "recentDate": faker.dateTime().recent(7),  # 過去7日以内
        "birthDate": faker.dateTime().birthday(20, 30),  # 20-30歳
        "between": faker.dateTime().between(
          "2024-01-01T00:00:00Z",
          "2024-12-31T23:59:59Z"
        )
      }
    test: |
      time.before(current.pastDate, time.now()) &&
      time.after(current.futureDate, time.now())
```

### ローカライズされたデータ

```yaml
steps:
  faker_localized:
    dump: |
      {
        # 日本語のデータ生成
        "jaName": faker.locale("ja").name(),
        "jaAddress": faker.locale("ja").address(),
        "jaCompany": faker.locale("ja").company(),
        # 他の言語
        "frName": faker.locale("fr").name(),
        "deName": faker.locale("de").name()
      }
```

## 実用的な使用例

### APIテストでの活用

```yaml
desc: ユーザー登録APIのテスト
steps:
  create_test_user:
    req:
      /users:
        post:
          body:
            application/json:
              name: faker.name()
              email: faker.email()
              age: faker.randomInt(20, 60)
              bio: faker.sentence()
              avatar: faker.imageURL()
    test: |
      current.res.status == 201 &&
      current.res.body.id != null
    
  verify_user:
    req:
      /users/{{ steps.create_test_user.res.body.id }}:
        get:
    test: |
      current.res.status == 200 &&
      # pickを使って比較対象を限定
      pick(current.res.body, ["name", "email"]) == 
      pick(steps.create_test_user.req.body, ["name", "email"])
```

### データ変換パイプライン

```yaml
vars:
  rawData:
    - timestamp: "2024-01-01T10:00:00Z"
      value: "123.45"
      tags: "important,urgent,todo"
    - timestamp: "2024-01-01T11:00:00Z"
      value: "67.89"
      tags: "normal,done"
steps:
  transform_data:
    dump: |
      map(vars.rawData, {
        # 時刻をフォーマット
        "date": time.format(
          time.parse(.timestamp, time.RFC3339),
          "2006-01-02"
        ),
        # 文字列を数値に変換
        "numericValue": float(.value),
        # タグを配列に分割
        "tagList": split(.tags, ","),
        # タグの数をカウント
        "tagCount": len(split(.tags, ","))
      })
    test: |
      current[0].date == "2024-01-01" &&
      current[0].numericValue == 123.45 &&
      current[0].tagList == ["important", "urgent", "todo"] &&
      current[0].tagCount == 3
```

### 複雑なデータ検証

```yaml
steps:
  complex_validation:
    req:
      /products:
        get:
    test: |
      # レスポンスの基本検証
      current.res.status == 200 &&
      
      # 全商品の価格が正の数
      all(current.res.body.products, {.price > 0}) &&
      
      # カテゴリごとに商品をグループ化して検証
      let grouped = groupBy(current.res.body.products, {.category});
      len(grouped.electronics) >= 5 &&
      len(grouped.books) >= 3 &&
      
      # 価格の統計情報を計算
      let prices = map(current.res.body.products, {.price});
      min(prices) >= 10 &&
      max(prices) <= 10000 &&
      
      # 在庫切れ商品の抽出
      let outOfStock = filter(current.res.body.products, {.stock == 0});
      len(outOfStock) < len(current.res.body.products) * 0.1  # 10%未満
```

## パフォーマンスとベストプラクティス

### 効率的な関数の使用

```yaml
steps:
  # 良い例：必要なデータのみを処理
  efficient_example:
    dump: |
      # 大きなオブジェクトから必要な部分のみ抽出
      let users = map(vars.largeDataset, {
        pick(., ["id", "name", "active"])
      });
      # アクティブユーザーのみフィルタ
      filter(users, {.active})
  
  # 避けるべき例：不必要な処理
  inefficient_example:
    dump: |
      # 全データを処理してからフィルタ（非効率）
      let allProcessed = map(vars.largeDataset, {
        merge(., {"processed": true, "timestamp": time.now()})
      });
      filter(allProcessed, {.active})
```

### エラーハンドリング

```yaml
steps:
  safe_operations:
    dump: |
      {
        # ファイル読み込みのエラーハンドリング
        "config": file("./config.json") ?? {"default": true},
        
        # 時刻パースのエラーハンドリング
        "parsedDate": time.parse(vars.dateString, "2006-01-02") ?? time.now(),
        
        # JSONパースのエラーハンドリング
        "data": fromJSON(vars.jsonString) ?? {}
      }
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