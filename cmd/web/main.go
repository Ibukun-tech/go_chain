package main

import (
	"net/http"

	"github.com/Ibukun-tech/go_chain/pkg/Handler"
)

func main() {
	http.HandleFunc("/", Handler.NewBook)
}
