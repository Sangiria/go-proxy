from PySide6.QtWidgets import QDialog, QMainWindow, QTreeWidgetItem
from PySide6.QtCore import Qt
from design.mainwindow import Ui_MainWindow
from design.addsubscriptiondialog import Ui_Dialog
from controller.node_controller import AddRequestController

class AddDialog(QDialog, Ui_Dialog):
    def __init__(self, parent=None):
        super().__init__(parent)
        self.setupUi(self)

        self.lineEdit.textChanged.connect(lambda: self.labelError.clear())
class MainWindow(QMainWindow, Ui_MainWindow):
    def __init__(self):
        super().__init__()
        self.setupUi(self)

        self.add_controller = AddRequestController(self)
        self.btnAddSubscription.clicked.connect(self.add_controller.handle_add)
    def add_node(self, item):
        row_data = [
            str(item.name),
            str(item.type),
            str(item.address),
            str(item.transport),
            str(item.port),
            str(item.tls)
        ]
        node = QTreeWidgetItem(row_data)
        self.treeWidget.addTopLevelItem(node)
    def add_sub(self, item):
        sub = QTreeWidgetItem([str(item.name)])
        self.treeWidget.addTopLevelItem(sub)

        for n in item.nodes:
            row_data = [
                str(n.name),
                str(n.type),
                str(n.address),
                str(n.transport),
                str(n.port),
                str(n.tls)
            ]

            node = QTreeWidgetItem(row_data)
            sub.addChild(node)

        self.treeWidget.expandItem(sub)