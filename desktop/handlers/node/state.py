from handlers.grpc import GrpcHandler
from model.worker import GrpcWorker, stub
from model.generated import proxy_pb2

class GetFullStateHandler(GrpcHandler):
    def __init__(self, view):
        super().__init__(view)
    def handle_get_state(self):
        if self.worker and self.worker.isRunning():
            return
        self.worker = GrpcWorker(stub.GetFullState, proxy_pb2.Null())
        self.worker.success.connect(self.get_state_success) 
        self.worker.finished.connect(self.worker.deleteLater)
        self.worker.start()
    def get_state_success(self, response):
        manual_map = {m.id: m for m in response.manual}
        sub_map = {s.id: s for s in response.subscription}

        for item in response.order:
            item_id = item.id
            
            if item_id in manual_map:
                self.view.add_node(manual_map[item_id])
            elif item_id in sub_map:
                self.view.add_sub(sub_map[item_id])
