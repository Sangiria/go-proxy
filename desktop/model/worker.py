from PySide6.QtCore import QThread, Signal
from model.generated import proxy_pb2_grpc
import grpc

channel = grpc.insecure_channel('localhost:3333')
stub = proxy_pb2_grpc.NodeServiceStub(channel)

class GrpcWorker(QThread):
    success = Signal(object)
    error = Signal(object)

    def __init__(self, method, *args):
        super().__init__()
        self.method = method
        self.args = args

    def run(self):
        try:
            response = self.method(*self.args)
            self.success.emit(response)
        except grpc.RpcError as e:
            self.error.emit(e) 
        except Exception as e:
            self.error.emit(e)