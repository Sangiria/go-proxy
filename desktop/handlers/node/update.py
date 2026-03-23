from PySide6.QtWidgets import QProgressDialog
from PySide6.QtCore import Qt
from handlers.grpc import GrpcHandler
from model.worker import GrpcWorker, node_stub
from model.generated import proxy_pb2
import grpc

class UpdateHandler(GrpcHandler):
    def __init__(self, view):
        super().__init__(view)
        self.progress = None

    def show_progress(self):
        flags = (Qt.Window | Qt.WindowTitleHint | Qt.CustomizeWindowHint)
        flags &= ~Qt.WindowCloseButtonHint
        
        self.progress = QProgressDialog("Updating subscription...", None, 0, 0, self.view)
        self.progress.setWindowTitle("Please, wait")
        self.progress.setWindowModality(Qt.WindowModal)
        self.progress.setCancelButton(None)
        self.progress.setWindowFlags(flags)
        self.progress.reject = lambda: None 
        self.progress.show()

    def handle_update(self, item):
        self.item = item
        item_id = item.data(0, Qt.UserRole)

        self.show_progress()

        self.worker = GrpcWorker(node_stub.UpdateSubscription, proxy_pb2.Id(id=str(item_id)))
        self.worker.success.connect(self.update_success)
        self.worker.error.connect(self.update_error)
        self.worker.finished.connect(self.cleanup)
        self.worker.start()
    def update_success(self, response):
        self.view.update_sub_nodes(self.item, response.nodes)

    def cleanup(self):
        if self.progress:
            self.progress.close()

    def update_error(self, err):
        if isinstance(err, grpc.RpcError):
            code = err.code()
            details = err.details()

            if code == grpc.StatusCode.CANCELLED:
                message = f"{details}"
            elif code == grpc.StatusCode.INTERNAL:
                message = f"internal error, operation cancelled"
            else:
                message = f"error [{code}]: {details}"
        else:
            message = f"unknown error: {str(err)}"
        
        self.view.show_notification(f"{message}", is_error=True)