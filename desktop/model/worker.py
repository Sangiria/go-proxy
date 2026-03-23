from PySide6.QtCore import QThread, Signal
from model.generated import proxy_pb2_grpc, proxy_pb2
import grpc

channel = grpc.insecure_channel('localhost:3333')

node_stub = proxy_pb2_grpc.NodeServiceStub(channel)
proxy_stub = proxy_pb2_grpc.ProxyServiceStub(channel)

class StatusListener(QThread):
    status_updated = Signal(object)
    connection_lost = Signal(str)

    def __init__(self):
        super().__init__()
        self._is_running = True

    def run(self):
        try:
            request = proxy_pb2.Null() 
            self.responses = proxy_stub.SubscribeStatus(request)
            
            for response in self.responses:
                if not self._is_running:
                    break
                self.status_updated.emit(response)
        
        except grpc.RpcError as e:
            if self._is_running:
                self.connection_lost.emit(f"{e.code()}")
        except Exception as e:
            if self._is_running:
                self.connection_lost.emit(str(e))

    def stop(self):
        self._is_running = False
        if hasattr(self, 'responses'):
            self.responses.cancel()
        self.quit()
        self.wait()

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