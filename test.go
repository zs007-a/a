package main

import(
    "fmt"
    "net/http"
)

func main(){
    http.HandleFunc("/",helloHandler)
    http.ListenAndServe(":8090",nil)
}

func helloHandler(w http.ResponseWrite,r *http.Resquest){
    fmt.Fprintln("hello world")
}