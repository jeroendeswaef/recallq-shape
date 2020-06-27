package main

import (
	"errors"
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appID = "github.com.jeroendeswaef.recallq-shape"

var mainWindow *gtk.Window
var builder *gtk.Builder

func main() {
	// Create a new application.
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)
	mainWindow = application.GetActiveWindow()

	data, err := Asset("ui/ui.glade")
	errorCheck(err)

	// Connect function to application activate event
	application.Connect("activate", func() {
		// Get the GtkBuilder UI definition in the glade file.
		builder, err = gtk.BuilderNew()
		builder.AddFromString(string(data))
		errorCheck(err)

		// Map the handlers to callback functions, and connect the signals
		// to the Builder.
		signals := map[string]interface{}{
			"on_main_window_destroy":     onMainWindowDestroy,
			"on_button_clicked":          onChooseBackgroundImage,
			"on_choose_background_image": onChooseBackgroundImage,
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

func onMainWindowDestroy() {
	os.Exit(0)
}

func onChooseBackgroundImage() {
	fileChooserDialog, err := gtk.FileChooserNativeDialogNew("Choose background image", mainWindow, gtk.FILE_CHOOSER_ACTION_OPEN, "_Open", "_Cancel")
	errorCheck(err)

	imageFileFilter, err := gtk.FileFilterNew()
	errorCheck(err)

	imageFileFilter.AddMimeType("image/*")
	fileChooserDialog.SetFilter(imageFileFilter)

	res := fileChooserDialog.Run()
	filename := fileChooserDialog.GetFilename()

	if gtk.ResponseType(res) == gtk.RESPONSE_ACCEPT {
		log.Printf("Accept: %s\n", filename)
		backgroundImageObj, err := builder.GetObject("background_image")
		errorCheck(err)
		backgroundImage := backgroundImageObj.(*gtk.Image)
		backgroundImage.SetFromFile(filename)
		backgroundStackObj, err := builder.GetObject("background_stack")
		errorCheck(err)

		backgroundStack := backgroundStackObj.(*gtk.Stack)
		backgroundStack.SetVisibleChildName("page1")
	}

}
