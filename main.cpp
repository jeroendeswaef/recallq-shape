#include <gtkmm.h>

#include "mainwindow.h"

int main(int argc, char *argv[])
{
 
  Gtk::Main kit(argc, argv);
  MainWindow GtkmmTutorial;
  // GtkmmTutorial.set_default_size(400, 400);
  // GtkmmTutorial.set_title("My title");
  Gtk::Main::run(GtkmmTutorial);

  return 0;
}