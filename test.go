package main

import(
    "fmt"
    "net/http"
)

func main(){
    http.HandleFunc("/",helloHandler)
    http.ListenAndServe(":8090",nil)
}

func helloHandler(w http.ResponseWriter,r *http.Request){
    fmt.Fprintln(w, "Hello, World!!3!")
}