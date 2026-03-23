from PySide6.QtCore import Qt
from handlers.grpc import GrpcHandler
from model.worker import GrpcWorker, node_stub
from model.generated import proxy_pb2
from handlers.node.edit import EditHandler

class GetHandler(GrpcHandler):
    def __init__(self, view):
        super().__init__(view)
        self.edit_handler = None
    def handle_get(self, item):
        item_id = item.data(0, Qt.UserRole)
        item_role = item.data(0, Qt.UserRole + 1)

        self.edit_handler = EditHandler(self.view)

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

            self.edit_handler.show_edit_dialog(response, item_role)

        if item_role == "node":
            self.worker = GrpcWorker(node_stub.GetNode, proxy_pb2.Id(**id_args))
        else:
            self.worker = GrpcWorker(node_stub.GetSubscription, proxy_pb2.Id(id=str(item_id)))
        
        self.worker.success.connect(on_success)
        self.worker.start()