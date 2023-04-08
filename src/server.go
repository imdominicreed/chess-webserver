package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = 8080

var handlerMap = map[string]func(http.ResponseWriter, *http.Request){"/hello": handleHello}

func main() {
	for handle, function := range handlerMap {
		http.HandleFunc(handle, function)
	}

	fmt.Printf("Starting server at port %d\n", PORT)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Fatal(err)
		return
	}
}
