/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//
//Copyright 2024 The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// To regenerate api.pb.go run `hack/update-codegen.sh protobindings`

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: staging/src/k8s.io/externaljwt/apis/v1alpha1/api.proto

package v1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExternalJWTSigner_Sign_FullMethodName      = "/v1alpha1.ExternalJWTSigner/Sign"
	ExternalJWTSigner_FetchKeys_FullMethodName = "/v1alpha1.ExternalJWTSigner/FetchKeys"
	ExternalJWTSigner_Metadata_FullMethodName  = "/v1alpha1.ExternalJWTSigner/Metadata"
)

// ExternalJWTSignerClient is the client API for ExternalJWTSigner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// This service is served by a process on a local Unix Domain Socket.
type ExternalJWTSignerClient interface {
	// Sign takes a serialized JWT payload, and returns the serialized header and
	// signature. The caller can then assemble the JWT from the header, payload,
	// and signature. Signature can be generated by signing
	// `base64url(header) + "." + base64url(payload)` with signing key.
	//
	// The plugin MUST set a key id in the returned JWT header.
	Sign(ctx context.Context, in *SignJWTRequest, opts ...grpc.CallOption) (*SignJWTResponse, error)
	// FetchKeys returns the set of public keys that are trusted to sign
	// Kubernetes service account tokens. Kube-apiserver will call this RPC:
	//
	// * Every time it tries to validate a JWT from the service account issuer with an unknown key ID, and
	//
	//   - Periodically, so it can serve reasonably-up-to-date keys from the OIDC
	//     JWKs endpoint.
	FetchKeys(ctx context.Context, in *FetchKeysRequest, opts ...grpc.CallOption) (*FetchKeysResponse, error)
	// Metadata is meant to be called once on startup.
	// Enables sharing metadata with kube-apiserver (eg: the max token lifetime that signer supports)
	Metadata(ctx context.Context, in *MetadataRequest, opts ...grpc.CallOption) (*MetadataResponse, error)
}

type externalJWTSignerClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalJWTSignerClient(cc grpc.ClientConnInterface) ExternalJWTSignerClient {
	return &externalJWTSignerClient{cc}
}

func (c *externalJWTSignerClient) Sign(ctx context.Context, in *SignJWTRequest, opts ...grpc.CallOption) (*SignJWTResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignJWTResponse)
	err := c.cc.Invoke(ctx, ExternalJWTSigner_Sign_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalJWTSignerClient) FetchKeys(ctx context.Context, in *FetchKeysRequest, opts ...grpc.CallOption) (*FetchKeysResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchKeysResponse)
	err := c.cc.Invoke(ctx, ExternalJWTSigner_FetchKeys_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalJWTSignerClient) Metadata(ctx context.Context, in *MetadataRequest, opts ...grpc.CallOption) (*MetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MetadataResponse)
	err := c.cc.Invoke(ctx, ExternalJWTSigner_Metadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalJWTSignerServer is the server API for ExternalJWTSigner service.
// All implementations must embed UnimplementedExternalJWTSignerServer
// for forward compatibility.
//
// This service is served by a process on a local Unix Domain Socket.
type ExternalJWTSignerServer interface {
	// Sign takes a serialized JWT payload, and returns the serialized header and
	// signature. The caller can then assemble the JWT from the header, payload,
	// and signature. Signature can be generated by signing
	// `base64url(header) + "." + base64url(payload)` with signing key.
	//
	// The plugin MUST set a key id in the returned JWT header.
	Sign(context.Context, *SignJWTRequest) (*SignJWTResponse, error)
	// FetchKeys returns the set of public keys that are trusted to sign
	// Kubernetes service account tokens. Kube-apiserver will call this RPC:
	//
	// * Every time it tries to validate a JWT from the service account issuer with an unknown key ID, and
	//
	//   - Periodically, so it can serve reasonably-up-to-date keys from the OIDC
	//     JWKs endpoint.
	FetchKeys(context.Context, *FetchKeysRequest) (*FetchKeysResponse, error)
	// Metadata is meant to be called once on startup.
	// Enables sharing metadata with kube-apiserver (eg: the max token lifetime that signer supports)
	Metadata(context.Context, *MetadataRequest) (*MetadataResponse, error)
	mustEmbedUnimplementedExternalJWTSignerServer()
}

// UnimplementedExternalJWTSignerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExternalJWTSignerServer struct{}

func (UnimplementedExternalJWTSignerServer) Sign(context.Context, *SignJWTRequest) (*SignJWTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedExternalJWTSignerServer) FetchKeys(context.Context, *FetchKeysRequest) (*FetchKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchKeys not implemented")
}
func (UnimplementedExternalJWTSignerServer) Metadata(context.Context, *MetadataRequest) (*MetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Metadata not implemented")
}
func (UnimplementedExternalJWTSignerServer) mustEmbedUnimplementedExternalJWTSignerServer() {}
func (UnimplementedExternalJWTSignerServer) testEmbeddedByValue()                           {}

// UnsafeExternalJWTSignerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalJWTSignerServer will
// result in compilation errors.
type UnsafeExternalJWTSignerServer interface {
	mustEmbedUnimplementedExternalJWTSignerServer()
}

func RegisterExternalJWTSignerServer(s grpc.ServiceRegistrar, srv ExternalJWTSignerServer) {
	// If the following call pancis, it indicates UnimplementedExternalJWTSignerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExternalJWTSigner_ServiceDesc, srv)
}

func _ExternalJWTSigner_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignJWTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalJWTSignerServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalJWTSigner_Sign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalJWTSignerServer).Sign(ctx, req.(*SignJWTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalJWTSigner_FetchKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalJWTSignerServer).FetchKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalJWTSigner_FetchKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalJWTSignerServer).FetchKeys(ctx, req.(*FetchKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalJWTSigner_Metadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalJWTSignerServer).Metadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalJWTSigner_Metadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalJWTSignerServer).Metadata(ctx, req.(*MetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExternalJWTSigner_ServiceDesc is the grpc.ServiceDesc for ExternalJWTSigner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalJWTSigner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.ExternalJWTSigner",
	HandlerType: (*ExternalJWTSignerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sign",
			Handler:    _ExternalJWTSigner_Sign_Handler,
		},
		{
			MethodName: "FetchKeys",
			Handler:    _ExternalJWTSigner_FetchKeys_Handler,
		},
		{
			MethodName: "Metadata",
			Handler:    _ExternalJWTSigner_Metadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staging/src/k8s.io/externaljwt/apis/v1alpha1/api.proto",
}
