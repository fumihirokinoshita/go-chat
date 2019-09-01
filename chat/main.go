package main

import (
	"flag"
	"go-chat/trace"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

// templは１つのテンプレートを表す
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServerHTTPはHTTPリクエストを処理する
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {

	// Gomniauthのセットアップ
	gomniauth.SetSecurityKey("TODO:セキュリティキー")
	gomniauth.WithProviders(
		facebook.New("クライアントID", "秘密鍵", "http://localhost:8080/auth/callback/facebook"),
		github.New("クライアントID", "秘密鍵", "http://localhost:8080/auth/callback/github"),
		google.New("462225854968-iprrepgp6ff08af4euqpr0htqcfihdkj.apps.googleusercontent.com", "lYoQBbPczT0x9QQqqK495tIG", "http://localhost:8080/auth/callback/google"),
	)

	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() // フラグを解釈する
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	// チャットルームを開始する
	go r.run()
	// Webサーバを起動
	log.Println("Webサーバを開始します。ポート：", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndserve:", err)
	}
}
