from PySide6.QtCore import Qt
from handlers.grpc import GrpcHandler
from model.worker import GrpcWorker, stub
from model.generated import proxy_pb2
import grpc

class UpdateHandler(GrpcHandler):
    def __init__(self, view):
        super().__init__(view)
    def handle_update(self, item):
        self.item = item
        item_id = item.data(0, Qt.UserRole)

        self.view.statusBar().showMessage("updating subscription...")

        self.worker = GrpcWorker(stub.UpdateSubscription, proxy_pb2.Id(id=str(item_id)))
        self.worker.success.connect(self.update_success)
        self.worker.error.connect(self.update_error)
        self.worker.start()
    def update_success(self, response):
        self.view.statusBar().clearMessage()
        self.view.update_sub_nodes(self.item, response.nodes)
        self.view.show_notification("Subscription updated successfully!")
    def update_error(self, err):
        self.view.statusBar().clearMessage()
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