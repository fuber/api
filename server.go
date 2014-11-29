package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fuber/util" // HL
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
	bufrw.WriteString(util.ChunkedHeader())
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
	fmt.Printf("Running Fuber API server at :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
