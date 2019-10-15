# go-chat
ブラウザ上でチャットを行える

## 環境設定(Mac)

go のインストール
```
brew install go
```

Path の設定
```
mkdir $HOME/go
echo 'export GOPATH=$(go env GOPATH)' >> ~/.bash_profile
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bash_profile
source ~/.bash_profile
```

