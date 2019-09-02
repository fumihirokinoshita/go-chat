package main

import (
	"flag"
	"go-chat/trace"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/stretchr/objx"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

// 現在アクティブなAvatarの実装
var avatars Avatar = TryAvatars{
	UseFileSystemAvatar,
	UseAuthAvatar,
	UseGravatar}

// templは１つのテンプレートを表す
type templateHandler struct {
	filename string
	templ    *template.Template
}

// ServeHTTPはHTTPリクエストを処理する
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if t.templ == nil {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	}
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, data)
}

var host = flag.String("host", ":8080", "アプリケーションのアドレス")

func main() {

	flag.Parse() // フラグを解釈する

	// Gomniauthのセットアップ
	gomniauth.SetSecurityKey("xv1vjk3xgoz8ic90zgnryxdz")
	gomniauth.WithProviders(
		facebook.New("2369387316463639", "6a942a6882208ffada387542e4d1033d", "http://localhost:8080/auth/callback/facebook"),
		github.New("2bfb18da1ac0a5bc59bb", "a0a02644be5916ca03635dd05469778ad1658b97", "http://localhost:8080/auth/callback/github"),
		google.New("462225854968-iprrepgp6ff08af4euqpr0htqcfihdkj.apps.googleusercontent.com", "lYoQBbPczT0x9QQqqK495tIG", "http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/uploader", uploaderHandler)

	http.Handle("/avatars/", http.StripPrefix("/avatars/", http.FileServer(http.Dir("./avatars"))))

	// チャットルームを開始する
	go r.run()
	// Webサーバを起動
	log.Println("Webサーバを開始します。ポート：", *host)
	if err := http.ListenAndServe(*host, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
