// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v5.28.0
// source: account/profile/profile.proto

package profile

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Photo string          `protobuf:"bytes,2,opt,name=photo,proto3" json:"photo,omitempty"`
	T     *structpb.Value `protobuf:"bytes,3,opt,name=T,proto3" json:"T,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_profile_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_account_profile_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_account_profile_profile_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Profile) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *Profile) GetT() *structpb.Value {
	if x != nil {
		return x.T
	}
	return nil
}

type Profiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Profile `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Profiles) Reset() {
	*x = Profiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_profile_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profiles) ProtoMessage() {}

func (x *Profiles) ProtoReflect() protoreflect.Message {
	mi := &file_account_profile_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profiles.ProtoReflect.Descriptor instead.
func (*Profiles) Descriptor() ([]byte, []int) {
	return file_account_profile_profile_proto_rawDescGZIP(), []int{1}
}

func (x *Profiles) GetItems() []*Profile {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_account_profile_profile_proto protoreflect.FileDescriptor

var file_account_profile_profile_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x0a, 0x01, 0x54, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x01, 0x54, 0x22, 0x32, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x62, 0x6c, 0x6f, 0x6e, 0x64, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_profile_profile_proto_rawDescOnce sync.Once
	file_account_profile_profile_proto_rawDescData = file_account_profile_profile_proto_rawDesc
)

func file_account_profile_profile_proto_rawDescGZIP() []byte {
	file_account_profile_profile_proto_rawDescOnce.Do(func() {
		file_account_profile_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_profile_profile_proto_rawDescData)
	})
	return file_account_profile_profile_proto_rawDescData
}

var file_account_profile_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_account_profile_profile_proto_goTypes = []interface{}{
	(*Profile)(nil),        // 0: profile.Profile
	(*Profiles)(nil),       // 1: profile.Profiles
	(*structpb.Value)(nil), // 2: google.protobuf.Value
}
var file_account_profile_profile_proto_depIdxs = []int32{
	2, // 0: profile.Profile.T:type_name -> google.protobuf.Value
	0, // 1: profile.Profiles.items:type_name -> profile.Profile
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_account_profile_profile_proto_init() }
func file_account_profile_profile_proto_init() {
	if File_account_profile_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_profile_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_account_profile_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profiles); i {
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
			RawDescriptor: file_account_profile_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_profile_profile_proto_goTypes,
		DependencyIndexes: file_account_profile_profile_proto_depIdxs,
		MessageInfos:      file_account_profile_profile_proto_msgTypes,
	}.Build()
	File_account_profile_profile_proto = out.File
	file_account_profile_profile_proto_rawDesc = nil
	file_account_profile_profile_proto_goTypes = nil
	file_account_profile_profile_proto_depIdxs = nil
}
