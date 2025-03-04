package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/tranchida/echotest/internal/handler"
)

//go: embed static templates
//var contentFS embed.FS

func newServer() *http.ServeMux {

	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("GET /", handler.IndexHandler(tmpl))
	http.HandleFunc("GET /host", handler.HostInfoHandler(tmpl))
	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return http.DefaultServeMux
}

func main() {

	fmt.Println("open browser on : http://localhost:8080")

	if err := http.ListenAndServe(":8080", newServer()); err != nil {
		panic(err)
	}

}
