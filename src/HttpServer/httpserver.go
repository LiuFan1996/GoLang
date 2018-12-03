package main

import (
	"net/http"
	"os"
)

func allpattern(res http.ResponseWriter, req *http.Request) {
	url := req.RequestURI
	filepath := url[1:]
	file, _ := os.Open(filepath)
	buffer := make([]byte, 1024)
	defer file.Close()
	for {
		n, _ := file.Read(buffer)
		if n == 0 {
			break
		}
		res.Write(buffer[:n])
	}
}
func main() {
	http.HandleFunc("/", allpattern)
	http.ListenAndServe("127.0.0.1:8000", nil)

}
