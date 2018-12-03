package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/tianqi.json")

	defer resp.Body.Close()

	if err != nil {
		fmt.Println("http.Get err:", err)
		return
	}
	json := ""
	buf := make([]byte, 1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		json += string(buf[:n])
	}
	fmt.Println(json)
}
