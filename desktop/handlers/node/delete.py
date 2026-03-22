from PySide6.QtCore import Qt
from handlers.grpc import GrpcHandler
from model.worker import GrpcWorker, stub
from model.generated import proxy_pb2

class DeleteHandler(GrpcHandler):
    def __init__(self, view):
        super().__init__(view)
    def handle_delete(self, item):
        item_id = item.data(0, Qt.UserRole)
        item_role = item.data(0, Qt.UserRole + 1)

        id_args = {"id": str(item_id)}

        parent_item = item.parent()
        source_id = None
        if parent_item:
            source_id = parent_item.data(0, Qt.UserRole)
            if source_id:
                id_args["source_id"] = str(source_id)

        if item_role == "node":
            self.worker = GrpcWorker(stub.DeleteNode, proxy_pb2.Id(**id_args))
        else:
            self.worker = GrpcWorker(stub.DeleteSubscription, proxy_pb2.Id(id=str(item_id)))
        
        self.worker.success.connect(lambda: self.delete_success(item))
        self.worker.start()
    def delete_success(self, item):
        self.view.remove_item(item)
