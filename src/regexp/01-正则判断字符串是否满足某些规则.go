package regexp

import (
	"fmt"
	"regexp"
)

func main() {
	/*str := `12345678901`
	reg := regexp.MustCompile(`\d{11}`)
	result := reg.MatchString(str)
	fmt.Println(result)*/

	/*str := `bbbb 12345678901 aaaa`
	reg := regexp.MustCompile(`^\d{11}$`)
	result := reg.MatchString(str)
	fmt.Println(result)*/

	/*str := `15852086521`
	reg := regexp.MustCompile(`^1[35789]\d{9}$`)
	result := reg.MatchString(str)
	fmt.Println(result)*/

	str := `15852086521`
	reg := regexp.MustCompile(`^(\+86)?1[35789]\d{9}$`)
	result := reg.MatchString(str)
	fmt.Println(result)

}
