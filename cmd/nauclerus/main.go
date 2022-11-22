package main

import (
	"log"

	"github.com/juwit/nauclerus/http"
)

func main() {
	r := http.NewRouter()
	log.Fatal(r.Run(":8080"))
}
