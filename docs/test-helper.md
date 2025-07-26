# 第5章：テストヘルパーとしての利用

runnはGoテストヘルパーとして使用でき、`go test`と統合してシナリオベースのテストを実行できます。

## 基本的な使い方

### プロジェクト構造

```
myproject/
├── main.go
├── main_test.go
├── go.mod
└── testdata/
    └── api_test.yml
```

### 実装例

以下は、シンプルなユーザーAPI のテスト例です。

`main.go`:
```go
{{ includex("examples/test-helper/simple/main.go") }}
```

`main_test.go`:
```go
{{ includex("examples/test-helper/simple/main_test.go") }}
```

`testdata/api_test.yml`:
```yaml
{{ includex("examples/test-helper/simple/testdata/api_test.yml") }}
```

### 実行方法

```bash
go test -v
```

## 主なオプション

- `runn.T(t)` - testing.Tを渡してテスト統合
- `runn.Runner("name", url)` - HTTPランナーの設定
- `runn.DBRunner("name", db)` - データベースランナーの設定
- `runn.Var("key", value)` - 変数の設定
- `runn.Debug(true)` - デバッグモードの有効化

## CI/CDでの実行

GitHub Actionsでの設定例:

`.github/workflows/test.yml`:
```yaml
{{ includex("examples/test-helper/simple/.github/workflows/test.yml") }}
```

この設定により、プッシュのたびに自動的にrunnを使用したテストが実行されます。