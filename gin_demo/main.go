package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method)
		fmt.Println(request.Host)
		fmt.Println(request.Header)
		_, _ = fmt.Fprintln(writer, "gin demo")
	})

	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}
