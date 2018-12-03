package main

import "card"

func main() {
	var c card.CcardStr
	var pc1 card.PlayerCardStr
	var pc2 card.PlayerCardStr
	var pc3 card.PlayerCardStr
	pc1.Pname = "张三"
	pc2.Pname = "李四"
	pc3.Pname = "王五"
	//生成扑克牌
	c.ShengCheng()
	//洗牌
	c.XiPai()
	var j int = 0
	//发牌
	for i := 0; i < 17; i++ {
		pc1.Pcard[i] = c.Ccard[j]
		j++
		pc2.Pcard[i] = c.Ccard[j]
		j++
		pc3.Pcard[i] = c.Ccard[j]
		j++
	}

	pc1.ToString()
	pc2.ToString()
	pc3.ToString()
}
