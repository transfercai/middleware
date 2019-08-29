package main

import (
	"net/http"
	"transfercai/middleware/middleware"
	"transfercai/middleware/service"
)

func main() {
	ch := make(chan int, 5)
	go middleware.LimitReq(10, ch)
	http.HandleFunc("/test", middleware.MakeHandleFunc(service.Test, ch))
	http.ListenAndServe(":8080", nil)
}
