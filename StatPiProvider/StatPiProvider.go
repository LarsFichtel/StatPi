package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./frontends/html/templates")))
	http.ListenAndServe(":9090", nil)
}
