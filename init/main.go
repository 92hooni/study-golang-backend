package main

import (
	"fmt"
	"net/http"
)

func main() {
	// root("/")에 대한 handler function을 등록해준다.
	http.HandleFunc("/", helloWorld)

	// go의 http 서버에 대해 8080 port로 listener 를 등록해준다.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("에러가 발생했습니다.")
		panic(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}
