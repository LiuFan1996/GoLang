package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var quit chan int
var save chan string
var begenPage int
var endPage int
var savePath string

//获取当前页的所有文章页面内容
func getOneAllPage(uri string, page int, savepath string) {
	savepath = savepath + "/" + strconv.Itoa(page) + "/"
	err := os.MkdirAll(savepath, 666)
	if err != nil {
		fmt.Println(" os.MkdirAll err:", err)
		return
	}
	uri = uri + strconv.Itoa(page)

	//获取本页的所有对应链接
	content := getContent(uri)
	//确定根路径
	root := "https://studygolang.com"
	//对获取的页面进行解析
	reg := regexp.MustCompile(`<h2><a href="(/articles/\d+)" target="_blank" title="(.+)">`)
	arr := reg.FindAllStringSubmatch(content, -1)

	for i, v := range arr {
		go saveFile(root+v[1], i+1, v[2], savepath)
	}
	for i := begenPage; i < endPage; i++ {
		quit <- i
	}

}

//获取本页内容
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
	content = buffer.String()
	return
}

//保存当前uri 的内容
func saveFile(uri string, num int, filename string, savepath string) {
	content := getContent(uri)
	filename = strconv.Itoa(num) + "_" + filename + ".html"
	compile := regexp.MustCompile(`[\/:\*"<>\?\|]`)
	filename = compile.ReplaceAllString(filename, "")

	file, err := os.Create(savepath + filename)
	defer file.Close()
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	file.Write([]byte(content))
}
func main() {
	uri := "https://studygolang.com/articles?p="
	for {

		fmt.Println("请输入起始页面")
		fmt.Scan(&begenPage)
		fmt.Println("请输入结束页面")
		fmt.Scan(&endPage)
		fmt.Println("请输入保存路径如C:/a")
		fmt.Scan(&savePath)

		for i := begenPage; i <= endPage; i++ {
			go getOneAllPage(uri, i, savePath)
		}
		fmt.Println("操作成功:", <-quit)
		for {
			s := <-quit
			fmt.Println("quit=", s)
			if 20 == s {
				fmt.Println("操作成功，是否继续(Y/N)")
				var choose string
				fmt.Scan(&choose)
				if "Y" == choose || "y" == choose {
					continue
				} else {
					break
				}
			}
		}

	}
}
