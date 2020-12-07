package pprof

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func DoInHttp() {
	work()
	log.Fatal(http.ListenAndServe("0.0.0.0:6060", nil))
}
func work() {
	go func() {
		data := []byte("abc")
		var data2 string = string(data)
		fmt.Println(data2)
	}()
}
