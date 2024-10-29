package main

import (
    "log"
    "net/http"
    "path/filepath"
    "strings"
)

func html(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/style.css" {
        http.ServeFile(w, r, "/html/style.css")
        return
    }
    if strings.HasPrefix(r.URL.Path, "/static/") {
        http.ServeFile(w, r, r.URL.Path)
        return
    }
    if r.URL.Path == "/" {
        http.ServeFile(w, r, "/html/index.html")
        return
    }
    http.ServeFile(w, r, filepath.Join("/html", r.URL.Path + ".html"))
}

func main() {
    http.HandleFunc("/", html)

    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal(err)
    }
}
