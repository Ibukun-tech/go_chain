package main

import (
	"fmt"
	"net/http"

	"github.com/Ibukun-tech/go_chain/pkg/Handler"
)

func main() {
	http.HandleFunc("/new", Handler.NewBook)
	// http.HandleFunc("/", Handler.GetBlockchain)
	http.HandleFunc("/", Handler.WriteBlock)
	fmt.Println("Its running on Port 4000")
	http.ListenAndServe(":4000", nil)
}
