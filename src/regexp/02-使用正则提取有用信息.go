package regexp

import (
	"fmt"
	"regexp"
)

func main() {
	// 提取数字
	/*str := "hello 12  jack 345 xxx 5  uuu"
	reg := regexp.MustCompile(`\d+`)
	slice := reg.FindAllString(str, -1)
	fmt.Println("clice=", slice)*/
	str := "hello +12 jack 3.45 xxx 5 uuu 3.14 kk -3  mm -1.2"
	//reg := regexp.MustCompile(`(-|\+)?\d+(\.\d+)?`)
	reg := regexp.MustCompile(`[+-]?\d+(\.\d+)?`)
	slice := reg.FindAllString(str, -1)
	fmt.Println("clice=", slice)
}
