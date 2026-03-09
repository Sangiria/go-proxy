#include "client.h"

const string address = "localhost:3333";

NodeClient::NodeClient()
{
    auto channel = grpc::CreateChannel(
        address,
        grpc::InsecureChannelCredentials()
    );
    stub_ = NodeService::NewStub(channel);
}