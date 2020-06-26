#include "drawimage.h"
#include <cairomm/context.h>
#include <giomm/resource.h>
#include <gdkmm/general.h>
#include <glibmm/fileutils.h>
#include <iostream>

DrawImage::DrawImage(const std::string &filename)
{
    try
    {
        myImage = Gdk::Pixbuf::create_from_file(filename);
    }
    catch (...)
    {
        std::cerr << "An error occured while loading the image file." << std::endl;
    }

    if (myImage)
        set_size_request(myImage->get_width(), myImage->get_height());
    set_halign(Gtk::ALIGN_CENTER);
    set_valign(Gtk::ALIGN_CENTER);
}

DrawImage::~DrawImage()
{
    //dtor
}

bool DrawImage::on_draw(const Cairo::RefPtr<Cairo::Context> &cr)
{
    if (!myImage)
        return false;

    Gtk::Allocation allocation = get_allocation();
    const int height = allocation.get_height();
    const int width = allocation.get_width();

    Gdk::Cairo::set_source_pixbuf(cr, myImage, width - myImage->get_width(), height - myImage->get_height());

    cr->paint();
    return true;
}
