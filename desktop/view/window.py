from PySide6.QtWidgets import QMainWindow, QTreeWidgetItem, QMenu, QTreeWidgetItemIterator
from PySide6.QtCore import Qt, QTimer, QEvent
from design.mainwindow import Ui_MainWindow
from view.notification import Notification
from handlers.node.add import AddHandler
from handlers.node.state import GetFullStateHandler
from handlers.node.get import GetHandler
from handlers.node.delete import DeleteHandler
from handlers.node.update import UpdateHandler

class MainWindow(QMainWindow, Ui_MainWindow):
    def __init__(self):
        super().__init__()
        self.setupUi(self)

        self.treeWidget.setColumnWidth(0, 200)
        self.typeCBox.installEventFilter(self)
        self.typeCBox.currentTextChanged.connect(self.filter_tree)

        self.add_handler = AddHandler(self)
        self.get_state_handler = GetFullStateHandler(self)

        self.btnAdd.clicked.connect(self.add_handler.handle_add)
        self.treeWidget.setContextMenuPolicy(Qt.CustomContextMenu)
        self.treeWidget.customContextMenuRequested.connect(self.open_context_menu)

        self.columns = {
            "name": 0,
            "address": 2,
            "transport": 3,
            "port": 4,
            "security": 5
        }

        QTimer.singleShot(0, self.get_state_handler.handle_get_state)

    def eventFilter(self, obj, event):
        if obj is self.typeCBox and event.type() == QEvent.Wheel:
            if not self.typeCBox.view().isVisible():
                return True
        
        return super().eventFilter(obj, event)

    def open_context_menu(self, position):
        item = self.treeWidget.itemAt(position)
        if not item:
            return
        
        item_type = item.data(0, Qt.UserRole + 1)

        menu = QMenu()
        update_action = None
        if item_type == "sub":
            update_action = menu.addAction("Update")
            menu.addSeparator()

        edit_action = menu.addAction("Edit")
        delete_action = menu.addAction("Delete")
    
        action = menu.exec(self.treeWidget.viewport().mapToGlobal(position))
        if action == edit_action:
            self.get_handler = GetHandler(self)
            self.get_handler.handle_get(item)
        if action == delete_action:
            self.delete_handler = DeleteHandler(self)
            self.delete_handler.handle_delete(item)
        if action == update_action:
            self.update_handler = UpdateHandler(self)
            self.update_handler.handle_update(item)

    def update_sub_nodes(self, sub, nodes_data):
        sub.takeChildren()
        for n in nodes_data: 
            node = self.create_node_item(n, "node")
            sub.addChild(node)
        self.treeWidget.expandItem(sub)

    def show_notification(self, text, is_error=False):
        self.toast = Notification(self, text, is_error)
    def remove_item(self, item):
        parent = item.parent()
    
        if parent:
            parent.removeChild(item)
        else:
            index = self.treeWidget.indexOfTopLevelItem(item)
            if index != -1:
                self.treeWidget.takeTopLevelItem(index)
        del item

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
    
    def filter_tree(self):
        filter_value = self.typeCBox.currentText().lower()
    
        for i in range(self.treeWidget.topLevelItemCount()):
            item = self.treeWidget.topLevelItem(i)
            item_type = item.data(0, Qt.UserRole + 1)
        
            if filter_value == "all":
                item.setHidden(False)
            elif filter_value == "subscriptions":
                item.setHidden(item_type != "sub")
            elif filter_value == "manual":
                item.setHidden(item_type != "node")

    def add_node(self, item):
        node = self.create_node_item(item, "node")
        self.treeWidget.addTopLevelItem(node)

    def add_sub(self, item):
        sub = self.create_node_item(item, "sub")
        self.treeWidget.addTopLevelItem(sub)
        
        for n in item.nodes.nodes: 
            node = self.create_node_item(n, "node")
            sub.addChild(node)

        self.treeWidget.expandItem(sub)
    def update_item(self, data):
        iterator = QTreeWidgetItemIterator(self.treeWidget)
        target = None

        while iterator.value():
            item = iterator.value()
            if item.data(0, Qt.UserRole) == data.id:
                target = item
                break
            iterator += 1

        if not target:
            return
        
        for field, value in data.ListFields():
            attr = field.name
            if field.name in self.columns:
                col_idx = self.columns[attr]
                display_value = str(value)
                target.setText(col_idx, display_value)