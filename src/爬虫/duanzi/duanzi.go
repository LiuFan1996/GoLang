package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//https://www.pengfu.com/xiaohua_1.html
	uri := "https://www.pengfu.com/xiaohua_"
	result := ""
	for i := 1; i <= 50; i++ {
		content := getContent(uri + strconv.Itoa(i) + ".html")
		r := saveFile(content, uri+strconv.Itoa(i))
		result += r
	}
	file, _ := os.Create("段子.txt")
	defer file.Close()
	file.Write([]byte(result))
	fmt.Println("爬取完成")
}
func getContent(uri string) (content string) {
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("http.Get err:", err)
		return
	}
	buf := make([]byte, 1024)
	buffer := bytes.Buffer{}
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("resp.Body.Read err:", err)
			return
		}
		if n == 0 {
			break
		}
		buffer.Write(buf[:n])
	}
	fmt.Println("爬取：", uri, "成功")
	content = buffer.String()
	return
}
func saveFile(content string, uri string) (result string) {

	reg := regexp.MustCompile(`<div class="content-img clearfix pt10 relative">(?s:(.*?))</div>`)
	arr1 := reg.FindAllStringSubmatch(content, -1)
	reg = regexp.MustCompile(`<h1 class="dp-b"><a href="https://.+" target="_blank">(?s:(.*?))</a>`)
	arr2 := reg.FindAllStringSubmatch(content, -1)
	for k, a := range arr1 {
		result += "\n                 标题：" + arr2[k][1] + "\n"
		con := a[1]
		con = strings.Replace(con, "<br/>", "\n", -1)
		con = strings.Replace(con, "<br>", "", -1)
		con = strings.Replace(con, "<br />", "", -1)
		con = strings.Replace(con, "										", "", -1)
		con = strings.Replace(con, "&nbsp;", " ", -1)
		con = strings.Replace(con, "&nbsp", " ", -1)
		result += con + "\n"
	}
	result += "\n"
	fmt.Println("保存：", uri, "成功")
	return
}
