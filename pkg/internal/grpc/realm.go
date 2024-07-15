package grpc

import (
	"context"
	"git.solsynth.dev/hydrogen/dealer/pkg/hyper"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"google.golang.org/grpc"
)

func (v *Server) ListCommunityRealm(ctx context.Context, request *proto.ListRealmRequest) (*proto.ListRealmResponse, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.ListRealmResponse, error) {
			out, err := proto.NewRealmClient(conn).ListCommunityRealm(ctx, request)
			return out, err
		},
	)
}

func (v *Server) ListAvailableRealm(ctx context.Context, request *proto.LookupUserRealmRequest) (*proto.ListRealmResponse, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.ListRealmResponse, error) {
			out, err := proto.NewRealmClient(conn).ListAvailableRealm(ctx, request)
			return out, err
		},
	)
}

func (v *Server) ListOwnedRealm(ctx context.Context, request *proto.LookupUserRealmRequest) (*proto.ListRealmResponse, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.ListRealmResponse, error) {
			out, err := proto.NewRealmClient(conn).ListOwnedRealm(ctx, request)
			return out, err
		},
	)
}

func (v *Server) GetRealm(ctx context.Context, request *proto.LookupRealmRequest) (*proto.RealmInfo, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.RealmInfo, error) {
			out, err := proto.NewRealmClient(conn).GetRealm(ctx, request)
			return out, err
		},
	)
}

func (v *Server) ListRealmMember(ctx context.Context, request *proto.RealmMemberLookupRequest) (*proto.ListRealmMemberResponse, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.ListRealmMemberResponse, error) {
			out, err := proto.NewRealmClient(conn).ListRealmMember(ctx, request)
			return out, err
		},
	)
}

func (v *Server) GetRealmMember(ctx context.Context, request *proto.RealmMemberLookupRequest) (*proto.MemberInfo, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.MemberInfo, error) {
			out, err := proto.NewRealmClient(conn).GetRealmMember(ctx, request)
			return out, err
		},
	)
}

func (v *Server) CheckRealmMemberPerm(ctx context.Context, request *proto.CheckRealmPermRequest) (*proto.CheckRealmPermResponse, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.CheckRealmPermResponse, error) {
			out, err := proto.NewRealmClient(conn).CheckRealmMemberPerm(ctx, request)
			return out, err
		},
	)
}
