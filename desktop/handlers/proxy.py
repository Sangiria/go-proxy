from PySide6.QtCore import Qt
from handlers.grpc import GrpcHandler
from model.worker import StatusListener, GrpcWorker, proxy_stub
from model.generated import proxy_pb2

class ProxyControlHandler(GrpcHandler):
    def __init__(self, view):
        super().__init__(view)
    def start_proxy(self, item):
        item_id = item.data(0, Qt.UserRole)
        id_args = {"id": str(item_id)}

        parent_item = item.parent()
        source_id = None
        if parent_item:
            source_id = parent_item.data(0, Qt.UserRole)
            id_args["source_id"] = str(source_id)

        self.worker = GrpcWorker(proxy_stub.StartProxy, proxy_pb2.Id(**id_args))
        self.worker.start()

    def stop_proxy(self):
        self.worker = GrpcWorker(proxy_stub.StopProxy, proxy_pb2.Null())
        self.worker.start()

class StatusHandler(GrpcHandler):
    def __init__(self, view):
        super().__init__(view)
        self.is_initialized = False
        self.listener = StatusListener()
        self.listener.status_updated.connect(self._parse_status)
        self.listener.connection_lost.connect(self._on_connection_lost)

    def start_monitoring(self):
        self.listener.start()

    def _parse_status(self, status):
        self.is_initialized = True
        state = status.state
        msg = status.message

        node_id = status.active_node_id
        self.view.active_node_id = node_id

        print(f"DEBUG: Received status {status.state}, node_id: '{node_id}'")

        if state == proxy_pb2.CONNECTED:
            self.view.change_button("connected")
            self.view.statusBar().showMessage("running")
            if node_id:
                self.view.highlight_node(node_id)
            
        elif state == proxy_pb2.CONNECTING:
            self.view.change_button("connecting")
            self.view.statusBar().showMessage("connecting...")
            
        elif state in [proxy_pb2.ERROR, proxy_pb2.DISCONNECTED]:
            self.view.active_node_id = None
            self.view.reset_tree_styles()
            
            if state == proxy_pb2.ERROR:
                self.view.change_button("error")
                self.view.statusBar().showMessage("error")
                self.view.show_notification(f"Error: {msg}", is_error=True)
            else:
                self.view.change_button("stopped")
                self.view.statusBar().showMessage("stopped")

    def _on_connection_lost(self):
        if not self.is_initialized:
            return
        self.view.change_button("error")
        self.view.statusBar().showMessage("unavailable")
        self.view.show_notification("Service unavailable", is_error=True)