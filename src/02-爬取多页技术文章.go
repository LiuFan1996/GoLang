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

// 用于提取链接内容的函数
func getContent(url string) (content string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get err:", err)
		return
	}

	defer resp.Body.Close()

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

func saveFile(no int, url string, title string, savePath string) {
	content := getContent(url)
	// 创建文件
	fileName := strconv.Itoa(no) + "_" + title + ".html"
	// 文件名中不允许包含 \/:*"<>?|
	replaceReg := regexp.MustCompile(`[\/:\*"<>\?\|]`)
	// 将文件名中不允许的字符替换成空，理想的办法是替换成类似的符号：例如："\"替换成"＼"
	fileName = replaceReg.ReplaceAllString(fileName, "")
	fmt.Println(fileName)

	file, err := os.Create(savePath + fileName)

	defer file.Close()

	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	file.Write([]byte(content))
}

func getOnePage(url string, page int, savePath string) {

	content := getContent(url + strconv.Itoa(page))
	//fmt.Println(content)

	// 创建文件夹savePath+dirName
	// savepath: c:/
	savePath = savePath + "/" + strconv.Itoa(page) + "/"
	err := os.MkdirAll(savePath, 666)
	if err != nil {
		fmt.Println("os.MkdirAll err:", err)
		return
	}
	//url:  "https://studygolang.com/articles?p=1"
	// 从url中提取服务器的根路径 https://studygolang.com
	regRoot := regexp.MustCompile(`https?://[^/]+`)
	// root是服务器的根路径: http://studygolang.com
	root := regRoot.FindString(url)

	// content是爬取页面的内容，需要从内容中提取文章链接，然后把链接中的内容再次爬取出来，然后保存成html文件
	// 提取的内容类似 “<h2><a href="/articles/15778" target="_blank" title="运维监控系统之Open-Falcon">
	reg := regexp.MustCompile(`<h2><a href="(/articles/\d+)" target="_blank" title="(.+)">`)
	arr := reg.FindAllStringSubmatch(content, -1)
	// 0_运维监控系统之Open-Falcon.html
	// 1_运维监控系统之Open-Falcon.html
	// 从数组中提取uri
	// v[0]=<h2><a href="(/articles/\d+)" target="_blank" title="(.+)">
	// v[1]=/articles/\d+
	// v[2]=标题
	for i, v := range arr {
		//fmt.Printf("i=%d uri=%v title=%v\n", i, root+v[1], v[2])
		// 将链接对应的每篇文章爬取下来，保存成本地html文件
		// i+1是文件的序号  root+v[1]是文件的url   v[2]文章标题
		saveFile(i+1, root+v[1], v[2], savePath)
	}
}

func main() {
	var (
		beginPage int    // 起始页
		endPage   int    // 结束页
		savePath  string // 保存路径
	)

	for {
		fmt.Print("请输入起始页：")
		fmt.Scan(&beginPage)
		fmt.Print("请输入结束页：")
		fmt.Scan(&endPage)
		fmt.Print("请输入保存路径：")
		fmt.Scan(&savePath)

		url := "https://studygolang.com/articles?p="
		for i := beginPage; i <= endPage; i++ {
			getOnePage(url, i, savePath)
		}

	}
	reg := `<link\s+.*\s+href="(.+\.css).*"[^>]>`
	fmt.Print(reg)
}
