# -*- coding: utf-8 -*-

################################################################################
## Form generated from reading UI file 'editnodedialog.ui'
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
from PySide6.QtWidgets import (QAbstractButton, QApplication, QComboBox, QDialog,
    QDialogButtonBox, QGridLayout, QLabel, QLineEdit,
    QScrollArea, QSizePolicy, QTextEdit, QWidget)

class Ui_Dialog(object):
    def setupUi(self, Dialog):
        if not Dialog.objectName():
            Dialog.setObjectName(u"Dialog")
        Dialog.resize(503, 352)
        Dialog.setMinimumSize(QSize(503, 352))
        Dialog.setMaximumSize(QSize(503, 352))
        self.gridLayout = QGridLayout(Dialog)
        self.gridLayout.setObjectName(u"gridLayout")
        self.buttonBox = QDialogButtonBox(Dialog)
        self.buttonBox.setObjectName(u"buttonBox")
        self.buttonBox.setOrientation(Qt.Horizontal)
        self.buttonBox.setStandardButtons(QDialogButtonBox.Cancel|QDialogButtonBox.Ok)

        self.gridLayout.addWidget(self.buttonBox, 1, 0, 1, 1)

        self.scrollArea = QScrollArea(Dialog)
        self.scrollArea.setObjectName(u"scrollArea")
        self.scrollArea.setWidgetResizable(True)
        self.scrollAreaWidgetContents = QWidget()
        self.scrollAreaWidgetContents.setObjectName(u"scrollAreaWidgetContents")
        self.scrollAreaWidgetContents.setGeometry(QRect(0, 0, 466, 479))
        self.gridLayout_2 = QGridLayout(self.scrollAreaWidgetContents)
        self.gridLayout_2.setObjectName(u"gridLayout_2")
        self.comboBox_2 = QComboBox(self.scrollAreaWidgetContents)
        self.comboBox_2.addItem("")
        self.comboBox_2.addItem("")
        self.comboBox_2.setObjectName(u"comboBox_2")
        sizePolicy = QSizePolicy(QSizePolicy.Policy.Fixed, QSizePolicy.Policy.Fixed)
        sizePolicy.setHorizontalStretch(0)
        sizePolicy.setVerticalStretch(0)
        sizePolicy.setHeightForWidth(self.comboBox_2.sizePolicy().hasHeightForWidth())
        self.comboBox_2.setSizePolicy(sizePolicy)

        self.gridLayout_2.addWidget(self.comboBox_2, 6, 1, 1, 1)

        self.label_3 = QLabel(self.scrollAreaWidgetContents)
        self.label_3.setObjectName(u"label_3")

        self.gridLayout_2.addWidget(self.label_3, 1, 0, 1, 1)

        self.label_10 = QLabel(self.scrollAreaWidgetContents)
        self.label_10.setObjectName(u"label_10")

        self.gridLayout_2.addWidget(self.label_10, 10, 0, 1, 1)

        self.lineEdit_2 = QLineEdit(self.scrollAreaWidgetContents)
        self.lineEdit_2.setObjectName(u"lineEdit_2")

        self.gridLayout_2.addWidget(self.lineEdit_2, 1, 1, 1, 1)

        self.lineEdit_3 = QLineEdit(self.scrollAreaWidgetContents)
        self.lineEdit_3.setObjectName(u"lineEdit_3")
        sizePolicy.setHeightForWidth(self.lineEdit_3.sizePolicy().hasHeightForWidth())
        self.lineEdit_3.setSizePolicy(sizePolicy)

        self.gridLayout_2.addWidget(self.lineEdit_3, 2, 1, 1, 1)

        self.lineEdit_5 = QLineEdit(self.scrollAreaWidgetContents)
        self.lineEdit_5.setObjectName(u"lineEdit_5")

        self.gridLayout_2.addWidget(self.lineEdit_5, 7, 1, 1, 1)

        self.comboBox = QComboBox(self.scrollAreaWidgetContents)
        self.comboBox.addItem("")
        self.comboBox.addItem("")
        self.comboBox.addItem("")
        self.comboBox.addItem("")
        self.comboBox.addItem("")
        self.comboBox.addItem("")
        self.comboBox.addItem("")
        self.comboBox.addItem("")
        self.comboBox.setObjectName(u"comboBox")
        sizePolicy.setHeightForWidth(self.comboBox.sizePolicy().hasHeightForWidth())
        self.comboBox.setSizePolicy(sizePolicy)

        self.gridLayout_2.addWidget(self.comboBox, 4, 1, 1, 2)

        self.lineEdit = QLineEdit(self.scrollAreaWidgetContents)
        self.lineEdit.setObjectName(u"lineEdit")

        self.gridLayout_2.addWidget(self.lineEdit, 0, 1, 1, 1)

        self.label_5 = QLabel(self.scrollAreaWidgetContents)
        self.label_5.setObjectName(u"label_5")

        self.gridLayout_2.addWidget(self.label_5, 4, 0, 1, 1)

        self.label_2 = QLabel(self.scrollAreaWidgetContents)
        self.label_2.setObjectName(u"label_2")

        self.gridLayout_2.addWidget(self.label_2, 0, 0, 1, 1)

        self.label_8 = QLabel(self.scrollAreaWidgetContents)
        self.label_8.setObjectName(u"label_8")

        self.gridLayout_2.addWidget(self.label_8, 9, 0, 1, 1)

        self.label_9 = QLabel(self.scrollAreaWidgetContents)
        self.label_9.setObjectName(u"label_9")

        self.gridLayout_2.addWidget(self.label_9, 5, 0, 1, 1)

        self.comboBox_3 = QComboBox(self.scrollAreaWidgetContents)
        self.comboBox_3.addItem("")
        self.comboBox_3.addItem("")
        self.comboBox_3.addItem("")
        self.comboBox_3.addItem("")
        self.comboBox_3.setObjectName(u"comboBox_3")
        sizePolicy.setHeightForWidth(self.comboBox_3.sizePolicy().hasHeightForWidth())
        self.comboBox_3.setSizePolicy(sizePolicy)

        self.gridLayout_2.addWidget(self.comboBox_3, 5, 1, 1, 1)

        self.label_4 = QLabel(self.scrollAreaWidgetContents)
        self.label_4.setObjectName(u"label_4")

        self.gridLayout_2.addWidget(self.label_4, 2, 0, 1, 1)

        self.label_6 = QLabel(self.scrollAreaWidgetContents)
        self.label_6.setObjectName(u"label_6")

        self.gridLayout_2.addWidget(self.label_6, 6, 0, 1, 1)

        self.label = QLabel(self.scrollAreaWidgetContents)
        self.label.setObjectName(u"label")

        self.gridLayout_2.addWidget(self.label, 3, 0, 1, 1)

        self.lineEdit_6 = QLineEdit(self.scrollAreaWidgetContents)
        self.lineEdit_6.setObjectName(u"lineEdit_6")

        self.gridLayout_2.addWidget(self.lineEdit_6, 9, 1, 1, 1)

        self.label_7 = QLabel(self.scrollAreaWidgetContents)
        self.label_7.setObjectName(u"label_7")

        self.gridLayout_2.addWidget(self.label_7, 7, 0, 1, 1)

        self.textEdit = QTextEdit(self.scrollAreaWidgetContents)
        self.textEdit.setObjectName(u"textEdit")

        self.gridLayout_2.addWidget(self.textEdit, 10, 1, 1, 1)

        self.lineEdit_4 = QLineEdit(self.scrollAreaWidgetContents)
        self.lineEdit_4.setObjectName(u"lineEdit_4")

        self.gridLayout_2.addWidget(self.lineEdit_4, 3, 1, 1, 1)

        self.label_11 = QLabel(self.scrollAreaWidgetContents)
        self.label_11.setObjectName(u"label_11")

        self.gridLayout_2.addWidget(self.label_11, 8, 0, 1, 1)

        self.comboBox_4 = QComboBox(self.scrollAreaWidgetContents)
        self.comboBox_4.addItem("")
        self.comboBox_4.addItem("")
        self.comboBox_4.addItem("")
        self.comboBox_4.addItem("")
        self.comboBox_4.addItem("")
        self.comboBox_4.addItem("")
        self.comboBox_4.addItem("")
        self.comboBox_4.addItem("")
        self.comboBox_4.addItem("")
        self.comboBox_4.addItem("")
        self.comboBox_4.setObjectName(u"comboBox_4")
        sizePolicy.setHeightForWidth(self.comboBox_4.sizePolicy().hasHeightForWidth())
        self.comboBox_4.setSizePolicy(sizePolicy)

        self.gridLayout_2.addWidget(self.comboBox_4, 8, 1, 1, 1)

        self.scrollArea.setWidget(self.scrollAreaWidgetContents)

        self.gridLayout.addWidget(self.scrollArea, 0, 0, 1, 1)


        self.retranslateUi(Dialog)
        self.buttonBox.accepted.connect(Dialog.accept)
        self.buttonBox.rejected.connect(Dialog.reject)

        QMetaObject.connectSlotsByName(Dialog)
    # setupUi

    def retranslateUi(self, Dialog):
        Dialog.setWindowTitle(QCoreApplication.translate("Dialog", u"Edit Node", None))
        self.comboBox_2.setItemText(0, QCoreApplication.translate("Dialog", u"reality", None))
        self.comboBox_2.setItemText(1, QCoreApplication.translate("Dialog", u"tls", None))

        self.label_3.setText(QCoreApplication.translate("Dialog", u"Address", None))
        self.label_10.setText(QCoreApplication.translate("Dialog", u"Extra", None))
        self.comboBox.setItemText(0, QCoreApplication.translate("Dialog", u"tcp", None))
        self.comboBox.setItemText(1, QCoreApplication.translate("Dialog", u"kcp", None))
        self.comboBox.setItemText(2, QCoreApplication.translate("Dialog", u"ws", None))
        self.comboBox.setItemText(3, QCoreApplication.translate("Dialog", u"httpupgrade", None))
        self.comboBox.setItemText(4, QCoreApplication.translate("Dialog", u"xttp", None))
        self.comboBox.setItemText(5, QCoreApplication.translate("Dialog", u"h2", None))
        self.comboBox.setItemText(6, QCoreApplication.translate("Dialog", u"quic", None))
        self.comboBox.setItemText(7, QCoreApplication.translate("Dialog", u"grpc", None))

        self.label_5.setText(QCoreApplication.translate("Dialog", u"Transport", None))
        self.label_2.setText(QCoreApplication.translate("Dialog", u"Name", None))
        self.label_8.setText(QCoreApplication.translate("Dialog", u"Public Key", None))
        self.label_9.setText(QCoreApplication.translate("Dialog", u"Transport Mode", None))
        self.comboBox_3.setItemText(0, QCoreApplication.translate("Dialog", u"auto", None))
        self.comboBox_3.setItemText(1, QCoreApplication.translate("Dialog", u"packet-up", None))
        self.comboBox_3.setItemText(2, QCoreApplication.translate("Dialog", u"stream-up", None))
        self.comboBox_3.setItemText(3, QCoreApplication.translate("Dialog", u"stream-one", None))

        self.label_4.setText(QCoreApplication.translate("Dialog", u"Port", None))
        self.label_6.setText(QCoreApplication.translate("Dialog", u"Security", None))
        self.label.setText(QCoreApplication.translate("Dialog", u"UUID", None))
        self.label_7.setText(QCoreApplication.translate("Dialog", u"SNI", None))
        self.label_11.setText(QCoreApplication.translate("Dialog", u"Fingerprint", None))
        self.comboBox_4.setItemText(0, QCoreApplication.translate("Dialog", u"chrome", None))
        self.comboBox_4.setItemText(1, QCoreApplication.translate("Dialog", u"firefox", None))
        self.comboBox_4.setItemText(2, QCoreApplication.translate("Dialog", u"safari", None))
        self.comboBox_4.setItemText(3, QCoreApplication.translate("Dialog", u"ios", None))
        self.comboBox_4.setItemText(4, QCoreApplication.translate("Dialog", u"android", None))
        self.comboBox_4.setItemText(5, QCoreApplication.translate("Dialog", u"edge", None))
        self.comboBox_4.setItemText(6, QCoreApplication.translate("Dialog", u"360", None))
        self.comboBox_4.setItemText(7, QCoreApplication.translate("Dialog", u"qq", None))
        self.comboBox_4.setItemText(8, QCoreApplication.translate("Dialog", u"random", None))
        self.comboBox_4.setItemText(9, QCoreApplication.translate("Dialog", u"randomized", None))

    # retranslateUi

