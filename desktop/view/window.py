from PySide6.QtWidgets import QMainWindow, QTreeWidgetItem, QMenu, QTreeWidgetItemIterator
from PySide6.QtCore import Qt, QTimer
from design.mainwindow import Ui_MainWindow
from handlers.node_handler import AddHandler, GetFullStateHandler, GetHandler

class MainWindow(QMainWindow, Ui_MainWindow):
    def __init__(self):
        super().__init__()
        self.setupUi(self)

        self.add_handler = AddHandler(self)
        self.get_state_handler = GetFullStateHandler(self)

        self.btnAddSubscription.clicked.connect(self.add_handler.handle_add)
        self.treeWidget.setContextMenuPolicy(Qt.CustomContextMenu)
        self.treeWidget.customContextMenuRequested.connect(self.open_context_menu)

        QTimer.singleShot(0, self.get_state_handler.handle_get_state)
    def open_context_menu(self, position):
        item = self.treeWidget.itemAt(position)
        if not item:
            return

        menu = QMenu()
        edit_action = menu.addAction("Edit")
        delete_action = menu.addAction("Delete")
    
        action = menu.exec(self.treeWidget.viewport().mapToGlobal(position))
        if action == edit_action:
            self.get_handler = GetHandler(self)
            self.get_handler.handle_get(item)
        if action == delete_action:
            pass
    def create_node_item(self, item, item_type="node"):
        if item_type == "node":
            row_data = [
                str(item.name), str(item.type), str(item.address),
                str(item.transport), str(item.port), str(item.security)
            ]
        else:
            row_data = [str(item.name)]

        node = QTreeWidgetItem(row_data)
        node.setData(0, Qt.UserRole, item.id)
        node.setData(0, Qt.UserRole + 1, item_type)
        return node

    def add_node(self, item):
        node = self.create_node_item(item, "node")
        self.treeWidget.addTopLevelItem(node)

    def add_sub(self, item):
        sub = self.create_node_item(item, "sub")
        self.treeWidget.addTopLevelItem(sub)

        for n in item.nodes:
            node = self.create_node_item(n, "node")
            sub.addChild(node)

        self.treeWidget.expandItem(sub)
    def update_item(self, item_data):
        iterator = QTreeWidgetItemIterator(self.treeWidget)
        while iterator.value():
            item = iterator.value()
            if item.data(0, Qt.UserRole) == item_data.id:
                item.setText(0, str(item_data.name))
                if item.data(0, Qt.UserRole + 1) == "node":
                    item.setText(0, str(item_data.name))
                    item.setText(2, str(item_data.address))
                    item.setText(3, str(item_data.transport))
                    item.setText(4, str(item_data.port))
                    item.setText(5, str(item_data.security))
                if item.data(0, Qt.UserRole + 1) == "sub":
                    item.setText(0, str(item_data.name))
                return
            iterator += 1