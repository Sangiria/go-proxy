# -*- coding: utf-8 -*-

################################################################################
## Form generated from reading UI file 'editsubdialog.ui'
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
from PySide6.QtWidgets import (QAbstractButton, QApplication, QDialog, QDialogButtonBox,
    QGridLayout, QHBoxLayout, QLabel, QLineEdit,
    QSizePolicy, QTextEdit, QVBoxLayout, QWidget)

class Ui_Dialog(object):
    def setupUi(self, Dialog):
        if not Dialog.objectName():
            Dialog.setObjectName(u"Dialog")
        Dialog.resize(503, 189)
        Dialog.setMinimumSize(QSize(503, 189))
        Dialog.setMaximumSize(QSize(503, 189))
        self.verticalLayout = QVBoxLayout(Dialog)
        self.verticalLayout.setObjectName(u"verticalLayout")
        self.widget_2 = QWidget(Dialog)
        self.widget_2.setObjectName(u"widget_2")
        self.horizontalLayout = QHBoxLayout(self.widget_2)
        self.horizontalLayout.setObjectName(u"horizontalLayout")
        self.label = QLabel(self.widget_2)
        self.label.setObjectName(u"label")
        sizePolicy = QSizePolicy(QSizePolicy.Policy.Fixed, QSizePolicy.Policy.Preferred)
        sizePolicy.setHorizontalStretch(0)
        sizePolicy.setVerticalStretch(0)
        sizePolicy.setHeightForWidth(self.label.sizePolicy().hasHeightForWidth())
        self.label.setSizePolicy(sizePolicy)

        self.horizontalLayout.addWidget(self.label)

        self.labelError = QLabel(self.widget_2)
        self.labelError.setObjectName(u"labelError")
        self.labelError.setStyleSheet(u"color: rgb(255, 0, 0);")

        self.horizontalLayout.addWidget(self.labelError)


        self.verticalLayout.addWidget(self.widget_2)

        self.widget = QWidget(Dialog)
        self.widget.setObjectName(u"widget")
        self.gridLayout = QGridLayout(self.widget)
        self.gridLayout.setObjectName(u"gridLayout")
        self.name = QLabel(self.widget)
        self.name.setObjectName(u"name")

        self.gridLayout.addWidget(self.name, 0, 0, 1, 1)

        self.url = QLabel(self.widget)
        self.url.setObjectName(u"url")
        self.url.setAlignment(Qt.AlignLeading|Qt.AlignLeft|Qt.AlignTop)

        self.gridLayout.addWidget(self.url, 1, 0, 1, 1)

        self.nameEdit = QLineEdit(self.widget)
        self.nameEdit.setObjectName(u"nameEdit")
        sizePolicy1 = QSizePolicy(QSizePolicy.Policy.Expanding, QSizePolicy.Policy.Fixed)
        sizePolicy1.setHorizontalStretch(0)
        sizePolicy1.setVerticalStretch(0)
        sizePolicy1.setHeightForWidth(self.nameEdit.sizePolicy().hasHeightForWidth())
        self.nameEdit.setSizePolicy(sizePolicy1)

        self.gridLayout.addWidget(self.nameEdit, 0, 1, 1, 1)

        self.urlEdit = QTextEdit(self.widget)
        self.urlEdit.setObjectName(u"urlEdit")

        self.gridLayout.addWidget(self.urlEdit, 1, 1, 1, 1)


        self.verticalLayout.addWidget(self.widget)

        self.buttonBox = QDialogButtonBox(Dialog)
        self.buttonBox.setObjectName(u"buttonBox")
        self.buttonBox.setOrientation(Qt.Horizontal)
        self.buttonBox.setStandardButtons(QDialogButtonBox.Cancel|QDialogButtonBox.Ok)

        self.verticalLayout.addWidget(self.buttonBox)


        self.retranslateUi(Dialog)
        self.buttonBox.accepted.connect(Dialog.accept)
        self.buttonBox.rejected.connect(Dialog.reject)

        QMetaObject.connectSlotsByName(Dialog)
    # setupUi

    def retranslateUi(self, Dialog):
        Dialog.setWindowTitle(QCoreApplication.translate("Dialog", u"Edit Subscription", None))
        self.label.setText(QCoreApplication.translate("Dialog", u"Subscription: ", None))
        self.labelError.setText("")
        self.name.setText(QCoreApplication.translate("Dialog", u"Name", None))
        self.url.setText(QCoreApplication.translate("Dialog", u"URL", None))
    # retranslateUi

