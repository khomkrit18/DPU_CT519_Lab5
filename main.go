package main

import (
    "fmt"
    "net/http"
)
func about(w http.ResponseWriter, req *http.Request){
        p := "./www/about.html"
        http.ServeFile(w, req, p)

}

func hello(w http.ResponseWriter, req *http.Request){
        p := "./www/index.html"
        http.ServeFile(w, req, p)

}

func myresearch(w http.ResponseWriter, req *http.Request){
        p := "./www/myresearch.html"
        http.ServeFile(w, req, p)

}


func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/", hello)
    http.HandleFunc("/about", about)
    http.HandleFunc("/myresearch", myresearch)
    http.HandleFunc("/headers", headers)
    http.ListenAndServe(":8080", nil)
}

