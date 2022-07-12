// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package banking

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BankingServiceClient is the client API for BankingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankingServiceClient interface {
	GetAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error)
	GetAccountBalance(ctx context.Context, in *GetAccountBalanceRequest, opts ...grpc.CallOption) (*GetAccountBalanceResponse, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetTransfers(ctx context.Context, in *ListTransfersRequest, opts ...grpc.CallOption) (*ListTransfersResponse, error)
	CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferResponse, error)
}

type bankingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBankingServiceClient(cc grpc.ClientConnInterface) BankingServiceClient {
	return &bankingServiceClient{cc}
}

func (c *bankingServiceClient) GetAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	out := new(ListAccountsResponse)
	err := c.cc.Invoke(ctx, "/banking.BankingService/GetAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankingServiceClient) GetAccountBalance(ctx context.Context, in *GetAccountBalanceRequest, opts ...grpc.CallOption) (*GetAccountBalanceResponse, error) {
	out := new(GetAccountBalanceResponse)
	err := c.cc.Invoke(ctx, "/banking.BankingService/GetAccountBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankingServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/banking.BankingService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankingServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/banking.BankingService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankingServiceClient) GetTransfers(ctx context.Context, in *ListTransfersRequest, opts ...grpc.CallOption) (*ListTransfersResponse, error) {
	out := new(ListTransfersResponse)
	err := c.cc.Invoke(ctx, "/banking.BankingService/GetTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankingServiceClient) CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferResponse, error) {
	out := new(CreateTransferResponse)
	err := c.cc.Invoke(ctx, "/banking.BankingService/CreateTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankingServiceServer is the server API for BankingService service.
// All implementations must embed UnimplementedBankingServiceServer
// for forward compatibility
type BankingServiceServer interface {
	GetAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	GetAccountBalance(context.Context, *GetAccountBalanceRequest) (*GetAccountBalanceResponse, error)
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetTransfers(context.Context, *ListTransfersRequest) (*ListTransfersResponse, error)
	CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferResponse, error)
	mustEmbedUnimplementedBankingServiceServer()
}

// UnimplementedBankingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBankingServiceServer struct {
}

func (UnimplementedBankingServiceServer) GetAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}
func (UnimplementedBankingServiceServer) GetAccountBalance(context.Context, *GetAccountBalanceRequest) (*GetAccountBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountBalance not implemented")
}
func (UnimplementedBankingServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedBankingServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBankingServiceServer) GetTransfers(context.Context, *ListTransfersRequest) (*ListTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransfers not implemented")
}
func (UnimplementedBankingServiceServer) CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransfer not implemented")
}
func (UnimplementedBankingServiceServer) mustEmbedUnimplementedBankingServiceServer() {}

// UnsafeBankingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankingServiceServer will
// result in compilation errors.
type UnsafeBankingServiceServer interface {
	mustEmbedUnimplementedBankingServiceServer()
}

func RegisterBankingServiceServer(s grpc.ServiceRegistrar, srv BankingServiceServer) {
	s.RegisterService(&BankingService_ServiceDesc, srv)
}

func _BankingService_GetAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankingServiceServer).GetAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banking.BankingService/GetAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankingServiceServer).GetAccounts(ctx, req.(*ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankingService_GetAccountBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankingServiceServer).GetAccountBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banking.BankingService/GetAccountBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankingServiceServer).GetAccountBalance(ctx, req.(*GetAccountBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankingService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankingServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banking.BankingService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankingServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankingService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankingServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banking.BankingService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankingServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankingService_GetTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankingServiceServer).GetTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banking.BankingService/GetTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankingServiceServer).GetTransfers(ctx, req.(*ListTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankingService_CreateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankingServiceServer).CreateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banking.BankingService/CreateTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankingServiceServer).CreateTransfer(ctx, req.(*CreateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BankingService_ServiceDesc is the grpc.ServiceDesc for BankingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "banking.BankingService",
	HandlerType: (*BankingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccounts",
			Handler:    _BankingService_GetAccounts_Handler,
		},
		{
			MethodName: "GetAccountBalance",
			Handler:    _BankingService_GetAccountBalance_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _BankingService_CreateAccount_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _BankingService_Login_Handler,
		},
		{
			MethodName: "GetTransfers",
			Handler:    _BankingService_GetTransfers_Handler,
		},
		{
			MethodName: "CreateTransfer",
			Handler:    _BankingService_CreateTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/banking/banking.proto",
}