/**
this demo based on web.go.

install:
go get github.com/hoisie/web

run:
go run hello_web.go

view in your browser:
http://localhost:9999/world
*/


package main

import (
    "github.com/hoisie/web"
)

func hello(val string) string { return "hello " + val } 

func main() {
    web.Get("/(.*)", hello)
    web.Run("0.0.0.0:9999")
}