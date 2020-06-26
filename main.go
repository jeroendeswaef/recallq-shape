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

	win.Add(img)

	// Set the default window size.
	win.SetDefaultSize(400, 400)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
