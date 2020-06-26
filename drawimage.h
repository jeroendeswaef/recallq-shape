#ifndef DRAWIMAGE_H  
#define DRAWIMAGE_H  
   
#include <gtkmm/drawingarea.h>  
#include <gdkmm/pixbuf.h>  
   
   
 class DrawImage : public Gtk::DrawingArea  
 {  
   public:  
     DrawImage(const std::string& filename);  
     virtual ~DrawImage();  
   protected:  
     Glib::RefPtr<Gdk::Pixbuf> myImage;  
     bool on_draw(const Cairo::RefPtr<Cairo::Context>& cr) override;  
 };  
   
 #endif // DRAWIMAGE_H  