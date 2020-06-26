package main

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	pixbuf, err := gdk.PixbufNewFromFile("0YzJMi-8_400x400.jpg")
	if err != nil {
		log.Fatal("Unable to create pixbuf", err)
	}

	img, err := gtk.ImageNewFromPixbuf(pixbuf)
	if err != nil {
		log.Fatal("Unable to create image", err)
	}

	// l, err := gtk.LabelNew("Hello, gotk3!")
	// if err != nil {
	// 	log.Fatal("Unable to create label:", err)
	// }
	// button, err := gtk.ButtonNewWithLabel("Click me")
	// if err != nil {
	// 	log.Fatal("Unable to create button", err)
	// }

	menuBar, err := gtk.MenuBarNew()
	if err != nil {
		log.Fatal("Unable to create menu bar", err)
	}

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("Unable to create box", err)
	}

	// button.Connect("clicked", func() {
	// 	log.Print("Clicked")
	// })

	box.Add(menuBar)
	box.Add(img)

	win.Add(box)

	// Set the default window size.
	win.SetDefaultSize(400, 400)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
