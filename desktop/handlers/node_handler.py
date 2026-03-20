from PySide6.QtCore import Qt
from PySide6.QtWidgets import QDialog
from model.worker import GrpcWorker, stub
from model.generated import proxy_pb2
import grpc

class GrpcHandler:
    def __init__(self, view):
        self.view = view
        self.worker = None

class GetFullStateHandler(GrpcHandler):
    def __init__(self, view):
        super().__init__(view)
    def handle_get_state(self):
        self.worker = GrpcWorker(stub.GetFullState, proxy_pb2.Null())
        self.worker.success.connect(self.get_state_success) 
        self.worker.start()
    def get_state_success(self, response):
        for m in response.manual:
            self.view.add_node(m)
        for s in response.subscription:
            self.view.add_sub(s)
class AddHandler(GrpcHandler):
    def __init__(self, view):
        super().__init__(view)
        self.dialog = None
    def handle_add(self):
        from view.dialog import AddDialog
        self.dialog = AddDialog(self.view)
        self.dialog.buttonBox.accepted.disconnect()
        self.dialog.buttonBox.accepted.connect(self.process_add)
        self.dialog.exec()
    def process_add(self):
        url = self.dialog.lineEdit.text().strip()
        if not url:
            self.dialog.labelError.setStyleSheet("color: red;")
            self.dialog.labelError.setText("field cannot be empty")
            return

        self.dialog.labelError.setStyleSheet("color: black;")
        self.dialog.labelError.setText("adding...")

        self.dialog.setEnabled(False)
        if url.startswith(("http", "https")):
            self.worker = GrpcWorker(stub.AddSubscription, proxy_pb2.Url(url=url))
        else:
            self.worker = GrpcWorker(stub.AddNode, proxy_pb2.Url(url=url))

        self.worker.success.connect(self.add_success) 
        self.worker.error.connect(self.add_error)
        self.worker.start()
    def add_success(self, response):
        self.dialog.accept()
        if isinstance(response, proxy_pb2.Node):
            self.view.add_node(response)
        elif isinstance(response, proxy_pb2.Subscription):
            self.view.add_sub(response)
    def add_error(self, err):
        self.dialog.setEnabled(True)
        self.dialog.labelError.setStyleSheet("color: red;")

        if isinstance(err, grpc.RpcError):
            code = err.code()
            details = err.details()

            if code == grpc.StatusCode.ALREADY_EXISTS:
                message = f"({details})"
            elif code == grpc.StatusCode.UNAVAILABLE:
                message = "server is unavailable, please check the connection"
            elif code == grpc.StatusCode.INVALID_ARGUMENT:
                message = f"invalid url, {details}"
            elif code == grpc.StatusCode.CANCELLED and self.worker.method == stub.AddSubscription:
                message = f"failed to get the subscription"
            elif code == grpc.StatusCode.INTERNAL:
                message = f"internal error, operation cancelled"
            else:
                message = f"error [{code}]: {details}"
        else:
            message = f"unknown error: {str(err)}"

        self.dialog.labelError.setText(message)

class GetHandler(GrpcHandler):
    def __init__(self, view):
        super().__init__(view)
    def handle_get(self, item):
        item_id = item.data(0, Qt.UserRole)
        item_role = item.data(0, Qt.UserRole + 1)
        edit_handler = EditHandler(self.view)

        id_args = {"id": str(item_id)}
        
        parent_item = item.parent()
        source_id = None
        if parent_item:
            source_id = parent_item.data(0, Qt.UserRole)
            if source_id:
                id_args["source_id"] = str(source_id)

        def on_success(response):
            if not response.id:
                response.id = str(item_id)
            if source_id:
                response.source_id = str(source_id)

            edit_handler.show_edit_dialog(response, item_role)

        if item_role == "node":
            self.worker = GrpcWorker(stub.GetNode, proxy_pb2.Id(**id_args))
        else:
            self.worker = GrpcWorker(stub.GetSubscription, proxy_pb2.Id(id=str(item_id)))
        
        self.worker.success.connect(on_success)
        self.worker.start()
        
class EditHandler(GrpcHandler):
    def __init__(self, view):
        super().__init__(view)
        self.dialog = None
    def show_edit_dialog(self, data, role):
        if role == "node":
            from view.dialog import EditNodeDialog
            self.dialog = EditNodeDialog(data, self.view)
        else:
            from view.dialog import EditSubscribtionDialog
            self.dialog = EditSubscribtionDialog(data, self.view)

        self.dialog.start_edit = lambda: self.process_edit(role)
        self.dialog.exec()
    def process_edit(self, role):
        delta = getattr(self.dialog, 'delta', None)
        full = getattr(self.dialog, 'full', None)

        if delta:
            self.dialog.setEnabled(False)
            method = stub.EditNode if role == "node" else stub.EditSubscription
            self.handle_edit(method, delta, full)
    def handle_edit(self, method, delta, full):
        self.worker = GrpcWorker(method, delta)
        self.worker.success.connect(lambda: self.edit_success(full))
        self.worker.error.connect(self.edit_error)
        self.worker.start()
    def edit_success(self, data):
        if self.dialog:
            super(type(self.dialog), self.dialog).accept()
        
        self.view.update_item(data)
    def edit_error(self, err):
        self.dialog.setEnabled(True)

        if isinstance(err, grpc.RpcError):
            code = err.code()
            details = err.details()

            if code == grpc.StatusCode.ALREADY_EXISTS:
                message = f"({details})"
            elif code == grpc.StatusCode.UNAVAILABLE:
                message = "server is unavailable, please check the connection"
            elif code == grpc.StatusCode.INVALID_ARGUMENT:
                message = f"{details}"
            elif code == grpc.StatusCode.INTERNAL:
                message = f"internal error, operation cancelled"
            else:
                message = f"error [{code}]: {details}"
        else:
            message = f"unknown error: {str(err)}"

        self.dialog.labelError.setText(message)


    
    
        
    