from PySide6.QtWidgets import QDialog, QMainWindow
from design.mainwindow import Ui_MainWindow
from design.addsubscriptiondialog import Ui_Dialog
from generated import proxy_pb2_grpc
import grpc
from PySide6.QtCore import Qt

class AddDialog(QDialog, Ui_Dialog):
    def __init__(self, parent=None):
        super().__init__(parent)
        self.setupUi(self)

        self.lineEdit.textChanged.connect(lambda: self.labelError.clear())
        self.setAttribute(Qt.WidgetAttribute.WA_DeleteOnClose)
    def accept(self):
        if self.lineEdit.text() == "":
            self.labelError.setText("field cannot be empty!")
        else:
            super().accept()
class MainWindow(QMainWindow, Ui_MainWindow):
    def __init__(self):
        super().__init__()
        self.setupUi(self)

        self.channel = grpc.insecure_channel('localhost:30000')
        self.stub = proxy_pb2_grpc.NodeServiceStub(self.channel)

        self.btnAddSubscription.clicked.connect(self.add)
    def add(self):
        dialog = AddDialog(self)
        if dialog.exec():
            url = dialog.lineEdit.text()

            if url.startswith(("http", "https")):
                print("its url")
            else:
                print("other")
        else:
            return
        