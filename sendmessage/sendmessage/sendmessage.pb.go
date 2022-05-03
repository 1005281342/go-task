// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.14.0
// source: sendmessage.proto

package sendmessage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageType int32

const (
	MessageType_text     MessageType = 0
	MessageType_markdown MessageType = 1
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "text",
		1: "markdown",
	}
	MessageType_value = map[string]int32{
		"text":     0,
		"markdown": 1,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_sendmessage_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_sendmessage_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_sendmessage_proto_rawDescGZIP(), []int{0}
}

type AccountType int32

const (
	AccountType_OA AccountType = 0
)

// Enum value maps for AccountType.
var (
	AccountType_name = map[int32]string{
		0: "OA",
	}
	AccountType_value = map[string]int32{
		"OA": 0,
	}
)

func (x AccountType) Enum() *AccountType {
	p := new(AccountType)
	*p = x
	return p
}

func (x AccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_sendmessage_proto_enumTypes[1].Descriptor()
}

func (AccountType) Type() protoreflect.EnumType {
	return &file_sendmessage_proto_enumTypes[1]
}

func (x AccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountType.Descriptor instead.
func (AccountType) EnumDescriptor() ([]byte, []int) {
	return file_sendmessage_proto_rawDescGZIP(), []int{1}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Type    MessageType `protobuf:"varint,2,opt,name=type,proto3,enum=sendmessage.MessageType" json:"type,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sendmessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_sendmessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_sendmessage_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Message) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_text
}

type Sender struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type AccountType `protobuf:"varint,2,opt,name=type,proto3,enum=sendmessage.AccountType" json:"type,omitempty"`
}

func (x *Sender) Reset() {
	*x = Sender{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sendmessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sender) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sender) ProtoMessage() {}

func (x *Sender) ProtoReflect() protoreflect.Message {
	mi := &file_sendmessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sender.ProtoReflect.Descriptor instead.
func (*Sender) Descriptor() ([]byte, []int) {
	return file_sendmessage_proto_rawDescGZIP(), []int{1}
}

func (x *Sender) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sender) GetType() AccountType {
	if x != nil {
		return x.Type
	}
	return AccountType_OA
}

type Receiver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type AccountType `protobuf:"varint,2,opt,name=type,proto3,enum=sendmessage.AccountType" json:"type,omitempty"`
}

func (x *Receiver) Reset() {
	*x = Receiver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sendmessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receiver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receiver) ProtoMessage() {}

func (x *Receiver) ProtoReflect() protoreflect.Message {
	mi := &file_sendmessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receiver.ProtoReflect.Descriptor instead.
func (*Receiver) Descriptor() ([]byte, []int) {
	return file_sendmessage_proto_rawDescGZIP(), []int{2}
}

func (x *Receiver) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Receiver) GetType() AccountType {
	if x != nil {
		return x.Type
	}
	return AccountType_OA
}

type SendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg      *Message  `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Sender   *Sender   `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver *Receiver `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (x *SendReq) Reset() {
	*x = SendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sendmessage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReq) ProtoMessage() {}

func (x *SendReq) ProtoReflect() protoreflect.Message {
	mi := &file_sendmessage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReq.ProtoReflect.Descriptor instead.
func (*SendReq) Descriptor() ([]byte, []int) {
	return file_sendmessage_proto_rawDescGZIP(), []int{3}
}

func (x *SendReq) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *SendReq) GetSender() *Sender {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *SendReq) GetReceiver() *Receiver {
	if x != nil {
		return x.Receiver
	}
	return nil
}

type SendRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendRsp) Reset() {
	*x = SendRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sendmessage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRsp) ProtoMessage() {}

func (x *SendRsp) ProtoReflect() protoreflect.Message {
	mi := &file_sendmessage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRsp.ProtoReflect.Descriptor instead.
func (*SendRsp) Descriptor() ([]byte, []int) {
	return file_sendmessage_proto_rawDescGZIP(), []int{4}
}

var File_sendmessage_proto protoreflect.FileDescriptor

var file_sendmessage_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x51, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x46, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x65,
	0x6e, 0x64, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x48, 0x0a, 0x08, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x26, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x6e, 0x64,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x22, 0x09, 0x0a, 0x07, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x73, 0x70, 0x2a, 0x25, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x10, 0x01, 0x2a, 0x15, 0x0a, 0x0b, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x41,
	0x10, 0x00, 0x32, 0x39, 0x0a, 0x03, 0x72, 0x70, 0x63, 0x12, 0x32, 0x0a, 0x04, 0x73, 0x65, 0x6e,
	0x64, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x73, 0x70, 0x42, 0x0e, 0x5a,
	0x0c, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sendmessage_proto_rawDescOnce sync.Once
	file_sendmessage_proto_rawDescData = file_sendmessage_proto_rawDesc
)

func file_sendmessage_proto_rawDescGZIP() []byte {
	file_sendmessage_proto_rawDescOnce.Do(func() {
		file_sendmessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_sendmessage_proto_rawDescData)
	})
	return file_sendmessage_proto_rawDescData
}

var file_sendmessage_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_sendmessage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sendmessage_proto_goTypes = []interface{}{
	(MessageType)(0), // 0: sendmessage.MessageType
	(AccountType)(0), // 1: sendmessage.AccountType
	(*Message)(nil),  // 2: sendmessage.Message
	(*Sender)(nil),   // 3: sendmessage.Sender
	(*Receiver)(nil), // 4: sendmessage.Receiver
	(*SendReq)(nil),  // 5: sendmessage.SendReq
	(*SendRsp)(nil),  // 6: sendmessage.SendRsp
}
var file_sendmessage_proto_depIdxs = []int32{
	0, // 0: sendmessage.Message.type:type_name -> sendmessage.MessageType
	1, // 1: sendmessage.Sender.type:type_name -> sendmessage.AccountType
	1, // 2: sendmessage.Receiver.type:type_name -> sendmessage.AccountType
	2, // 3: sendmessage.SendReq.msg:type_name -> sendmessage.Message
	3, // 4: sendmessage.SendReq.sender:type_name -> sendmessage.Sender
	4, // 5: sendmessage.SendReq.receiver:type_name -> sendmessage.Receiver
	5, // 6: sendmessage.rpc.send:input_type -> sendmessage.SendReq
	6, // 7: sendmessage.rpc.send:output_type -> sendmessage.SendRsp
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_sendmessage_proto_init() }
func file_sendmessage_proto_init() {
	if File_sendmessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sendmessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sendmessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sender); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sendmessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receiver); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sendmessage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sendmessage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sendmessage_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sendmessage_proto_goTypes,
		DependencyIndexes: file_sendmessage_proto_depIdxs,
		EnumInfos:         file_sendmessage_proto_enumTypes,
		MessageInfos:      file_sendmessage_proto_msgTypes,
	}.Build()
	File_sendmessage_proto = out.File
	file_sendmessage_proto_rawDesc = nil
	file_sendmessage_proto_goTypes = nil
	file_sendmessage_proto_depIdxs = nil
}
