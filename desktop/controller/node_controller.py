from model.worker import GrpcWorker, stub
from model.generated import proxy_pb2
import grpc

class AddRequestController:
    def __init__(self, view):
        self.view = view
        self.dialog = None
        self.manual, self.subs = dict(), dict()
    def handle_add(self):
        from view.window import AddDialog
        self.dialog = AddDialog(self.view)

        def process_add():
            self.dialog.labelError.setStyleSheet("color: black;")
            self.dialog.labelError.setText("adding...")
            url = self.dialog.lineEdit.text().strip()
            if not url:
                self.dialog.labelError.setText("field cannot be empty")
                return

            self.dialog.setEnabled(False)
            if url.startswith(("http", "https")):
                self.worker = GrpcWorker(stub.AddSubscription, proxy_pb2.Url(url=url))
            else:
                self.worker = GrpcWorker(stub.AddManual, proxy_pb2.Url(url=url))

            self.worker.success.connect(self.add_success) 
            self.worker.error.connect(self.add_error) 
            self.worker.start()

        self.dialog.buttonBox.accepted.disconnect()
        self.dialog.buttonBox.accepted.connect(process_add)
        self.dialog.exec()
    def add_success(self, response):
        self.dialog.accept()
        if isinstance(response, proxy_pb2.Node):
            self.manual[response.id] = response.name
            self.view.add_node(response)
            print("success adding node")
        elif isinstance(response, proxy_pb2.Subscription):
            self.subs[response.id] = {
                "name": response.name,
                "nodes": {node.id: node.name for node in response.nodes}
            }
            self.view.add_sub(response)
            print("success adding sub")
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
                message = f"invalid url, operation cancelled"
            elif code == grpc.StatusCode.CANCELLED and self.worker.method == stub.AddSubscription:
                message = f"failed to get the subscription"
            elif code == grpc.StatusCode.INTERNAL:
                message = f"internal error, operation cancelled"
            else:
                message = f"error [{code}]: {details}"
        else:
            message = f"unknown error: {str(err)}"

        self.dialog.labelError.setText(message)