from PySide6.QtWidgets import QMainWindow, QTreeWidgetItem, QMenu, QTreeWidgetItemIterator
from PySide6.QtCore import Qt, QTimer, QEvent
from PySide6.QtGui import QIcon
from design.mainwindow import Ui_MainWindow
from view.notification import Notification
from handlers.node.add import AddHandler
from handlers.node.state import GetFullStateHandler
from handlers.node.get import GetHandler
from handlers.node.delete import DeleteHandler
from handlers.node.update import UpdateHandler
from handlers.proxy import ProxyControlHandler, StatusHandler
from pathlib import Path
import sys
from pathlib import Path

def get_resource_path():
    if hasattr(sys, '_MEIPASS'):
        return Path(sys._MEIPASS)
    return Path(__file__).resolve().parent.parent

ICON_PATH = get_resource_path() / "design" / "form" / "icon.png"

class MainWindow(QMainWindow, Ui_MainWindow):
    def __init__(self):
        super().__init__()
        self.setupUi(self)

        self.treeWidget.setColumnWidth(0, 200)
        self.typeCBox.installEventFilter(self)
        self.typeCBox.currentTextChanged.connect(self.filter_tree)

        self.add_handler = AddHandler(self)
        self.get_state_handler = GetFullStateHandler(self)
        self.proxy_control = ProxyControlHandler(self)
        self.status_monitor = StatusHandler(self)
        self.active_node_id = None

        self.btnStart.clicked.connect(self.on_start_clicked)
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
        self.status_monitor.start_monitoring()
        QTimer.singleShot(0, self.get_state_handler.handle_get_state)

    def closeEvent(self, event):
        if hasattr(self, 'status_monitor'):
            self.status_monitor.listener.stop()
            
        if hasattr(self, 'proxy_control') and hasattr(self.proxy_control, 'worker'):
            if self.proxy_control.worker and self.proxy_control.worker.isRunning():
                self.proxy_control.worker.quit()
                self.proxy_control.worker.wait()
        event.accept()

    def highlight_node(self, node_id):
        if not node_id: return
        self.active_node_id = str(node_id)
        
        self.reset_tree_styles()
        active_icon = QIcon(str(ICON_PATH))

        it = QTreeWidgetItemIterator(self.treeWidget)
        while it.value():
            item = it.value()
            if str(item.data(0, Qt.UserRole)) == self.active_node_id:
                item.setIcon(0, active_icon)
                self.treeWidget.scrollToItem(item)
                break
            it += 1
    
    def reset_tree_styles(self):
        it = QTreeWidgetItemIterator(self.treeWidget)
        while it.value():
            item = it.value()
            item.setIcon(0, QIcon())
            it += 1

    def change_button(self, state_type):
        self.btnStart.setProperty("state", state_type)

        if state_type == "connected":
            self.btnStart.setStyleSheet("background-color: #2ecc71; color: white;")
            self.btnStart.setEnabled(True)
        elif state_type == "connecting":
            self.btnStart.setStyleSheet("")
            self.btnStart.setEnabled(False)
        else:
            self.btnStart.setStyleSheet("") 
            self.btnStart.setEnabled(True)

    def on_start_clicked(self):
        if self.btnStart.property("state") == "connected":
            self.proxy_control.stop_proxy()
            return

        item = self.treeWidget.currentItem()
        if not item:
            self.show_notification("Please select node first", is_error=True)
            return

        self.proxy_control.start_proxy(item)

    def eventFilter(self, obj, event):
        if obj is self.typeCBox and event.type() == QEvent.Wheel:
            if not self.typeCBox.view().isVisible():
                return True
        
        return super().eventFilter(obj, event)

    def open_context_menu(self, position):
        if self.btnStart.property("state") == "connected":
            return
    
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
        
        if not action:
            return

        if update_action and action == update_action:
            self.update_handler = UpdateHandler(self)
            self.update_handler.handle_update(item)
            
        elif action == edit_action:
            self.get_handler = GetHandler(self)
            self.get_handler.handle_get(item)
            
        elif action == delete_action:
            self.delete_handler = DeleteHandler(self)
            self.delete_handler.handle_delete(item)

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