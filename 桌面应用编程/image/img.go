package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

var imageOneName string = "1.JPG"
var imageTwoName string = "2.JPG"
var flag bool = true

func main() {
	gtk.Init(&os.Args)
	builder := gtk.NewBuilder()
	builder.AddFromFile("img.glade")
	windows1 := gtk.WindowFromObject(builder.GetObject("window1"))
	img := gtk.ImageFromObject(builder.GetObject("image1"))
	button1 := gtk.ButtonFromObject(builder.GetObject("button1"))

	var w, h int
	img.GetSizeRequest(&w, &h)
	fmt.Println(w, h)
	pixbuf1, _ := gdkpixbuf.NewPixbufFromFileAtScale(imageOneName, w, h, false)
	img.SetFromPixbuf(pixbuf1)
	button1.Clicked(func() {
		if flag {
			pixbuf2, _ := gdkpixbuf.NewPixbufFromFileAtScale(imageTwoName, w, h, false)
			img.SetFromPixbuf(pixbuf2)
			flag = false
		} else {
			img.SetFromPixbuf(pixbuf1)
			flag = true
		}
		fmt.Println("已经切换")
	})
	windows1.Show()
	gtk.Main()

}
