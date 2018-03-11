package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/msbond/gorouter/cool"
)

func main(){
    fmt.Print("cool")
    r := mux.NewRouter()
    r.HandleFunc("/", cool.Handler)
    http.ListenAndServe(":8080", r)
}



