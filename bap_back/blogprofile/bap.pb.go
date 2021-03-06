// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: blogprofile/bap.proto

package blogprofile

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

type BlogId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BlogId) Reset() {
	*x = BlogId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogprofile_bap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogId) ProtoMessage() {}

func (x *BlogId) ProtoReflect() protoreflect.Message {
	mi := &file_blogprofile_bap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogId.ProtoReflect.Descriptor instead.
func (*BlogId) Descriptor() ([]byte, []int) {
	return file_blogprofile_bap_proto_rawDescGZIP(), []int{0}
}

func (x *BlogId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NoId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoId) Reset() {
	*x = NoId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogprofile_bap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoId) ProtoMessage() {}

func (x *NoId) ProtoReflect() protoreflect.Message {
	mi := &file_blogprofile_bap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoId.ProtoReflect.Descriptor instead.
func (*NoId) Descriptor() ([]byte, []int) {
	return file_blogprofile_bap_proto_rawDescGZIP(), []int{1}
}

type BlogDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Article string   `protobuf:"bytes,2,opt,name=article,proto3" json:"article,omitempty"`
	Open    bool     `protobuf:"varint,3,opt,name=open,proto3" json:"open,omitempty"`
	Tag     []string `protobuf:"bytes,4,rep,name=tag,proto3" json:"tag,omitempty"`
	Title   string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Date    string   `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *BlogDetail) Reset() {
	*x = BlogDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogprofile_bap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogDetail) ProtoMessage() {}

func (x *BlogDetail) ProtoReflect() protoreflect.Message {
	mi := &file_blogprofile_bap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogDetail.ProtoReflect.Descriptor instead.
func (*BlogDetail) Descriptor() ([]byte, []int) {
	return file_blogprofile_bap_proto_rawDescGZIP(), []int{2}
}

func (x *BlogDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BlogDetail) GetArticle() string {
	if x != nil {
		return x.Article
	}
	return ""
}

func (x *BlogDetail) GetOpen() bool {
	if x != nil {
		return x.Open
	}
	return false
}

func (x *BlogDetail) GetTag() []string {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *BlogDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlogDetail) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type BlogList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Blogs []*BlogDetail `protobuf:"bytes,2,rep,name=blogs,proto3" json:"blogs,omitempty"`
	Tags  []string      `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *BlogList) Reset() {
	*x = BlogList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogprofile_bap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogList) ProtoMessage() {}

func (x *BlogList) ProtoReflect() protoreflect.Message {
	mi := &file_blogprofile_bap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogList.ProtoReflect.Descriptor instead.
func (*BlogList) Descriptor() ([]byte, []int) {
	return file_blogprofile_bap_proto_rawDescGZIP(), []int{3}
}

func (x *BlogList) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *BlogList) GetBlogs() []*BlogDetail {
	if x != nil {
		return x.Blogs
	}
	return nil
}

func (x *BlogList) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type ProfileDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Date    string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ProfileDetail) Reset() {
	*x = ProfileDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogprofile_bap_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileDetail) ProtoMessage() {}

func (x *ProfileDetail) ProtoReflect() protoreflect.Message {
	mi := &file_blogprofile_bap_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileDetail.ProtoReflect.Descriptor instead.
func (*ProfileDetail) Descriptor() ([]byte, []int) {
	return file_blogprofile_bap_proto_rawDescGZIP(), []int{4}
}

func (x *ProfileDetail) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ProfileDetail) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

var File_blogprofile_bap_proto protoreflect.FileDescriptor

var file_blogprofile_bap_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x62, 0x61,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x06,
	0x0a, 0x04, 0x4e, 0x6f, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x67, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6f,
	0x70, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22,
	0x63, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x2d, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x42,
	0x6c, 0x6f, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x67, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x22, 0x3d, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x32, 0xae, 0x01, 0x0a, 0x03, 0x42, 0x61, 0x70, 0x12, 0x36, 0x0a, 0x04, 0x42,
	0x6c, 0x6f, 0x67, 0x12, 0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x1a, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x11, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4e, 0x6f, 0x49, 0x64, 0x1a,
	0x15, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x42, 0x6c,
	0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x4e, 0x6f, 0x49, 0x64, 0x1a, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x70, 0x2f, 0x62, 0x61, 0x70, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blogprofile_bap_proto_rawDescOnce sync.Once
	file_blogprofile_bap_proto_rawDescData = file_blogprofile_bap_proto_rawDesc
)

func file_blogprofile_bap_proto_rawDescGZIP() []byte {
	file_blogprofile_bap_proto_rawDescOnce.Do(func() {
		file_blogprofile_bap_proto_rawDescData = protoimpl.X.CompressGZIP(file_blogprofile_bap_proto_rawDescData)
	})
	return file_blogprofile_bap_proto_rawDescData
}

var file_blogprofile_bap_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_blogprofile_bap_proto_goTypes = []interface{}{
	(*BlogId)(nil),        // 0: blogprofile.BlogId
	(*NoId)(nil),          // 1: blogprofile.NoId
	(*BlogDetail)(nil),    // 2: blogprofile.BlogDetail
	(*BlogList)(nil),      // 3: blogprofile.BlogList
	(*ProfileDetail)(nil), // 4: blogprofile.ProfileDetail
}
var file_blogprofile_bap_proto_depIdxs = []int32{
	2, // 0: blogprofile.BlogList.blogs:type_name -> blogprofile.BlogDetail
	0, // 1: blogprofile.Bap.Blog:input_type -> blogprofile.BlogId
	1, // 2: blogprofile.Bap.Blogs:input_type -> blogprofile.NoId
	1, // 3: blogprofile.Bap.Profile:input_type -> blogprofile.NoId
	2, // 4: blogprofile.Bap.Blog:output_type -> blogprofile.BlogDetail
	3, // 5: blogprofile.Bap.Blogs:output_type -> blogprofile.BlogList
	4, // 6: blogprofile.Bap.Profile:output_type -> blogprofile.ProfileDetail
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blogprofile_bap_proto_init() }
func file_blogprofile_bap_proto_init() {
	if File_blogprofile_bap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blogprofile_bap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogId); i {
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
		file_blogprofile_bap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoId); i {
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
		file_blogprofile_bap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogDetail); i {
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
		file_blogprofile_bap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogList); i {
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
		file_blogprofile_bap_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileDetail); i {
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
			RawDescriptor: file_blogprofile_bap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blogprofile_bap_proto_goTypes,
		DependencyIndexes: file_blogprofile_bap_proto_depIdxs,
		MessageInfos:      file_blogprofile_bap_proto_msgTypes,
	}.Build()
	File_blogprofile_bap_proto = out.File
	file_blogprofile_bap_proto_rawDesc = nil
	file_blogprofile_bap_proto_goTypes = nil
	file_blogprofile_bap_proto_depIdxs = nil
}
