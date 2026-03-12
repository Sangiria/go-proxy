# -*- coding: utf-8 -*-

################################################################################
## Form generated from reading UI file 'mainwindow.ui'
##
## Created by: Qt User Interface Compiler version 6.9.2
##
## WARNING! All changes made in this file will be lost when recompiling UI file!
################################################################################

from PySide6.QtCore import (QCoreApplication, QDate, QDateTime, QLocale,
    QMetaObject, QObject, QPoint, QRect,
    QSize, QTime, QUrl, Qt)
from PySide6.QtGui import (QBrush, QColor, QConicalGradient, QCursor,
    QFont, QFontDatabase, QGradient, QIcon,
    QImage, QKeySequence, QLinearGradient, QPainter,
    QPalette, QPixmap, QRadialGradient, QTransform)
from PySide6.QtWidgets import (QApplication, QHBoxLayout, QHeaderView, QMainWindow,
    QPushButton, QRadioButton, QSizePolicy, QSpacerItem,
    QStatusBar, QToolButton, QTreeWidget, QTreeWidgetItem,
    QVBoxLayout, QWidget)

class Ui_MainWindow(object):
    def setupUi(self, MainWindow):
        if not MainWindow.objectName():
            MainWindow.setObjectName(u"MainWindow")
        MainWindow.resize(820, 599)
        MainWindow.setMinimumSize(QSize(716, 0))
        MainWindow.setMaximumSize(QSize(820, 600))
        self.centralwidget = QWidget(MainWindow)
        self.centralwidget.setObjectName(u"centralwidget")
        self.verticalLayout = QVBoxLayout(self.centralwidget)
        self.verticalLayout.setObjectName(u"verticalLayout")
        self.widget = QWidget(self.centralwidget)
        self.widget.setObjectName(u"widget")
        self.widget.setStyleSheet(u"")
        self.horizontalLayout = QHBoxLayout(self.widget)
        self.horizontalLayout.setObjectName(u"horizontalLayout")
        self.btnAddSubscription = QPushButton(self.widget)
        self.btnAddSubscription.setObjectName(u"btnAddSubscription")
        self.btnAddSubscription.setStyleSheet(u"")

        self.horizontalLayout.addWidget(self.btnAddSubscription)

        self.btnUpdate = QPushButton(self.widget)
        self.btnUpdate.setObjectName(u"btnUpdate")

        self.horizontalLayout.addWidget(self.btnUpdate)

        self.horizontalSpacer = QSpacerItem(40, 20, QSizePolicy.Policy.Expanding, QSizePolicy.Policy.Minimum)

        self.horizontalLayout.addItem(self.horizontalSpacer)


        self.verticalLayout.addWidget(self.widget)

        self.treeWidget = QTreeWidget(self.centralwidget)
        self.treeWidget.setObjectName(u"treeWidget")

        self.verticalLayout.addWidget(self.treeWidget)

        self.widget_2 = QWidget(self.centralwidget)
        self.widget_2.setObjectName(u"widget_2")
        self.horizontalLayout_2 = QHBoxLayout(self.widget_2)
        self.horizontalLayout_2.setObjectName(u"horizontalLayout_2")
        self.btnStart = QToolButton(self.widget_2)
        self.btnStart.setObjectName(u"btnStart")

        self.horizontalLayout_2.addWidget(self.btnStart)

        self.rdbtnTun = QRadioButton(self.widget_2)
        self.rdbtnTun.setObjectName(u"rdbtnTun")
        self.rdbtnTun.setAutoExclusive(False)

        self.horizontalLayout_2.addWidget(self.rdbtnTun)

        self.rdbtnSysProxy = QRadioButton(self.widget_2)
        self.rdbtnSysProxy.setObjectName(u"rdbtnSysProxy")
        self.rdbtnSysProxy.setAutoExclusive(False)

        self.horizontalLayout_2.addWidget(self.rdbtnSysProxy)

        self.horizontalSpacer_2 = QSpacerItem(40, 20, QSizePolicy.Policy.Expanding, QSizePolicy.Policy.Minimum)

        self.horizontalLayout_2.addItem(self.horizontalSpacer_2)


        self.verticalLayout.addWidget(self.widget_2)

        MainWindow.setCentralWidget(self.centralwidget)
        self.statusbar = QStatusBar(MainWindow)
        self.statusbar.setObjectName(u"statusbar")
        MainWindow.setStatusBar(self.statusbar)

        self.retranslateUi(MainWindow)

        QMetaObject.connectSlotsByName(MainWindow)
    # setupUi

    def retranslateUi(self, MainWindow):
        MainWindow.setWindowTitle(QCoreApplication.translate("MainWindow", u"UI_TITLE_123", None))
        self.btnAddSubscription.setText(QCoreApplication.translate("MainWindow", u"Add", None))
        self.btnUpdate.setText(QCoreApplication.translate("MainWindow", u"Update", None))
        ___qtreewidgetitem = self.treeWidget.headerItem()
        ___qtreewidgetitem.setText(7, QCoreApplication.translate("MainWindow", u"Speed", None));
        ___qtreewidgetitem.setText(6, QCoreApplication.translate("MainWindow", u"TLS", None));
        ___qtreewidgetitem.setText(5, QCoreApplication.translate("MainWindow", u"Port", None));
        ___qtreewidgetitem.setText(4, QCoreApplication.translate("MainWindow", u"Transport", None));
        ___qtreewidgetitem.setText(3, QCoreApplication.translate("MainWindow", u"Address", None));
        ___qtreewidgetitem.setText(2, QCoreApplication.translate("MainWindow", u"Protocol", None));
        ___qtreewidgetitem.setText(1, QCoreApplication.translate("MainWindow", u"Type", None));
        ___qtreewidgetitem.setText(0, QCoreApplication.translate("MainWindow", u"Name", None));
        self.btnStart.setText("")
        self.rdbtnTun.setText(QCoreApplication.translate("MainWindow", u"TUN", None))
        self.rdbtnSysProxy.setText(QCoreApplication.translate("MainWindow", u"System Proxy", None))
    # retranslateUi

