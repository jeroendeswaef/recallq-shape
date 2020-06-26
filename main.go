package main

import (
	"errors"
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appID = "com.github.gotk3.gotk3-examples.glade"

func main() {
	// Create a new application.
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)

	data, err := Asset("ui/ui.glade")
	errorCheck(err)

	// Connect function to application activate event
	application.Connect("activate", func() {
		log.Println("application activate")

		// Get the GtkBuilder UI definition in the glade file.
		builder, err := gtk.BuilderNew()
		builder.AddFromString(string(data))
		// builder, err := gtk.BuilderNewFromFile("ui/ui.glade")
		errorCheck(err)

		// Map the handlers to callback functions, and connect the signals
		// to the Builder.
		signals := map[string]interface{}{
			"on_main_window_destroy": onMainWindowDestroy,
			"on_button_clicked": func() {
				log.Println("Clicked!")
				backgroundStackObj, err := builder.GetObject("background_stack")
				errorCheck(err)

				backgroundStack := backgroundStackObj.(*gtk.Stack)
				backgroundStack.SetVisibleChildName("page1")
			},
		}
		builder.ConnectSignals(signals)

		// Get the object with the id of "main_window".
		obj, err := builder.GetObject("main_window")
		errorCheck(err)

		// Verify that the object is a pointer to a gtk.ApplicationWindow.
		win, err := isWindow(obj)
		errorCheck(err)

		// Show the Window and all of its components.
		win.Show()
		application.AddWindow(win)
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})

	// Launch the application
	os.Exit(application.Run(os.Args))
}

func isWindow(obj glib.IObject) (*gtk.Window, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}

// onMainWindowDestory is the callback that is linked to the
// on_main_window_destroy handler. It is not required to map this,
// and is here to simply demo how to hook-up custom callbacks.
func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
	os.Exit(0)
}
