package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var filenameURL string

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
	//dec := mahonia.NewDecoder("GB18030")
	//content = dec.ConvertString(content)
	return
}
func getAllPageURL(filenameURL string) (AllPageURL [][]string, filename [][]string) {
	content := getContent(filenameURL)
	compile := regexp.MustCompile(`<dd><a href="(.+?)">.*?</a></dd>`)
	AllPageURL = compile.FindAllStringSubmatch(content, -1)
	mustCompile := regexp.MustCompile(`<h1>(.+?)</h1>`)
	filename = mustCompile.FindAllStringSubmatch(content, -1)
	//fmt.Println(AllPageURL)
	return
}
func saveOneFile() {
	//获取目录页的所有章节的URL数组
	AllPageURL, filename := getAllPageURL(filenameURL)
	//fmt.Println(AllPageURL, "\nfilename=", filename)
	file, _ := os.Create("F:/" + filename[0][1] + ".txt")
	defer file.Close()
	result := filename[0][1] + "\n"
	fmt.Println(AllPageURL)
	for k, v := range AllPageURL {
		s := v[1]
		s = strings.Replace(s, `class="empty"`, "", -1)
		fmt.Println("爬取第", k, "章")
		cont := getContent("https://www.xbiquge6.com/" + s)
		if strings.Contains(v[0], `class="empty"`) {
			break
		}
		tiltle, conts := fileRegexp(cont)
		result += "\n" + tiltle + "\n" + conts + "\n"
		fmt.Println(conts)
	}
	result += "\n"
	file.Write([]byte(result))
	fmt.Println("爬取完成")
}

func fileRegexp(cont string) (tiltle string, conts string) {
	compile := regexp.MustCompile(`<div id="content">(?s:(.*?))</div>`)
	conts = compile.FindAllStringSubmatch(cont, -1)[0][1]
	compile = regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	tiltle = compile.FindAllStringSubmatch(cont, -1)[0][1]
	conts = strings.Replace(conts, "&nbsp;", " ", -1)
	conts = strings.Replace(conts, "<br>", "", -1)
	conts = strings.Replace(conts, "<br />", "", -1)
	return
}

func SaveAllFile() {
	//获取目录页的所有章节的URL数组
	AllPageURL, dirname := getAllPageURL(filenameURL)
	os.MkdirAll("F:/"+dirname[0][1], 666)
	for k, v := range AllPageURL {
		fmt.Println("爬取第", k, "章")
		cont := getContent("http://www.biquge.com.tw" + v[1])
		tiltle, conts := fileRegexp(cont)
		file, _ := os.Create("F:/" + dirname[0][1] + "/" + tiltle + ".txt")
		file.Write([]byte(conts))
		file.Close()
	}
	fmt.Println("爬取完成")
}
func main() {
	var flag int
	fmt.Println("欢迎使用笔趣阁小说爬虫（http://www.biquge.com.tw）")
	fmt.Print("请输入你要爬取的笔趣阁小说目录页网址:")
	fmt.Scan(&filenameURL)
	fmt.Println("------------------输入序号选择保存方式--------------------")
	fmt.Println("1：保存一个文件")
	fmt.Println("2：保存一多文件")
	fmt.Print("请输入保存方式:")
	fmt.Scan(&flag)
	if flag == 1 {
		saveOneFile()
	} else {
		SaveAllFile()
	}
	fmt.Print("开始爬取...")

}
