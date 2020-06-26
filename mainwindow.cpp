#include "mainwindow.h"

MainWindow::MainWindow() : m_Vbox(Gtk::ORIENTATION_VERTICAL), myImage("0YzJMi-8_400x400.jpg")
{
    set_title("Display Image");
    set_size_request(400, 400);

    add(m_Vbox);  
   
    m_Vbox.pack_start(myImage, Gtk::PACK_SHRINK);  
   
    show_all_children();
}

MainWindow::~MainWindow()
{
    //dtor
}
