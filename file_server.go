package main

import (
    "net/http"
    "fmt"
)

func main(){
    h := http.FileServer(http.Dir("."))
    var port string
    fmt.Printf("Please input port Number: ")
    fmt.Scanf("%s",&port)
    http.ListenAndServe(":"+port, h)
}