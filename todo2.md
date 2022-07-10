# todoアプリ　前半

#　やったこと
- トップ画面
- 会員一覧画面
- 会員追加画面
- 会員編集画面

# 詳細
- Controllerを作成
```go
コントローラーはクライアントから来たリクエストURLに応じて、あらかじめ設定した処理
【処理手順】
クライアントからリクエストをコントローラーが受け取ります。
そして受け取った値をモデル経由でデータベースからデータを取得します。
モデル経由で受け取ったデータをビューに渡します。
動的生成されたHTMLをクライアントに渡します。
```

- Viewを作成する
```go
Webブラウザに表示するHTMLを動的に生成する部分
動的にHTMLを生成するとは、データベースから取得したデータをHTMLに埋め込んでアクセスするたびに、異なる HTML を生成することなどがあります。このようなサイトを動的なサイトと呼んだりもします。
```

# エラー
```go
github.com/menber/go-user-management/controllerをインポートできませんでしたのエラー
```

個人のGitHubで詰まったエラー
//main.go
```go
"github.com/murai/go-user-management/server"
```
を
```go
"gin_test/server"
```
に書き換えたらエラーが消えた

//router.go
```go
could not import gin_test/controller (no required module provides packag
gin_test / controllerをインポートできませんでした（必要なモジュールはpackagを提供しません
```
↓
```go
"gin_test/controller/top" imported but not used
「gin_test/controller / top」はインポートされましたが、使用されていませんのエラーが出ます
```
↓
```go
router.GETの中で使っているtop
userに書き換えたらエラーが消えた
```

# 課題
- ディレクトリを作る場所の箇所が違っていたのでディレクトリの作る場所を気を付ける<br>
- 今回はrouter.godeエラーが出ていたのでルーティングについて理解を深める
- エラーが出たときの対処法についてもっと学習する
