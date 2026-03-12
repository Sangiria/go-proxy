from PyQt6.QtCore import QThread, pyqtSignal
from generated import proxy_pb2_grpc
import grpc

channel = grpc.insecure_channel('localhost:3333')
stub = proxy_pb2_grpc.NodeServiceStub(channel)

class GrpcWorker(QThread):
    finished = pyqtSignal(object)
    error = pyqtSignal(str)

    def __init__(self, method, *args):
        super().__init__()
        self.method = method
        self.args = args

    def run(self):
        try:
            response = self.method(*self.args)
            self.finished.emit(response)
        except Exception as e:
            self.error.emit(str(e))