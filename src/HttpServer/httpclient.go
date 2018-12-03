package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://127.0.0.1:8000/q.json")
	buffer := make([]byte, 1024)
	defer resp.Body.Close()
	json := ""
	for {
		n, _ := resp.Body.Read(buffer)
		if n == 0 {
			break
		}
		json += string(buffer[:n])
	}
	fmt.Println(json)

}
