package main

import (
	"net/http"
	"strings"
	"os"
	"bufio"
)

const port = ":8000"

func main() {
	http.Handle("/", new(MyHandler))

	http.ListenAndServe(port, nil)
}

type MyHandler struct {
	http.Handler
}


func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request){
	path := "public/" +  req.URL.Path
	//data, err := ioutil.ReadFile(string(path))
	f, err := os.Open(path)

	if err == nil {
		bufferedReader := bufio.NewReader(f)
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/js"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "text/png"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content-Type", contentType)
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404) + " path " + path))
	}
}