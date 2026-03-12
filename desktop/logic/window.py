from PySide6.QtWidgets import QDialog, QMainWindow, QMessageBox
from PySide6.QtCore import Qt

from design.mainwindow import Ui_MainWindow
from design.addsubscriptiondialog import Ui_Dialog
from service.worker import GrpcWorker, stub
from generated import proxy_pb2

class AddDialog(QDialog, Ui_Dialog):
    def __init__(self, parent=None):
        super().__init__(parent)
        self.setupUi(self)
        self.url = ""
        self.lineEdit.textChanged.connect(lambda: self.labelError.clear())
        self.setAttribute(Qt.WidgetAttribute.WA_DeleteOnClose)
    def accept(self):
        self.url = self.lineEdit.text()
        if self.url == "":
            self.labelError.setText("field cannot be empty!")
        else:
            super().accept()
class MainWindow(QMainWindow, Ui_MainWindow):
    def __init__(self):
        super().__init__()
        self.setupUi(self)

        self.btnAddSubscription.clicked.connect(self.add)
    def add(self):
        dialog = AddDialog(self)
        if dialog.exec():
            self.btnAddSubscription.setEnabled(False)
            url = dialog.url.strip()

            if url.startswith(("http", "https")):
                self.worker = GrpcWorker(stub.AddSubscription, proxy_pb2.Url(url=url))
                self.worker.finished.connect(self.on_add_success)
                self.worker.error.connect(self.on_add_error)
                self.worker.start()
            else:
                print("other")
        else:
            return
    def on_add_success(self, response):
        self.btnAddSubscription.setEnabled(True)
        self.inputUrl.clear()

        print(f"Сервер ответил: {response.message}")

    def on_add_error(self, error_text):
        self.btnAddSubscription.setEnabled(True)
        QMessageBox.critical(self, "Ошибка gRPC", f"Не удалось добавить: {error_text}")
        