package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)
	builder := gtk.NewBuilder()
	builder.AddFromFile("demo.glade")
	windows1 := gtk.WindowFromObject(builder.GetObject("window1"))
	button1 := gtk.ButtonFromObject(builder.GetObject("button1"))
	button1.Clicked(func() {
		fmt.Println("已经注册")
	})
	windows1.Show()
	gtk.Main()

}
