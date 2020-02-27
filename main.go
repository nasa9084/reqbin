package main

import (
	"bufio"
	"bytes"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	if err := execute(); err != nil {
		log.Printf("error: %+v", err)
		os.Exit(1)
	}
}

func execute() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	http.HandleFunc("/webhook", webhookHandler)
	return http.ListenAndServe(":8080", nil)
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sc := bufio.NewScanner(bytes.NewBuffer(b))
	log.Print("==========")
	for sc.Scan() {
		log.Print(sc.Text())
	}
	log.Print("==========")
	if err := sc.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
