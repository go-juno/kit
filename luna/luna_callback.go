package luna

import (
	"context"
	"sync"

	"github.com/go-juno/kit/api/grpc/protos"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

// Handler 处理回调方法
// 在回调方法中禁止调用
type Handler func(context.Context, *protos.CallbackRequest) (*protos.CallbackReply, error)

// CallbackServer luna的回调服务
type CallbackServer interface {
	// 注册服务到grpc
	Register(srv grpc.ServiceRegistrar)
	// RegisterSubServer 注册子服务
	RegisterSubServer(processID int64, handle Handler) (err error)
	// UnRegisterSubServer 取消注册子服务
	UnRegisterSubServer(processID int64)
}

// CallbackServerConfig 配置选项
type CallbackServerConfig struct {
	Host string
	Port uint
}

type callbackServer struct {
	Config       *CallbackServerConfig
	SubServerMap *sync.Map
}

func (s *callbackServer) DoCallback(ctx context.Context, in *protos.CallbackRequest) (out *protos.CallbackReply, err error) {

	sub, ok := s.SubServerMap.Load(in.ProcessId)
	if !ok || sub == nil {
		err = xerrors.Errorf("未找到对应的流程监听业务")
		return
	}
	handler := sub.(Handler)
	out, err = handler(ctx, in)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *callbackServer) Register(srv grpc.ServiceRegistrar) {
	protos.RegisterLunaCallbackServiceServer(srv, s)
}

func (s *callbackServer) RegisterSubServer(processID int64, handle Handler) (err error) {

	if _, ok := s.SubServerMap.Load(processID); ok {
		err = xerrors.New("该服务已经在监听列表中")
		return
	}
	s.SubServerMap.Store(processID, handle)
	return
}

func (s *callbackServer) UnRegisterSubServer(processID int64) {
	s.SubServerMap.Delete(processID)
}

// NewCallbackServer 创建回调服务
func NewCallbackServer(cfg CallbackServerConfig) CallbackServer {
	return &callbackServer{
		Config:       &cfg,
		SubServerMap: new(sync.Map),
	}

}
