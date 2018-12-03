package main

import "fmt"

func main() {
		//关于Map的操作

		//初始化
		map1:=make(map[string]string ,5)
		map2:=make(map[string]string)
		map3:=map[string]string{}
		map4:=map[string]string{"a":"b","b":"c"}
		fmt.Println(map1,map2,map3,map4)

		//注意空值map和空map的区别也就是有没有初始化的区别
		var null_map map[string]string
		empty_map :=map[string]string{}
		fmt.Println(null_map==nil) //true 没有初始化的map其等于nil
		fmt.Println(empty_map==nil)//false 初始化但没有内容的map其不等于没有初始化的map


		//map 的遍历

		var city_map =map[string]string{}
		city_map["安徽"]="合肥，芜湖，铜陵"
		city_map["江苏"]="南京，无锡，苏州"

	for key,value:=range city_map  {
		fmt.Println(key,":" ,value)
	}


	//map的增删改查

	//增
	city_map["山东"]="济南，青岛，烟台"
	//删除
	delete(city_map, "山东")
	//改
	city_map["山东"]="济南，青岛，营口"
	//查
	if _, ok := city_map["山东"]; ok {
		fmt.Println(city_map["山东"])
	}

	//Go 语言中map和slice，func一样，不支持 == 操作符，
	// 就是不能直接比较。唯一合法的就是和nil作比较，判断该map是不是零值状态。

}
