package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>チャット</title>
				</head>
				<body>
					チャットしよう！
				</body>
			</html>
		`))
	})
	// start web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndserve:", err)
	}
}
