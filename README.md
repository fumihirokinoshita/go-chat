# go-chat
ブラウザ上でチャットを行える

## 環境設定(Mac)

go のインストール
```
$ brew install go
```

Path の設定
```
$ mkdir $HOME/go
$ echo 'export GOPATH=$(go env GOPATH)' >> ~/.bash_profile
$ echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bash_profile
$ source ~/.bash_profile
```

clone
```
$ cd $HOME/go
$ git clone github.com/fumihirokinoshita/go-chat
```

ライブラリのインストール
```
$ go get github.com/gorilla/websocket
$ go get github.com/stretchr/gomniauth/...
```

## 実行
```
$ cd chat
$ go build -o chat
$ ./chat -host=:8080
```

ブラウザから
```
http://localhost:8080
```
にアクセス
複数タブからログインするとチャット開始