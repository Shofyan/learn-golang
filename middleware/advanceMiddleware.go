package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logging() Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}

}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func Method(s string) Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			if s != r.Method {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(w, r)
		}
	}

}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}