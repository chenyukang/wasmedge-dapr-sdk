// The wasm impl of dapr sdk

package client

import (
	"context"
	"encoding/json"
	v1 "github.com/dapr/dapr/pkg/proto/common/v1"
	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var _             pb.DaprClient = (*defaultClientImpl)(nil)

type action int32

const (
	actionGet action = 1 + iota
	actionSet
	actionDelete
)

type clientImpl struct {
	defaultClientImpl
}

func (c *clientImpl) GetState(ctx context.Context, in *pb.GetStateRequest, opts ...grpc.CallOption) (*pb.GetStateResponse, error) {
	metadataStr, _ := json.Marshal(in.GetMetadata())
	resStr, err := daprState(actionGet, in.StoreName, in.Key, nil, int32(in.Consistency), string(metadataStr))
	if len(err) > 0 {
		return nil, errors.New(err)
	}
	var res pb.GetStateResponse
	Jerr := json.Unmarshal([]byte(resStr), &res)
	if Jerr != nil {
		return nil, Jerr
	}
	return &res, nil
}

func (c *clientImpl) SaveState(ctx context.Context, in *pb.SaveStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

func (c *clientImpl) DeleteState(ctx context.Context, in *pb.DeleteStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

type defaultClientImpl struct{}

func (c *defaultClientImpl) QueryStateAlpha1(ctx context.Context, in *pb.QueryStateRequest, opts ...grpc.CallOption) (*pb.QueryStateResponse, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) GetConfigurationAlpha1(ctx context.Context, in *pb.GetConfigurationRequest, opts ...grpc.CallOption) (*pb.GetConfigurationResponse, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) SubscribeConfigurationAlpha1(ctx context.Context, in *pb.SubscribeConfigurationRequest, opts ...grpc.CallOption) (pb.Dapr_SubscribeConfigurationAlpha1Client, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) InvokeService(ctx context.Context, in *pb.InvokeServiceRequest, opts ...grpc.CallOption) (*v1.InvokeResponse, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) GetState(ctx context.Context, in *pb.GetStateRequest, opts ...grpc.CallOption) (*pb.GetStateResponse, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) GetBulkState(ctx context.Context, in *pb.GetBulkStateRequest, opts ...grpc.CallOption) (*pb.GetBulkStateResponse, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) SaveState(ctx context.Context, in *pb.SaveStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) DeleteState(ctx context.Context, in *pb.DeleteStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) DeleteBulkState(ctx context.Context, in *pb.DeleteBulkStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) ExecuteStateTransaction(ctx context.Context, in *pb.ExecuteStateTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) PublishEvent(ctx context.Context, in *pb.PublishEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) InvokeBinding(ctx context.Context, in *pb.InvokeBindingRequest, opts ...grpc.CallOption) (*pb.InvokeBindingResponse, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) GetSecret(ctx context.Context, in *pb.GetSecretRequest, opts ...grpc.CallOption) (*pb.GetSecretResponse, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) GetBulkSecret(ctx context.Context, in *pb.GetBulkSecretRequest, opts ...grpc.CallOption) (*pb.GetBulkSecretResponse, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) RegisterActorTimer(ctx context.Context, in *pb.RegisterActorTimerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) UnregisterActorTimer(ctx context.Context, in *pb.UnregisterActorTimerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) RegisterActorReminder(ctx context.Context, in *pb.RegisterActorReminderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) UnregisterActorReminder(ctx context.Context, in *pb.UnregisterActorReminderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) GetActorState(ctx context.Context, in *pb.GetActorStateRequest, opts ...grpc.CallOption) (*pb.GetActorStateResponse, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) ExecuteActorStateTransaction(ctx context.Context, in *pb.ExecuteActorStateTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) InvokeActor(ctx context.Context, in *pb.InvokeActorRequest, opts ...grpc.CallOption) (*pb.InvokeActorResponse, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) GetMetadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pb.GetMetadataResponse, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) SetMetadata(ctx context.Context, in *pb.SetMetadataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}

func (c *defaultClientImpl) Shutdown(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, ErrUnimplemented
}
