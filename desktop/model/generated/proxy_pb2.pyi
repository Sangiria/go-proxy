from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Node(_message.Message):
    __slots__ = ("id", "type", "name", "address", "port", "transport", "tls")
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    PORT_FIELD_NUMBER: _ClassVar[int]
    TRANSPORT_FIELD_NUMBER: _ClassVar[int]
    TLS_FIELD_NUMBER: _ClassVar[int]
    id: str
    type: str
    name: str
    address: str
    port: int
    transport: str
    tls: str
    def __init__(self, id: _Optional[str] = ..., type: _Optional[str] = ..., name: _Optional[str] = ..., address: _Optional[str] = ..., port: _Optional[int] = ..., transport: _Optional[str] = ..., tls: _Optional[str] = ...) -> None: ...

class Subscription(_message.Message):
    __slots__ = ("id", "name", "nodes")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    NODES_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    nodes: _containers.RepeatedCompositeFieldContainer[Node]
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., nodes: _Optional[_Iterable[_Union[Node, _Mapping]]] = ...) -> None: ...

class Url(_message.Message):
    __slots__ = ("url",)
    URL_FIELD_NUMBER: _ClassVar[int]
    url: str
    def __init__(self, url: _Optional[str] = ...) -> None: ...
