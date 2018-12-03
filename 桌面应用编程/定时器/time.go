package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
)

func main() {
	gtk.Init(&os.Args)
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	table := gtk.NewTable(2, 2, true)
	btn1 := gtk.NewButtonWithLabel("start")
	btn2 := gtk.NewButtonWithLabel("stop")
	lable := gtk.NewLabel("0")
	win.SetSizeRequest(500, 300)
	win.SetTitle("定时器")
	lable.ModifyFontSize(50)
	btn2.SetSensitive(false)
	var id int
	btn1.Clicked(func() {
		id = glib.TimeoutAdd(1000, func() bool {
			i, _ := strconv.Atoi(lable.GetText())
			i++
			lable.SetText(strconv.Itoa(i))
			return true
		})
		btn2.SetSensitive(true)
		btn1.SetSensitive(false)
	})

	btn2.Clicked(func() {
		glib.TimeoutRemove(id)
		btn2.SetSensitive(false)
		btn1.SetSensitive(true)
	})
	table.AttachDefaults(lable, 0, 2, 0, 1)
	table.AttachDefaults(btn1, 0, 1, 1, 2)
	table.AttachDefaults(btn2, 1, 2, 1, 2)
	win.Add(table)
	win.Connect("destory", gtk.MainQuit)
	win.ShowAll()
	gtk.Main()

}
