package main

import (
	"fmt"
	"strings"
)

func main() {

	//判断是否以某字符串打头/结尾
	//
	//strings.HasPrefix(s, prefix string) bool
	//strings.HasSuffix(s string, suffix string) bool
	str1 := "书山有路,苦海无涯，天道酬勤，天不负有心人"
	fmt.Println(strings.HasPrefix(str1, "书"))
	fmt.Println(strings.HasSuffix(str1, "书"))
	//
	//字符串分割
	//
	//strings.Split(s string, sep string) []string
	//在go语言中通过split切割能得到一个完全切割数组
	fmt.Println(strings.Split(str1, ","))
	fmt.Println(strings.SplitAfter(str1, ","))
	fmt.Println(strings.SplitAfterN(str1, ",", 10))
	str2 := strings.SplitN(str1, ", , ,", 3)
	//for _, s1:=range str2  {
	//	fmt.Println(s1)
	//}
	//
	//返回子串索引
	//
	//strings.Index(s string, sub string) int
	//strings.LastIndex 最后一个匹配索引
	fmt.Println(strings.Index(str1, "山"))
	fmt.Println(len(str1))
	/*在go语言中如果string字符串类型中含有中文，默认一个中文占用两个字节，下标也是一个中文占用两个下标*/
	//
	//字符串连接
	//
	//strings.Join(a []string, sep string) string
	//另外可以直接使用“+”来连接两个字符串
	//str3:="lalala"
	//str4:="xixixi"
	fmt.Println(strings.Join(str2, "+"))

	//
	//字符串替换
	//
	//strings.Replace(s, old, new string, n int) string
	//
	str5 := "aAbAcAdAeAf"
	fmt.Println(strings.Replace(str5, "A", "", strings.Count(str5, "A")))
	/*将字符串替换和字符串计数配合使用：字符串替换第一个参数：目标字符串；第二个参数;目标字符串的子串；
	第三个参数：替换后的子串；第四个参数：替换目标的替换次数，可以使用Count计数方法获得子串在目标字符串
	中出现的次数，从而实现全部替换的功能,也小于0如-1实现全部替换*/
	//字符串转化为大小写
	//
	//strings.ToUpper(s string) string
	//strings.ToLower(s string) string
	//

	fmt.Println(strings.ToUpper(str5))
	fmt.Println(strings.ToLower(str5))
	//统计某个字符在字符串出现的次数
	//
	//strings.Count(s string, sep string) int
	//
	fmt.Println(strings.Count(str5, "A"))
	//判断字符串的包含关系

	//strings.Contains(s, substr string) bool
	fmt.Println(strings.Contains(str5, "A"))
	//与字符串相关的类型转换都是通过 strconv 包实现的。
	// strconv.Itoa(i int) string
}
