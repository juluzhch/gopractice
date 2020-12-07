package main

import (
	"github.com/juluzhch/gopractice/httpweb"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpweb.Echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	//flag.NewFlagSet()

}
