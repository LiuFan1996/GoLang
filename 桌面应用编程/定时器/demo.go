package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
)

func main() {
	// 初始化
	gtk.Init(&os.Args)
	// 创建窗口对象
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	// 设置窗口大小和标题
	win.SetSizeRequest(500, 300)
	win.SetTitle("定时器")
	// 创建表格
	table := gtk.NewTable(2, 2, true)

	// 创建label
	time := gtk.NewLabel("0")
	time.ModifyFontSize(50)

	// 开始和停止按钮
	startBtn := gtk.NewButtonWithLabel("start")
	stopBtn := gtk.NewButtonWithLabel("stop")
	stopBtn.SetSensitive(false) // 停止按钮默认不可点

	var id int // 用于保存定时器的标识
	// 点击开始按钮，开始计时，每秒递增1
	startBtn.Clicked(func() {
		//启动定时器, 1000毫秒为时间间隔，回调函数为匿名函数
		// id是定时器的标识，用于停止定时器的执行
		id = glib.TimeoutAdd(1000, func() bool {
			num, _ := strconv.Atoi(time.GetText())
			num++
			time.SetText(strconv.Itoa(num)) //给标签设置内容
			if num == 5 {
				return false // 定时器自动停止
			}
			return true //只要定时器没有停止，时间到自动调用回调函数
		})

		startBtn.SetSensitive(false) //启动按钮变灰，不能按
		stopBtn.SetSensitive(true)   //定时器启动后，停止按钮可以按
	})

	// 停止计时
	stopBtn.Clicked(func() {
		glib.TimeoutRemove(id)      // id是定时器的标识
		startBtn.SetSensitive(true) //启动按钮变灰，不能按
		stopBtn.SetSensitive(false) //定时器启动后，停止按钮可以按
	})

	// 将按钮添加到表格
	table.AttachDefaults(time, 0, 2, 0, 1)
	table.AttachDefaults(startBtn, 0, 1, 1, 2)
	table.AttachDefaults(stopBtn, 1, 2, 1, 2)

	// 将table添加到win中
	win.Add(table)

	// 当收到销毁信号时，执行gtk.MainQuit函数，指定函数名即可
	win.Connect("destroy", gtk.MainQuit)
	// 显示窗口及其中的控件
	win.ShowAll()
	//主事件循环，等待用户操作
	gtk.Main()
}
