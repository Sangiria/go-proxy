from model.worker import GrpcWorker, stub
from model.generated import proxy_pb2
import grpc

manual, subs = dict(), dict()

class NodeController:
    def __init__(self, view):
        self.view = view
        self.worker = None

class GetStateController(NodeController):
    def __init__(self, view):
        super().__init__(view)
    def handle_get_state(self):
        self.worker = GrpcWorker(stub.GetFullState, proxy_pb2.Null())
        self.worker.success.connect(self.get_state_success) 
        self.worker.start()
    def get_state_success(self, response):
        for m in response.manual:
            manual[m.id] = m.name
            self.view.add_node(m)

        for s in response.subscription:
            subs[s.id] = {
                "name": s.name,
                "nodes": {n.id: n.name for n in s.nodes}
            }
            self.view.add_sub(s)

class AddController(NodeController):
    def __init__(self, view):
        super().__init__(view)
        self.dialog = None
    def handle_add(self):
        from view.window import AddDialog
        self.dialog = AddDialog(self.view)

        def process_add():
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

        self.dialog.buttonBox.accepted.disconnect()
        self.dialog.buttonBox.accepted.connect(process_add)
        self.dialog.exec()
    def add_success(self, response):
        self.dialog.accept()
        if isinstance(response, proxy_pb2.Node):
            manual[response.id] = response.name
            self.view.add_node(response)
        elif isinstance(response, proxy_pb2.Subscription):
            subs[response.id] = {
                "name": response.name,
                "nodes": {node.id: node.name for node in response.nodes}
            }
            self.view.add_sub(response)
    def add_error(self, err):
        self.dialog.setEnabled(True)
        self.dialog.labelError.setStyleSheet("color: red;")

        if isinstance(err, grpc.Call):
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