#ifndef MAINWINDOW_H  
#define MAINWINDOW_H  
   
#include <gtkmm.h>  
#include "drawimage.h"  
   
   
 class MainWindow : public Gtk::Window  
 {  
   public:  
     MainWindow();  
     virtual ~MainWindow();  
   protected:  
     Gtk::Box m_Vbox;
     DrawImage myImage;   
 };  
   
 #endif // MAINWINDOW_H  
