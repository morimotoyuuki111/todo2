# 本日やったこと
- Mysqlを使う準備
- DBの準備
- Modelを作成する
- 会員一覧画面の修正
- 会員編集画面の修正
- 会員追加画面の修正

# エラー
list.go
```go
github.com/murai/go-user-management/model"でパッケージがインポート出来ませんでしたのエラー
```

```go
"gin_test/model"　//解決
```

- router.go

```go
router.goのUserListDisplayAction が宣言されていない。
```

```go
ListDisplayAction //関数名が異なっていた
```

- router.go
```go
EditCompleteAction が usercompiler パッケージによって宣言されていない。と出ました。とエラー

//edig.goに必要なコードの書き忘れ
```

# 学んだこと
```go
1. エラーを読んで、
2. 宣言されていない場合は関数名やimport名、go getなど幅広く疑う
ということが大切
```
```go
関数を使うときは、<パッケージ名>.<関数名>で呼び出す。

"宣言されていない"ということは、"正常に呼び出せていない"ということなので、
・定義している側
・使っている側
の
・パッケージ名
・関数名
を確認
```

```go
router.GETの中に書かれているディレクトリは、「URLのパス」です。

例えば、
router.GET("/aaa/bbb", top.IndexDisplayAction)
の場合は、 https://〇〇.com/aaa/bbb でアクセスが来た時に、` top.IndexDisplayAction`関数を呼び出す。という処理。

ディレクトリと必ずしも一致しない
```

# 今後の課題
- エラーは決めつけず幅広く疑うことが大事
- アプリ開発にもっと慣れてわからないがわからないに触れる
- コードをもっと読めるようになる
- 一つ一つどことどこが繋がってるか理解する
