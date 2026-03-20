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
    QDialogButtonBox, QGridLayout, QHBoxLayout, QLabel,
    QLineEdit, QScrollArea, QSizePolicy, QSpinBox,
    QTextEdit, QWidget)

class Ui_Dialog(object):
    def setupUi(self, Dialog):
        if not Dialog.objectName():
            Dialog.setObjectName(u"Dialog")
        Dialog.resize(503, 352)
        Dialog.setMinimumSize(QSize(503, 352))
        Dialog.setMaximumSize(QSize(503, 352))
        self.gridLayout = QGridLayout(Dialog)
        self.gridLayout.setObjectName(u"gridLayout")
        self.scrollArea = QScrollArea(Dialog)
        self.scrollArea.setObjectName(u"scrollArea")
        self.scrollArea.setWidgetResizable(True)
        self.scrollAreaWidgetContents = QWidget()
        self.scrollAreaWidgetContents.setObjectName(u"scrollAreaWidgetContents")
        self.scrollAreaWidgetContents.setGeometry(QRect(0, 0, 466, 479))
        self.gridLayout_2 = QGridLayout(self.scrollAreaWidgetContents)
        self.gridLayout_2.setObjectName(u"gridLayout_2")
        self.securityCBox = QComboBox(self.scrollAreaWidgetContents)
        self.securityCBox.addItem("")
        self.securityCBox.addItem("")
        self.securityCBox.addItem("")
        self.securityCBox.setObjectName(u"securityCBox")
        sizePolicy = QSizePolicy(QSizePolicy.Policy.Fixed, QSizePolicy.Policy.Fixed)
        sizePolicy.setHorizontalStretch(0)
        sizePolicy.setVerticalStretch(0)
        sizePolicy.setHeightForWidth(self.securityCBox.sizePolicy().hasHeightForWidth())
        self.securityCBox.setSizePolicy(sizePolicy)

        self.gridLayout_2.addWidget(self.securityCBox, 6, 1, 1, 1)

        self.address = QLabel(self.scrollAreaWidgetContents)
        self.address.setObjectName(u"address")

        self.gridLayout_2.addWidget(self.address, 1, 0, 1, 1)

        self.extra = QLabel(self.scrollAreaWidgetContents)
        self.extra.setObjectName(u"extra")

        self.gridLayout_2.addWidget(self.extra, 10, 0, 1, 1)

        self.addressEdit = QLineEdit(self.scrollAreaWidgetContents)
        self.addressEdit.setObjectName(u"addressEdit")

        self.gridLayout_2.addWidget(self.addressEdit, 1, 1, 1, 1)

        self.sniEdit = QLineEdit(self.scrollAreaWidgetContents)
        self.sniEdit.setObjectName(u"sniEdit")

        self.gridLayout_2.addWidget(self.sniEdit, 7, 1, 1, 1)

        self.transportCBox = QComboBox(self.scrollAreaWidgetContents)
        self.transportCBox.addItem("")
        self.transportCBox.addItem("")
        self.transportCBox.addItem("")
        self.transportCBox.addItem("")
        self.transportCBox.addItem("")
        self.transportCBox.addItem("")
        self.transportCBox.addItem("")
        self.transportCBox.addItem("")
        self.transportCBox.setObjectName(u"transportCBox")
        sizePolicy.setHeightForWidth(self.transportCBox.sizePolicy().hasHeightForWidth())
        self.transportCBox.setSizePolicy(sizePolicy)

        self.gridLayout_2.addWidget(self.transportCBox, 4, 1, 1, 2)

        self.nameEdit = QLineEdit(self.scrollAreaWidgetContents)
        self.nameEdit.setObjectName(u"nameEdit")

        self.gridLayout_2.addWidget(self.nameEdit, 0, 1, 1, 1)

        self.transport = QLabel(self.scrollAreaWidgetContents)
        self.transport.setObjectName(u"transport")

        self.gridLayout_2.addWidget(self.transport, 4, 0, 1, 1)

        self.name = QLabel(self.scrollAreaWidgetContents)
        self.name.setObjectName(u"name")

        self.gridLayout_2.addWidget(self.name, 0, 0, 1, 1)

        self.pbk = QLabel(self.scrollAreaWidgetContents)
        self.pbk.setObjectName(u"pbk")

        self.gridLayout_2.addWidget(self.pbk, 9, 0, 1, 1)

        self.mode = QLabel(self.scrollAreaWidgetContents)
        self.mode.setObjectName(u"mode")

        self.gridLayout_2.addWidget(self.mode, 5, 0, 1, 1)

        self.modeCBox = QComboBox(self.scrollAreaWidgetContents)
        self.modeCBox.addItem("")
        self.modeCBox.addItem("")
        self.modeCBox.addItem("")
        self.modeCBox.addItem("")
        self.modeCBox.setObjectName(u"modeCBox")
        sizePolicy.setHeightForWidth(self.modeCBox.sizePolicy().hasHeightForWidth())
        self.modeCBox.setSizePolicy(sizePolicy)

        self.gridLayout_2.addWidget(self.modeCBox, 5, 1, 1, 1)

        self.port = QLabel(self.scrollAreaWidgetContents)
        self.port.setObjectName(u"port")

        self.gridLayout_2.addWidget(self.port, 2, 0, 1, 1)

        self.security = QLabel(self.scrollAreaWidgetContents)
        self.security.setObjectName(u"security")

        self.gridLayout_2.addWidget(self.security, 6, 0, 1, 1)

        self.uuid = QLabel(self.scrollAreaWidgetContents)
        self.uuid.setObjectName(u"uuid")

        self.gridLayout_2.addWidget(self.uuid, 3, 0, 1, 1)

        self.pbkEdit = QLineEdit(self.scrollAreaWidgetContents)
        self.pbkEdit.setObjectName(u"pbkEdit")

        self.gridLayout_2.addWidget(self.pbkEdit, 9, 1, 1, 1)

        self.sni = QLabel(self.scrollAreaWidgetContents)
        self.sni.setObjectName(u"sni")

        self.gridLayout_2.addWidget(self.sni, 7, 0, 1, 1)

        self.extraEdit = QTextEdit(self.scrollAreaWidgetContents)
        self.extraEdit.setObjectName(u"extraEdit")

        self.gridLayout_2.addWidget(self.extraEdit, 10, 1, 1, 1)

        self.uuidEdit = QLineEdit(self.scrollAreaWidgetContents)
        self.uuidEdit.setObjectName(u"uuidEdit")

        self.gridLayout_2.addWidget(self.uuidEdit, 3, 1, 1, 1)

        self.fp = QLabel(self.scrollAreaWidgetContents)
        self.fp.setObjectName(u"fp")

        self.gridLayout_2.addWidget(self.fp, 8, 0, 1, 1)

        self.fpCBox = QComboBox(self.scrollAreaWidgetContents)
        self.fpCBox.addItem("")
        self.fpCBox.addItem("")
        self.fpCBox.addItem("")
        self.fpCBox.addItem("")
        self.fpCBox.addItem("")
        self.fpCBox.addItem("")
        self.fpCBox.addItem("")
        self.fpCBox.addItem("")
        self.fpCBox.addItem("")
        self.fpCBox.addItem("")
        self.fpCBox.setObjectName(u"fpCBox")
        sizePolicy.setHeightForWidth(self.fpCBox.sizePolicy().hasHeightForWidth())
        self.fpCBox.setSizePolicy(sizePolicy)

        self.gridLayout_2.addWidget(self.fpCBox, 8, 1, 1, 1)

        self.portSBox = QSpinBox(self.scrollAreaWidgetContents)
        self.portSBox.setObjectName(u"portSBox")
        sizePolicy.setHeightForWidth(self.portSBox.sizePolicy().hasHeightForWidth())
        self.portSBox.setSizePolicy(sizePolicy)
        self.portSBox.setMinimum(1)
        self.portSBox.setMaximum(65535)

        self.gridLayout_2.addWidget(self.portSBox, 2, 1, 1, 1)

        self.scrollArea.setWidget(self.scrollAreaWidgetContents)

        self.gridLayout.addWidget(self.scrollArea, 1, 0, 1, 1)

        self.buttonBox = QDialogButtonBox(Dialog)
        self.buttonBox.setObjectName(u"buttonBox")
        self.buttonBox.setOrientation(Qt.Horizontal)
        self.buttonBox.setStandardButtons(QDialogButtonBox.Cancel|QDialogButtonBox.Ok)

        self.gridLayout.addWidget(self.buttonBox, 2, 0, 1, 1)

        self.widget = QWidget(Dialog)
        self.widget.setObjectName(u"widget")
        self.horizontalLayout = QHBoxLayout(self.widget)
        self.horizontalLayout.setObjectName(u"horizontalLayout")
        self.label = QLabel(self.widget)
        self.label.setObjectName(u"label")
        sizePolicy1 = QSizePolicy(QSizePolicy.Policy.Fixed, QSizePolicy.Policy.Preferred)
        sizePolicy1.setHorizontalStretch(0)
        sizePolicy1.setVerticalStretch(0)
        sizePolicy1.setHeightForWidth(self.label.sizePolicy().hasHeightForWidth())
        self.label.setSizePolicy(sizePolicy1)

        self.horizontalLayout.addWidget(self.label)

        self.labelError = QLabel(self.widget)
        self.labelError.setObjectName(u"labelError")
        self.labelError.setStyleSheet(u"color: rgb(255, 0, 0);")

        self.horizontalLayout.addWidget(self.labelError)


        self.gridLayout.addWidget(self.widget, 0, 0, 1, 1)


        self.retranslateUi(Dialog)
        self.buttonBox.accepted.connect(Dialog.accept)
        self.buttonBox.rejected.connect(Dialog.reject)

        QMetaObject.connectSlotsByName(Dialog)
    # setupUi

    def retranslateUi(self, Dialog):
        Dialog.setWindowTitle(QCoreApplication.translate("Dialog", u"Edit Node", None))
        self.securityCBox.setItemText(0, "")
        self.securityCBox.setItemText(1, QCoreApplication.translate("Dialog", u"reality", None))
        self.securityCBox.setItemText(2, QCoreApplication.translate("Dialog", u"tls", None))

        self.address.setText(QCoreApplication.translate("Dialog", u"Address", None))
        self.extra.setText(QCoreApplication.translate("Dialog", u"Extra", None))
        self.transportCBox.setItemText(0, QCoreApplication.translate("Dialog", u"tcp", None))
        self.transportCBox.setItemText(1, QCoreApplication.translate("Dialog", u"kcp", None))
        self.transportCBox.setItemText(2, QCoreApplication.translate("Dialog", u"ws", None))
        self.transportCBox.setItemText(3, QCoreApplication.translate("Dialog", u"httpupgrade", None))
        self.transportCBox.setItemText(4, QCoreApplication.translate("Dialog", u"xttp", None))
        self.transportCBox.setItemText(5, QCoreApplication.translate("Dialog", u"h2", None))
        self.transportCBox.setItemText(6, QCoreApplication.translate("Dialog", u"quic", None))
        self.transportCBox.setItemText(7, QCoreApplication.translate("Dialog", u"grpc", None))

        self.transport.setText(QCoreApplication.translate("Dialog", u"Transport", None))
        self.name.setText(QCoreApplication.translate("Dialog", u"Name", None))
        self.pbk.setText(QCoreApplication.translate("Dialog", u"Public Key", None))
        self.mode.setText(QCoreApplication.translate("Dialog", u"XHTTP Mode", None))
        self.modeCBox.setItemText(0, QCoreApplication.translate("Dialog", u"auto", None))
        self.modeCBox.setItemText(1, QCoreApplication.translate("Dialog", u"packet-up", None))
        self.modeCBox.setItemText(2, QCoreApplication.translate("Dialog", u"stream-up", None))
        self.modeCBox.setItemText(3, QCoreApplication.translate("Dialog", u"stream-one", None))

        self.port.setText(QCoreApplication.translate("Dialog", u"Port", None))
        self.security.setText(QCoreApplication.translate("Dialog", u"Security", None))
        self.uuid.setText(QCoreApplication.translate("Dialog", u"UUID", None))
        self.sni.setText(QCoreApplication.translate("Dialog", u"SNI", None))
        self.fp.setText(QCoreApplication.translate("Dialog", u"Fingerprint", None))
        self.fpCBox.setItemText(0, QCoreApplication.translate("Dialog", u"chrome", None))
        self.fpCBox.setItemText(1, QCoreApplication.translate("Dialog", u"firefox", None))
        self.fpCBox.setItemText(2, QCoreApplication.translate("Dialog", u"safari", None))
        self.fpCBox.setItemText(3, QCoreApplication.translate("Dialog", u"ios", None))
        self.fpCBox.setItemText(4, QCoreApplication.translate("Dialog", u"android", None))
        self.fpCBox.setItemText(5, QCoreApplication.translate("Dialog", u"edge", None))
        self.fpCBox.setItemText(6, QCoreApplication.translate("Dialog", u"360", None))
        self.fpCBox.setItemText(7, QCoreApplication.translate("Dialog", u"qq", None))
        self.fpCBox.setItemText(8, QCoreApplication.translate("Dialog", u"random", None))
        self.fpCBox.setItemText(9, QCoreApplication.translate("Dialog", u"randomized", None))

        self.label.setText(QCoreApplication.translate("Dialog", u"Node: ", None))
        self.labelError.setText("")
    # retranslateUi

