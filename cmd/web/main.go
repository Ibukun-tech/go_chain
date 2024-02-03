package main

import (
	"net/http"

	"github.com/Ibukun-tech/go_chain/pkg/Handler"
)

func main() {
	http.HandleFunc("/new", Handler.NewBook)
	http.HandleFunc("/", Handler.GetBlockchain)
	http.HandleFunc("/", Handler.WriteBlock)
	http.ListenAndServe(":4000", nil)
}
