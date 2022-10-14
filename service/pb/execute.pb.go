// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: execute.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname string     `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Account []*Account `protobuf:"bytes,3,rep,name=account,proto3" json:"account,omitempty"`
	Loan    []*Loan    `protobuf:"bytes,4,rep,name=loan,proto3" json:"loan,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_execute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_execute_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Request) GetAccount() []*Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *Request) GetLoan() []*Loan {
	if x != nil {
		return x.Loan
	}
	return nil
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentDateBalance int32 `protobuf:"varint,1,opt,name=payment_date_balance,json=paymentDateBalance,proto3" json:"payment_date_balance,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execute_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_execute_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_execute_proto_rawDescGZIP(), []int{1}
}

func (x *Account) GetPaymentDateBalance() int32 {
	if x != nil {
		return x.PaymentDateBalance
	}
	return 0
}

type Loan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment           int32                  `protobuf:"varint,1,opt,name=payment,proto3" json:"payment,omitempty"`
	PaymentDate       *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=payment_date,json=paymentDate,proto3" json:"payment_date,omitempty"`
	ActualPaymentDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=actual_payment_date,json=actualPaymentDate,proto3" json:"actual_payment_date,omitempty"`
}

func (x *Loan) Reset() {
	*x = Loan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execute_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Loan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loan) ProtoMessage() {}

func (x *Loan) ProtoReflect() protoreflect.Message {
	mi := &file_execute_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Loan.ProtoReflect.Descriptor instead.
func (*Loan) Descriptor() ([]byte, []int) {
	return file_execute_proto_rawDescGZIP(), []int{2}
}

func (x *Loan) GetPayment() int32 {
	if x != nil {
		return x.Payment
	}
	return 0
}

func (x *Loan) GetPaymentDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PaymentDate
	}
	return nil
}

func (x *Loan) GetActualPaymentDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ActualPaymentDate
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string    `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	ExecuteDate string    `protobuf:"bytes,2,opt,name=execute_date,json=executeDate,proto3" json:"execute_date,omitempty"`
	Results     []*Result `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execute_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_execute_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_execute_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Response) GetExecuteDate() string {
	if x != nil {
		return x.ExecuteDate
	}
	return ""
}

func (x *Response) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnoughMoneyByMonths          []int32 `protobuf:"varint,1,rep,packed,name=enough_money_by_months,json=enoughMoneyByMonths,proto3" json:"enough_money_by_months,omitempty"`
	DelinquencyByMonths          []int32 `protobuf:"varint,2,rep,packed,name=delinquency_by_months,json=delinquencyByMonths,proto3" json:"delinquency_by_months,omitempty"`
	DelinquencyDurationDaysTotal int32   `protobuf:"varint,3,opt,name=delinquency_duration_days_total,json=delinquencyDurationDaysTotal,proto3" json:"delinquency_duration_days_total,omitempty"`
	DelinquencySumTotal          int32   `protobuf:"varint,4,opt,name=delinquency_sum_total,json=delinquencySumTotal,proto3" json:"delinquency_sum_total,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execute_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_execute_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_execute_proto_rawDescGZIP(), []int{4}
}

func (x *Result) GetEnoughMoneyByMonths() []int32 {
	if x != nil {
		return x.EnoughMoneyByMonths
	}
	return nil
}

func (x *Result) GetDelinquencyByMonths() []int32 {
	if x != nil {
		return x.DelinquencyByMonths
	}
	return nil
}

func (x *Result) GetDelinquencyDurationDaysTotal() int32 {
	if x != nil {
		return x.DelinquencyDurationDaysTotal
	}
	return 0
}

func (x *Result) GetDelinquencySumTotal() int32 {
	if x != nil {
		return x.DelinquencySumTotal
	}
	return 0
}

var File_execute_proto protoreflect.FileDescriptor

var file_execute_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22,
	0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x04, 0x6c, 0x6f,
	0x61, 0x6e, 0x22, 0x3b, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a,
	0x14, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22,
	0xab, 0x01, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x4a, 0x0a, 0x13, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x61, 0x63, 0x74, 0x75,
	0x61, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x22, 0x73, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x22, 0xec, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a,
	0x16, 0x65, 0x6e, 0x6f, 0x75, 0x67, 0x68, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x62, 0x79,
	0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x13, 0x65,
	0x6e, 0x6f, 0x75, 0x67, 0x68, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x42, 0x79, 0x4d, 0x6f, 0x6e, 0x74,
	0x68, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x5f, 0x62, 0x79, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x13, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x79,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x12, 0x45, 0x0a, 0x1f, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x61, 0x79, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x1c, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a,
	0x15, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73, 0x75, 0x6d,
	0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x64, 0x65,
	0x6c, 0x69, 0x6e, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x75, 0x6d, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x32, 0x7f, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x3f, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_execute_proto_rawDescOnce sync.Once
	file_execute_proto_rawDescData = file_execute_proto_rawDesc
)

func file_execute_proto_rawDescGZIP() []byte {
	file_execute_proto_rawDescOnce.Do(func() {
		file_execute_proto_rawDescData = protoimpl.X.CompressGZIP(file_execute_proto_rawDescData)
	})
	return file_execute_proto_rawDescData
}

var file_execute_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_execute_proto_goTypes = []interface{}{
	(*Request)(nil),               // 0: strategy.Request
	(*Account)(nil),               // 1: strategy.Account
	(*Loan)(nil),                  // 2: strategy.Loan
	(*Response)(nil),              // 3: strategy.Response
	(*Result)(nil),                // 4: strategy.Result
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_execute_proto_depIdxs = []int32{
	1, // 0: strategy.Request.account:type_name -> strategy.Account
	2, // 1: strategy.Request.loan:type_name -> strategy.Loan
	5, // 2: strategy.Loan.payment_date:type_name -> google.protobuf.Timestamp
	5, // 3: strategy.Loan.actual_payment_date:type_name -> google.protobuf.Timestamp
	4, // 4: strategy.Response.results:type_name -> strategy.Result
	6, // 5: strategy.Executor.HealthCheck:input_type -> google.protobuf.Empty
	0, // 6: strategy.Executor.Execute:input_type -> strategy.Request
	6, // 7: strategy.Executor.HealthCheck:output_type -> google.protobuf.Empty
	3, // 8: strategy.Executor.Execute:output_type -> strategy.Response
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_execute_proto_init() }
func file_execute_proto_init() {
	if File_execute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_execute_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_execute_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_execute_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Loan); i {
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
		file_execute_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_execute_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_execute_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_execute_proto_goTypes,
		DependencyIndexes: file_execute_proto_depIdxs,
		MessageInfos:      file_execute_proto_msgTypes,
	}.Build()
	File_execute_proto = out.File
	file_execute_proto_rawDesc = nil
	file_execute_proto_goTypes = nil
	file_execute_proto_depIdxs = nil
}
