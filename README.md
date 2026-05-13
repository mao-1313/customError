# 課題：カスタムエラー型

Goのエラーハンドリングを体系的に理解する。

## 学習ステップ

### Step 1 - カスタムエラー型を定義する

```go
type NotFoundError struct {
    ID int
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("id %d not found", e.ID)
}
```

### Step 2 - エラーをラップする

```go
// %w でラップすると errors.Is / errors.As で辿れる
return fmt.Errorf("getUserByID: %w", &NotFoundError{ID: id})
```

### Step 3 - errors.Is / errors.As で判別する

```go
// errors.As: 特定の型かどうか確認して取り出す
var notFound *NotFoundError
if errors.As(err, &notFound) {
    fmt.Println("not found id:", notFound.ID)
}
```

### Step 4 - センチネルエラーを定義する

```go
var ErrUnauthorized = errors.New("unauthorized")

// errors.Is で一致確認
if errors.Is(err, ErrUnauthorized) { ... }
```

### Step 5 - テストを書く

`errors.As` / `errors.Is` が正しく動くかを確認する。

## 学びのポイント

- `error` interfaceの仕組み（`Error() string` を持つ型）
- ラップと `%w` の使い方
- `errors.Is`（同一性）と `errors.As`（型変換）の違い
- センチネルエラーのユースケース
