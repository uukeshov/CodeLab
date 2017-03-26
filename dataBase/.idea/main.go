package main

import (
	"github.com/uukeshov/codelab/api"
	"log"
	"flag"
)

func main() {
	bindAddr := flag.String("bind_addr", ":8080", "Set bind address")
	flag.Parse()
	a := api.New(*bindAddr)
	log.Fatal(a.Start())
}