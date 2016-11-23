package main

import (
	"net/http"
)

const port = ":8000"

func main() {
	http.ListenAndServe(port, http.FileServer(http.Dir("public")))
}

