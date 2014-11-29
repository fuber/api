package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RootHandler(w http.ResponseWriter, req *http.Request) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	conn, bufrw, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	bufrw.WriteString(`HTTP/1.1 200 OK
Date: Tue, 26 Nov 2014 22:28:41 GMT
Content-Type: text/plain; charset=UTF-8
Transfer-Encoding: chunked

`)
	for _ = range time.NewTicker(500 * time.Millisecond).C {
		served += 1
		bufrw.WriteString("5\r\n")
		bufrw.WriteString("💩 \r\n")
		bufrw.Flush()
	}
}

func CounterHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(fmt.Sprintf("%d\n", served)))
}

var served int = 0

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/counter", CounterHandler)
	port := 12345
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
