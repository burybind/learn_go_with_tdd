package main

import (
	"log"
	"net/http"
)

func main() {
	svr := PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", svr))
}
