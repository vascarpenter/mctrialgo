# mctrialgo

# 多施設共同研究用サーバー

node.jsで作っていましたが、
- コンパイラではない
- 型指定がいまいち
- ejsが助長
- ORMでない

などのいくつかの理由からやはり記述言語を変えるべきと思われました

候補にあがったのが swiftか goでしたが、ここはgoが親和性がよさそうだったので試してみました。

- webフレームワークには echo
- ORMには sqlboiler
- (Beegoはデータベースを読めなかったので外しました)

go開発環境のインストール
```
brew install go
```

gomodを使ったので
```
go build server.go
```
で必要なモジュールは自動的にダウンロードされるはず..

Sequel Proなどでデータベースを設計

sqlboiler.tomlファイルを作り
```
[mysql]
  dbname  = "studydb"
  host    = "localhost"
  port    = 3306
  user    = "oge"
  pass    = "hogehogeA00"
  sslmode = “false"
```

(mysqlのパスワードが@を含んでいるとsql URLに"@"が含まれていて問題だったのでやむなく変更しました
)

` sqlboiler mysql　`とすると
modelsディレクトリに自動でデータベースの内容を読み込んで go ファイルを作ってくれる

ディレクトリ構造
```
├── models
├── routes
├── static
│   └── css
└── views
```

データの準備
```
mysqlにoge, hogehogeA00というユーザを準備する
make import  (SQLが読み込まれます)
```

実行
```
go run server.go
あるいは
make run
```

更新
- login処理を作成
- go.mod環境に移行
- DB 構造に INTとUINTが混ざっていたので統一
- NULL可能なDBカラムは、sqlboilerでは、null.Stringとなり、これは .Valueをつけてhtmlの方から参照したり、null.StringFrom(string)等に変更しないと使えない
- VSCodeのlintはかなーり強力だが、html templateの中の typo までは指摘してくれず真っ白画面になる   html comment  {{/* */}} なんぞを使って切り分けしていくしかない

注意
- cookieのsecret keyがハードコードされてたりまだ未実装です。
- 更新中です

