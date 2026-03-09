#ifndef CLIENT_H
#define CLIENT_H

#include "proxy.grpc.pb.h"
#include <grpcpp/grpcpp.h>

using namespace std;
using goproxy::NodeService;

class NodeClient {
public:
    NodeClient();

private:
    std::unique_ptr<NodeService::Stub> stub_;
};

#endif