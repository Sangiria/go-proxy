/********************************************************************************
** Form generated from reading UI file 'mainwindow.ui'
**
** Created by: Qt User Interface Compiler version 6.9.2
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_MAINWINDOW_H
#define UI_MAINWINDOW_H

#include <QtCore/QVariant>
#include <QtWidgets/QApplication>
#include <QtWidgets/QHBoxLayout>
#include <QtWidgets/QHeaderView>
#include <QtWidgets/QMainWindow>
#include <QtWidgets/QPushButton>
#include <QtWidgets/QRadioButton>
#include <QtWidgets/QSpacerItem>
#include <QtWidgets/QStatusBar>
#include <QtWidgets/QToolButton>
#include <QtWidgets/QTreeWidget>
#include <QtWidgets/QVBoxLayout>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_MainWindow
{
public:
    QWidget *centralwidget;
    QVBoxLayout *verticalLayout;
    QWidget *widget;
    QHBoxLayout *horizontalLayout;
    QPushButton *btnAddSubscription;
    QPushButton *btnUpdate;
    QSpacerItem *horizontalSpacer;
    QTreeWidget *treeWidget;
    QWidget *widget_2;
    QHBoxLayout *horizontalLayout_2;
    QToolButton *btnStart;
    QRadioButton *rdbtnTun;
    QRadioButton *rdbtnSysProxy;
    QSpacerItem *horizontalSpacer_2;
    QStatusBar *statusbar;

    void setupUi(QMainWindow *MainWindow)
    {
        if (MainWindow->objectName().isEmpty())
            MainWindow->setObjectName("MainWindow");
        MainWindow->resize(800, 599);
        MainWindow->setMinimumSize(QSize(716, 0));
        MainWindow->setMaximumSize(QSize(800, 600));
        centralwidget = new QWidget(MainWindow);
        centralwidget->setObjectName("centralwidget");
        verticalLayout = new QVBoxLayout(centralwidget);
        verticalLayout->setObjectName("verticalLayout");
        widget = new QWidget(centralwidget);
        widget->setObjectName("widget");
        widget->setStyleSheet(QString::fromUtf8(""));
        horizontalLayout = new QHBoxLayout(widget);
        horizontalLayout->setObjectName("horizontalLayout");
        btnAddSubscription = new QPushButton(widget);
        btnAddSubscription->setObjectName("btnAddSubscription");
        btnAddSubscription->setStyleSheet(QString::fromUtf8(""));

        horizontalLayout->addWidget(btnAddSubscription);

        btnUpdate = new QPushButton(widget);
        btnUpdate->setObjectName("btnUpdate");

        horizontalLayout->addWidget(btnUpdate);

        horizontalSpacer = new QSpacerItem(40, 20, QSizePolicy::Policy::Expanding, QSizePolicy::Policy::Minimum);

        horizontalLayout->addItem(horizontalSpacer);


        verticalLayout->addWidget(widget);

        treeWidget = new QTreeWidget(centralwidget);
        treeWidget->setObjectName("treeWidget");

        verticalLayout->addWidget(treeWidget);

        widget_2 = new QWidget(centralwidget);
        widget_2->setObjectName("widget_2");
        horizontalLayout_2 = new QHBoxLayout(widget_2);
        horizontalLayout_2->setObjectName("horizontalLayout_2");
        btnStart = new QToolButton(widget_2);
        btnStart->setObjectName("btnStart");

        horizontalLayout_2->addWidget(btnStart);

        rdbtnTun = new QRadioButton(widget_2);
        rdbtnTun->setObjectName("rdbtnTun");

        horizontalLayout_2->addWidget(rdbtnTun);

        rdbtnSysProxy = new QRadioButton(widget_2);
        rdbtnSysProxy->setObjectName("rdbtnSysProxy");

        horizontalLayout_2->addWidget(rdbtnSysProxy);

        horizontalSpacer_2 = new QSpacerItem(40, 20, QSizePolicy::Policy::Expanding, QSizePolicy::Policy::Minimum);

        horizontalLayout_2->addItem(horizontalSpacer_2);


        verticalLayout->addWidget(widget_2);

        MainWindow->setCentralWidget(centralwidget);
        statusbar = new QStatusBar(MainWindow);
        statusbar->setObjectName("statusbar");
        MainWindow->setStatusBar(statusbar);

        retranslateUi(MainWindow);

        QMetaObject::connectSlotsByName(MainWindow);
    } // setupUi

    void retranslateUi(QMainWindow *MainWindow)
    {
        MainWindow->setWindowTitle(QCoreApplication::translate("MainWindow", "UI_TITLE_123", nullptr));
        btnAddSubscription->setText(QCoreApplication::translate("MainWindow", "Add", nullptr));
        btnUpdate->setText(QCoreApplication::translate("MainWindow", "Update", nullptr));
        QTreeWidgetItem *___qtreewidgetitem = treeWidget->headerItem();
        ___qtreewidgetitem->setText(7, QCoreApplication::translate("MainWindow", "Speed", nullptr));
        ___qtreewidgetitem->setText(6, QCoreApplication::translate("MainWindow", "TLS", nullptr));
        ___qtreewidgetitem->setText(5, QCoreApplication::translate("MainWindow", "Port", nullptr));
        ___qtreewidgetitem->setText(4, QCoreApplication::translate("MainWindow", "New Column", nullptr));
        ___qtreewidgetitem->setText(3, QCoreApplication::translate("MainWindow", "Address", nullptr));
        ___qtreewidgetitem->setText(2, QCoreApplication::translate("MainWindow", "Protocol", nullptr));
        ___qtreewidgetitem->setText(1, QCoreApplication::translate("MainWindow", "Type", nullptr));
        ___qtreewidgetitem->setText(0, QCoreApplication::translate("MainWindow", "Name", nullptr));
        btnStart->setText(QString());
        rdbtnTun->setText(QCoreApplication::translate("MainWindow", "TUN", nullptr));
        rdbtnSysProxy->setText(QCoreApplication::translate("MainWindow", "System Proxy", nullptr));
    } // retranslateUi

};

namespace Ui {
    class MainWindow: public Ui_MainWindow {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_MAINWINDOW_H
