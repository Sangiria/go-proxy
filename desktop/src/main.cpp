#include "mainwindow.h"
#include <QApplication>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    MainWindow w;
    w.show();
    return a.exec();
}

// step 1. button add on clicked function.
// step 2. in function analyze url string and call AddNode or AddSub
