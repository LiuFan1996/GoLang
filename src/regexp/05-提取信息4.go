package regexp

import (
	"fmt"
	"regexp"
)

func main() {
	// 提取div标签中的内容
	html := `11
	<div>11
1</div>
	<DIV>222</div>
	<div>333</div>
222
`
	// 使用()将需要捕获的内容扩起来
	// i表示忽略大小写   s表示使.能够匹配换号符\n
	reg := regexp.MustCompile(`(?is:<div>(.*?)</div>)`)
	// 返回结果是一个二维数组[[<div>111</div> 111] [<div>222</div> 222] [<div>333</div> 333]]
	result := reg.FindAllStringSubmatch(html, -1)
	fmt.Println("result = ", result)

	// 将提取的结果保存到切片中
	result2 := make([]string, len(result))
	for i, value := range result {
		result2[i] = value[1]
	}
	//fmt.Println("result2 = ", result2)
}
