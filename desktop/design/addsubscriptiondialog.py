# -*- coding: utf-8 -*-

################################################################################
## Form generated from reading UI file 'addsubscriptiondialog.ui'
##
## Created by: Qt User Interface Compiler version 6.10.2
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
    QHBoxLayout, QLabel, QLineEdit, QSizePolicy,
    QVBoxLayout, QWidget)

class Ui_Dialog(object):
    def setupUi(self, Dialog):
        if not Dialog.objectName():
            Dialog.setObjectName(u"Dialog")
        Dialog.resize(627, 151)
        Dialog.setMinimumSize(QSize(627, 151))
        Dialog.setMaximumSize(QSize(627, 151))
        self.verticalLayout = QVBoxLayout(Dialog)
        self.verticalLayout.setObjectName(u"verticalLayout")
        self.widget = QWidget(Dialog)
        self.widget.setObjectName(u"widget")
        self.horizontalLayout_2 = QHBoxLayout(self.widget)
        self.horizontalLayout_2.setObjectName(u"horizontalLayout_2")
        self.labelMessage = QLabel(self.widget)
        self.labelMessage.setObjectName(u"labelMessage")

        self.horizontalLayout_2.addWidget(self.labelMessage)

        self.labelError = QLabel(self.widget)
        self.labelError.setObjectName(u"labelError")
        font = QFont()
        font.setBold(False)
        self.labelError.setFont(font)
        self.labelError.setStyleSheet(u"color: rgb(255, 0, 0);")

        self.horizontalLayout_2.addWidget(self.labelError)

        self.horizontalLayout_2.setStretch(1, 1)

        self.verticalLayout.addWidget(self.widget)

        self.horizontalLayout = QHBoxLayout()
        self.horizontalLayout.setObjectName(u"horizontalLayout")
        self.labelUrl = QLabel(Dialog)
        self.labelUrl.setObjectName(u"labelUrl")

        self.horizontalLayout.addWidget(self.labelUrl)

        self.lineEdit = QLineEdit(Dialog)
        self.lineEdit.setObjectName(u"lineEdit")

        self.horizontalLayout.addWidget(self.lineEdit)


        self.verticalLayout.addLayout(self.horizontalLayout)

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
        Dialog.setWindowTitle(QCoreApplication.translate("Dialog", u"Dialog", None))
        self.labelMessage.setText(QCoreApplication.translate("Dialog", u"Enter subscription url:", None))
        self.labelError.setText("")
        self.labelUrl.setText(QCoreApplication.translate("Dialog", u"URL:", None))
    # retranslateUi

