package src

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", root)
	r.Get("/hello", hello)
	r.Get("/time", printTime)

	http.ListenAndServe(":3000", r)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is root page\n"))
	w.Write([]byte(r.URL.Path))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!\n"))
}

func printTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05 Mon"))
	w.Write([]byte("Current time: " + now.Format("2006-01-02 15:04:05 Mon") + "\n"))
}
