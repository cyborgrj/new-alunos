// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: protoimp/alunos.proto

package alunos_grpc

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

type NewAluno struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Serie          string `protobuf:"bytes,2,opt,name=serie,proto3" json:"serie,omitempty"`
	Cpf            string `protobuf:"bytes,3,opt,name=cpf,proto3" json:"cpf,omitempty"`
	Email          string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Datanascimento string `protobuf:"bytes,5,opt,name=datanascimento,proto3" json:"datanascimento,omitempty"`
	Cep            string `protobuf:"bytes,6,opt,name=cep,proto3" json:"cep,omitempty"`
}

func (x *NewAluno) Reset() {
	*x = NewAluno{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoimp_alunos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAluno) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAluno) ProtoMessage() {}

func (x *NewAluno) ProtoReflect() protoreflect.Message {
	mi := &file_protoimp_alunos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAluno.ProtoReflect.Descriptor instead.
func (*NewAluno) Descriptor() ([]byte, []int) {
	return file_protoimp_alunos_proto_rawDescGZIP(), []int{0}
}

func (x *NewAluno) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewAluno) GetSerie() string {
	if x != nil {
		return x.Serie
	}
	return ""
}

func (x *NewAluno) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

func (x *NewAluno) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NewAluno) GetDatanascimento() string {
	if x != nil {
		return x.Datanascimento
	}
	return ""
}

func (x *NewAluno) GetCep() string {
	if x != nil {
		return x.Cep
	}
	return ""
}

type Aluno struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Serie          string   `protobuf:"bytes,3,opt,name=serie,proto3" json:"serie,omitempty"`
	Cpf            string   `protobuf:"bytes,4,opt,name=cpf,proto3" json:"cpf,omitempty"`
	Email          string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Datanascimento string   `protobuf:"bytes,6,opt,name=datanascimento,proto3" json:"datanascimento,omitempty"`
	Idade          int32    `protobuf:"varint,7,opt,name=idade,proto3" json:"idade,omitempty"`
	Endereco       *Address `protobuf:"bytes,8,opt,name=endereco,proto3" json:"endereco,omitempty"`
}

func (x *Aluno) Reset() {
	*x = Aluno{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoimp_alunos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Aluno) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aluno) ProtoMessage() {}

func (x *Aluno) ProtoReflect() protoreflect.Message {
	mi := &file_protoimp_alunos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aluno.ProtoReflect.Descriptor instead.
func (*Aluno) Descriptor() ([]byte, []int) {
	return file_protoimp_alunos_proto_rawDescGZIP(), []int{1}
}

func (x *Aluno) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Aluno) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Aluno) GetSerie() string {
	if x != nil {
		return x.Serie
	}
	return ""
}

func (x *Aluno) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

func (x *Aluno) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Aluno) GetDatanascimento() string {
	if x != nil {
		return x.Datanascimento
	}
	return ""
}

func (x *Aluno) GetIdade() int32 {
	if x != nil {
		return x.Idade
	}
	return 0
}

func (x *Aluno) GetEndereco() *Address {
	if x != nil {
		return x.Endereco
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cep         string `protobuf:"bytes,1,opt,name=cep,proto3" json:"cep,omitempty"`
	Logradouro  string `protobuf:"bytes,2,opt,name=logradouro,proto3" json:"logradouro,omitempty"`
	Bairro      string `protobuf:"bytes,3,opt,name=bairro,proto3" json:"bairro,omitempty"`
	Localidade  string `protobuf:"bytes,4,opt,name=localidade,proto3" json:"localidade,omitempty"`
	Uf          string `protobuf:"bytes,5,opt,name=uf,proto3" json:"uf,omitempty"`
	Complemento string `protobuf:"bytes,6,opt,name=complemento,proto3" json:"complemento,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoimp_alunos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_protoimp_alunos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_protoimp_alunos_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetCep() string {
	if x != nil {
		return x.Cep
	}
	return ""
}

func (x *Address) GetLogradouro() string {
	if x != nil {
		return x.Logradouro
	}
	return ""
}

func (x *Address) GetBairro() string {
	if x != nil {
		return x.Bairro
	}
	return ""
}

func (x *Address) GetLocalidade() string {
	if x != nil {
		return x.Localidade
	}
	return ""
}

func (x *Address) GetUf() string {
	if x != nil {
		return x.Uf
	}
	return ""
}

func (x *Address) GetComplemento() string {
	if x != nil {
		return x.Complemento
	}
	return ""
}

type AlunosList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alunos []*Aluno `protobuf:"bytes,1,rep,name=alunos,proto3" json:"alunos,omitempty"`
}

func (x *AlunosList) Reset() {
	*x = AlunosList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoimp_alunos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlunosList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlunosList) ProtoMessage() {}

func (x *AlunosList) ProtoReflect() protoreflect.Message {
	mi := &file_protoimp_alunos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlunosList.ProtoReflect.Descriptor instead.
func (*AlunosList) Descriptor() ([]byte, []int) {
	return file_protoimp_alunos_proto_rawDescGZIP(), []int{3}
}

func (x *AlunosList) GetAlunos() []*Aluno {
	if x != nil {
		return x.Alunos
	}
	return nil
}

type GetAlunosParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpf  string `protobuf:"bytes,1,opt,name=cpf,proto3" json:"cpf,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAlunosParams) Reset() {
	*x = GetAlunosParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoimp_alunos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlunosParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlunosParams) ProtoMessage() {}

func (x *GetAlunosParams) ProtoReflect() protoreflect.Message {
	mi := &file_protoimp_alunos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlunosParams.ProtoReflect.Descriptor instead.
func (*GetAlunosParams) Descriptor() ([]byte, []int) {
	return file_protoimp_alunos_proto_rawDescGZIP(), []int{4}
}

func (x *GetAlunosParams) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

func (x *GetAlunosParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_protoimp_alunos_proto protoreflect.FileDescriptor

var file_protoimp_alunos_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x70, 0x2f, 0x61, 0x6c, 0x75, 0x6e, 0x6f,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d,
	0x70, 0x22, 0x96, 0x01, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x41, 0x6c, 0x75, 0x6e, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x72, 0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x65, 0x72, 0x69, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x66, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x26, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x61, 0x73, 0x63, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x61,
	0x73, 0x63, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x65, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x65, 0x70, 0x22, 0xd6, 0x01, 0x0a, 0x05, 0x41,
	0x6c, 0x75, 0x6e, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x65, 0x72, 0x69, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x70, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x66,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x61,
	0x73, 0x63, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x64, 0x61, 0x74, 0x61, 0x6e, 0x61, 0x73, 0x63, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x64, 0x61, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69,
	0x64, 0x61, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x63, 0x6f,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d,
	0x70, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x65, 0x63, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x65,
	0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x72, 0x61, 0x64, 0x6f, 0x75, 0x72, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x72, 0x61, 0x64, 0x6f, 0x75, 0x72,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x69, 0x72, 0x72, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x61, 0x69, 0x72, 0x72, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x75, 0x66, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x75, 0x66, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0a, 0x41,
	0x6c, 0x75, 0x6e, 0x6f, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x61, 0x6c, 0x75,
	0x6e, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x69, 0x6d, 0x70, 0x2e, 0x41, 0x6c, 0x75, 0x6e, 0x6f, 0x52, 0x06, 0x61, 0x6c, 0x75, 0x6e,
	0x6f, 0x73, 0x22, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x75, 0x6e, 0x6f, 0x73, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xb4, 0x02, 0x0a, 0x0f,
	0x41, 0x6c, 0x75, 0x6e, 0x6f, 0x4d, 0x61, 0x6e, 0x65, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x34, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x75, 0x6e, 0x6f, 0x12, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x70, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x6c, 0x75,
	0x6e, 0x6f, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x70, 0x2e, 0x41, 0x6c,
	0x75, 0x6e, 0x6f, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x75, 0x6e,
	0x6f, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x70, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x75, 0x6e, 0x6f, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x70, 0x2e, 0x41, 0x6c, 0x75, 0x6e, 0x6f, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x75, 0x6e,
	0x6f, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x70, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x75, 0x6e, 0x6f, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x70, 0x2e, 0x41, 0x6c, 0x75, 0x6e, 0x6f, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x75, 0x6e, 0x6f, 0x12, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x70, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x6c, 0x75,
	0x6e, 0x6f, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x70, 0x2e, 0x41, 0x6c,
	0x75, 0x6e, 0x6f, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x6c, 0x75, 0x6e, 0x6f, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x70, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x75, 0x6e, 0x6f, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x70, 0x2e, 0x41, 0x6c, 0x75, 0x6e, 0x6f,
	0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x63, 0x79, 0x62, 0x6f, 0x72, 0x67, 0x72, 0x6a, 0x2f, 0x6e,
	0x65, 0x77, 0x2d, 0x61, 0x6c, 0x75, 0x6e, 0x6f, 0x73, 0x3b, 0x61, 0x6c, 0x75, 0x6e, 0x6f, 0x73,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoimp_alunos_proto_rawDescOnce sync.Once
	file_protoimp_alunos_proto_rawDescData = file_protoimp_alunos_proto_rawDesc
)

func file_protoimp_alunos_proto_rawDescGZIP() []byte {
	file_protoimp_alunos_proto_rawDescOnce.Do(func() {
		file_protoimp_alunos_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoimp_alunos_proto_rawDescData)
	})
	return file_protoimp_alunos_proto_rawDescData
}

var file_protoimp_alunos_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protoimp_alunos_proto_goTypes = []interface{}{
	(*NewAluno)(nil),        // 0: protoimp.NewAluno
	(*Aluno)(nil),           // 1: protoimp.Aluno
	(*Address)(nil),         // 2: protoimp.Address
	(*AlunosList)(nil),      // 3: protoimp.AlunosList
	(*GetAlunosParams)(nil), // 4: protoimp.GetAlunosParams
}
var file_protoimp_alunos_proto_depIdxs = []int32{
	2, // 0: protoimp.Aluno.endereco:type_name -> protoimp.Address
	1, // 1: protoimp.AlunosList.alunos:type_name -> protoimp.Aluno
	0, // 2: protoimp.AlunoManegement.CreateAluno:input_type -> protoimp.NewAluno
	4, // 3: protoimp.AlunoManegement.GetAlunos:input_type -> protoimp.GetAlunosParams
	4, // 4: protoimp.AlunoManegement.GetAluno:input_type -> protoimp.GetAlunosParams
	0, // 5: protoimp.AlunoManegement.UpdateAluno:input_type -> protoimp.NewAluno
	4, // 6: protoimp.AlunoManegement.DeleteAluno:input_type -> protoimp.GetAlunosParams
	1, // 7: protoimp.AlunoManegement.CreateAluno:output_type -> protoimp.Aluno
	3, // 8: protoimp.AlunoManegement.GetAlunos:output_type -> protoimp.AlunosList
	1, // 9: protoimp.AlunoManegement.GetAluno:output_type -> protoimp.Aluno
	1, // 10: protoimp.AlunoManegement.UpdateAluno:output_type -> protoimp.Aluno
	1, // 11: protoimp.AlunoManegement.DeleteAluno:output_type -> protoimp.Aluno
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protoimp_alunos_proto_init() }
func file_protoimp_alunos_proto_init() {
	if File_protoimp_alunos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoimp_alunos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAluno); i {
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
		file_protoimp_alunos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Aluno); i {
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
		file_protoimp_alunos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_protoimp_alunos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlunosList); i {
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
		file_protoimp_alunos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlunosParams); i {
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
			RawDescriptor: file_protoimp_alunos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoimp_alunos_proto_goTypes,
		DependencyIndexes: file_protoimp_alunos_proto_depIdxs,
		MessageInfos:      file_protoimp_alunos_proto_msgTypes,
	}.Build()
	File_protoimp_alunos_proto = out.File
	file_protoimp_alunos_proto_rawDesc = nil
	file_protoimp_alunos_proto_goTypes = nil
	file_protoimp_alunos_proto_depIdxs = nil
}