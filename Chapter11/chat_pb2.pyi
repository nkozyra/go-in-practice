from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Comment(_message.Message):
    __slots__ = ("username", "text", "sent")
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    SENT_FIELD_NUMBER: _ClassVar[int]
    username: str
    text: str
    sent: _timestamp_pb2.Timestamp
    def __init__(self, username: _Optional[str] = ..., text: _Optional[str] = ..., sent: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class CommentMeta(_message.Message):
    __slots__ = ("commentLength", "previousCommentCount")
    COMMENTLENGTH_FIELD_NUMBER: _ClassVar[int]
    PREVIOUSCOMMENTCOUNT_FIELD_NUMBER: _ClassVar[int]
    commentLength: int
    previousCommentCount: int
    def __init__(self, commentLength: _Optional[int] = ..., previousCommentCount: _Optional[int] = ...) -> None: ...
