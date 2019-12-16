package main

import (
	"fmt"
	"github.com/appoptics/appoptics-apm-go/v1/ao"
	"net/http"
	"strconv"
)

const Port = 8090

func requestHandler(w http.ResponseWriter, r *http.Request) {
	_ = ao.TraceFromContext(r.Context())
}

func main() {
	fmt.Printf("Start\n")
	http.HandleFunc("/", ao.HTTPHandler(requestHandler))
	fmt.Printf("%v", http.ListenAndServe(":"+strconv.Itoa(Port), nil))
}
