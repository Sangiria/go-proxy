from handlers.grpc import GrpcHandler
from model.worker import GrpcWorker, stub
from model.generated import proxy_pb2
import grpc

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

        if url.startswith(("http", "https")):
            self.worker = GrpcWorker(stub.AddSubscription, proxy_pb2.Url(url=url))
        else:
            self.worker = GrpcWorker(stub.AddNode, proxy_pb2.Url(url=url))

        self.dialog.setEnabled(False)
        self.worker.success.connect(self.add_success) 
        self.worker.error.connect(self.add_error)
        self.worker.finished.connect(self.worker.deleteLater)
        self.worker.start()
    def add_success(self, response):
        if isinstance(response, proxy_pb2.Node):
            self.view.add_node(response)
        elif isinstance(response, proxy_pb2.Subscription):
            self.view.add_sub(response)   
        self.dialog.accept()
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