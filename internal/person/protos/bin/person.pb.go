// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: internal/person/protos/person.proto

package bin

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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_person_protos_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_internal_person_protos_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_internal_person_protos_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Person) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Person) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Person) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreatePersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreatePersonResponse) Reset() {
	*x = CreatePersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_person_protos_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonResponse) ProtoMessage() {}

func (x *CreatePersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_person_protos_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonResponse.ProtoReflect.Descriptor instead.
func (*CreatePersonResponse) Descriptor() ([]byte, []int) {
	return file_internal_person_protos_person_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePersonResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePersonResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type GetPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPersonRequest) Reset() {
	*x = GetPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_person_protos_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonRequest) ProtoMessage() {}

func (x *GetPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_person_protos_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonRequest.ProtoReflect.Descriptor instead.
func (*GetPersonRequest) Descriptor() ([]byte, []int) {
	return file_internal_person_protos_person_proto_rawDescGZIP(), []int{2}
}

func (x *GetPersonRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *GetPersonResponse) Reset() {
	*x = GetPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_person_protos_person_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonResponse) ProtoMessage() {}

func (x *GetPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_person_protos_person_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonResponse.ProtoReflect.Descriptor instead.
func (*GetPersonResponse) Descriptor() ([]byte, []int) {
	return file_internal_person_protos_person_proto_rawDescGZIP(), []int{3}
}

func (x *GetPersonResponse) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type DeletePersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePersonRequest) Reset() {
	*x = DeletePersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_person_protos_person_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePersonRequest) ProtoMessage() {}

func (x *DeletePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_person_protos_person_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePersonRequest.ProtoReflect.Descriptor instead.
func (*DeletePersonRequest) Descriptor() ([]byte, []int) {
	return file_internal_person_protos_person_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePersonRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeletePersonResponse) Reset() {
	*x = DeletePersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_person_protos_person_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePersonResponse) ProtoMessage() {}

func (x *DeletePersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_person_protos_person_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePersonResponse.ProtoReflect.Descriptor instead.
func (*DeletePersonResponse) Descriptor() ([]byte, []int) {
	return file_internal_person_protos_person_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePersonResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeletePersonResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_internal_person_protos_person_proto protoreflect.FileDescriptor

var file_internal_person_protos_person_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x68, 0x0a,
	0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0xce, 0x01, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x1a,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x62, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_person_protos_person_proto_rawDescOnce sync.Once
	file_internal_person_protos_person_proto_rawDescData = file_internal_person_protos_person_proto_rawDesc
)

func file_internal_person_protos_person_proto_rawDescGZIP() []byte {
	file_internal_person_protos_person_proto_rawDescOnce.Do(func() {
		file_internal_person_protos_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_person_protos_person_proto_rawDescData)
	})
	return file_internal_person_protos_person_proto_rawDescData
}

var file_internal_person_protos_person_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_person_protos_person_proto_goTypes = []interface{}{
	(*Person)(nil),               // 0: protos.Person
	(*CreatePersonResponse)(nil), // 1: protos.CreatePersonResponse
	(*GetPersonRequest)(nil),     // 2: protos.GetPersonRequest
	(*GetPersonResponse)(nil),    // 3: protos.GetPersonResponse
	(*DeletePersonRequest)(nil),  // 4: protos.DeletePersonRequest
	(*DeletePersonResponse)(nil), // 5: protos.DeletePersonResponse
}
var file_internal_person_protos_person_proto_depIdxs = []int32{
	0, // 0: protos.GetPersonResponse.person:type_name -> protos.Person
	0, // 1: protos.PersonService.Create:input_type -> protos.Person
	2, // 2: protos.PersonService.Get:input_type -> protos.GetPersonRequest
	4, // 3: protos.PersonService.Delete:input_type -> protos.DeletePersonRequest
	1, // 4: protos.PersonService.Create:output_type -> protos.CreatePersonResponse
	3, // 5: protos.PersonService.Get:output_type -> protos.GetPersonResponse
	5, // 6: protos.PersonService.Delete:output_type -> protos.DeletePersonResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_person_protos_person_proto_init() }
func file_internal_person_protos_person_proto_init() {
	if File_internal_person_protos_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_person_protos_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_internal_person_protos_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonResponse); i {
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
		file_internal_person_protos_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonRequest); i {
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
		file_internal_person_protos_person_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonResponse); i {
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
		file_internal_person_protos_person_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePersonRequest); i {
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
		file_internal_person_protos_person_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePersonResponse); i {
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
			RawDescriptor: file_internal_person_protos_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_person_protos_person_proto_goTypes,
		DependencyIndexes: file_internal_person_protos_person_proto_depIdxs,
		MessageInfos:      file_internal_person_protos_person_proto_msgTypes,
	}.Build()
	File_internal_person_protos_person_proto = out.File
	file_internal_person_protos_person_proto_rawDesc = nil
	file_internal_person_protos_person_proto_goTypes = nil
	file_internal_person_protos_person_proto_depIdxs = nil
}
