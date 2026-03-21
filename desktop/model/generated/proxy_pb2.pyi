from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class State(_message.Message):
    __slots__ = ("manual", "subscription", "order")
    MANUAL_FIELD_NUMBER: _ClassVar[int]
    SUBSCRIPTION_FIELD_NUMBER: _ClassVar[int]
    ORDER_FIELD_NUMBER: _ClassVar[int]
    manual: _containers.RepeatedCompositeFieldContainer[Node]
    subscription: _containers.RepeatedCompositeFieldContainer[Subscription]
    order: _containers.RepeatedCompositeFieldContainer[Id]
    def __init__(self, manual: _Optional[_Iterable[_Union[Node, _Mapping]]] = ..., subscription: _Optional[_Iterable[_Union[Subscription, _Mapping]]] = ..., order: _Optional[_Iterable[_Union[Id, _Mapping]]] = ...) -> None: ...

class Node(_message.Message):
    __slots__ = ("id", "type", "name", "address", "port", "transport", "security")
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    PORT_FIELD_NUMBER: _ClassVar[int]
    TRANSPORT_FIELD_NUMBER: _ClassVar[int]
    SECURITY_FIELD_NUMBER: _ClassVar[int]
    id: str
    type: str
    name: str
    address: str
    port: int
    transport: str
    security: str
    def __init__(self, id: _Optional[str] = ..., type: _Optional[str] = ..., name: _Optional[str] = ..., address: _Optional[str] = ..., port: _Optional[int] = ..., transport: _Optional[str] = ..., security: _Optional[str] = ...) -> None: ...

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

class NodeForm(_message.Message):
    __slots__ = ("id", "name", "address", "port", "uuid", "transport", "security", "sni", "fp", "pbk", "sid", "mode", "extra", "source_id")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    PORT_FIELD_NUMBER: _ClassVar[int]
    UUID_FIELD_NUMBER: _ClassVar[int]
    TRANSPORT_FIELD_NUMBER: _ClassVar[int]
    SECURITY_FIELD_NUMBER: _ClassVar[int]
    SNI_FIELD_NUMBER: _ClassVar[int]
    FP_FIELD_NUMBER: _ClassVar[int]
    PBK_FIELD_NUMBER: _ClassVar[int]
    SID_FIELD_NUMBER: _ClassVar[int]
    MODE_FIELD_NUMBER: _ClassVar[int]
    EXTRA_FIELD_NUMBER: _ClassVar[int]
    SOURCE_ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    address: str
    port: int
    uuid: str
    transport: str
    security: str
    sni: str
    fp: str
    pbk: str
    sid: str
    mode: str
    extra: str
    source_id: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., address: _Optional[str] = ..., port: _Optional[int] = ..., uuid: _Optional[str] = ..., transport: _Optional[str] = ..., security: _Optional[str] = ..., sni: _Optional[str] = ..., fp: _Optional[str] = ..., pbk: _Optional[str] = ..., sid: _Optional[str] = ..., mode: _Optional[str] = ..., extra: _Optional[str] = ..., source_id: _Optional[str] = ...) -> None: ...

class SubscriptionForm(_message.Message):
    __slots__ = ("id", "name", "url")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    url: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., url: _Optional[str] = ...) -> None: ...

class Id(_message.Message):
    __slots__ = ("id", "source_id")
    ID_FIELD_NUMBER: _ClassVar[int]
    SOURCE_ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    source_id: str
    def __init__(self, id: _Optional[str] = ..., source_id: _Optional[str] = ...) -> None: ...

class Null(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
