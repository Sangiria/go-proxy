from PySide6.QtCore import QTimer
from handlers.grpc import GrpcHandler
from model.worker import GrpcWorker, stub
import grpc

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
        if delta:
            self.dialog.setEnabled(False)
            method = stub.EditNode if role == "node" else stub.EditSubscription
            self.handle_edit(method, delta)
    def handle_edit(self, method, delta):
        self.worker = GrpcWorker(method, delta)
        self.worker.finished.connect(self.worker.deleteLater)
        self.worker.success.connect(lambda: self.edit_success(delta))
        self.worker.error.connect(self.edit_error)
        self.worker.start()
    def edit_success(self, data):
        self.view.update_item(data)
        QTimer.singleShot(0, self.dialog.accept)
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