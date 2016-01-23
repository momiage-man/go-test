package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

func main() {
	// root
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// web サーバを開始します
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// templは１つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServerHttpはHTTPリクエストを処理します
func (t *templateHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil) // 戻り値チェック
}
