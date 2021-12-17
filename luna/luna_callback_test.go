package luna_test

import (
	"context"
	"testing"

	"git.yupaopao.com/ops-public/kit/api/grpc/protos"
	"git.yupaopao.com/ops-public/kit/luna"
	"google.golang.org/grpc"
)

func TestLunaCallback(t *testing.T) {
	config := luna.CallbackServerConfig{
		Host: "luna-approval",
		Port: 8088,
	}
	srv := luna.NewCallbackServer(config)

	s := grpc.NewServer()
	// 注册主服务
	srv.Register(s)
	// 业务模块
	endpoint := &Endpoint{}
	// 注册子服务
	_ = srv.RegisterSubServer(884831321770168320, endpoint.DoSomething)
}

// 业务处理
type Endpoint struct {
}

// 具体业务代码
func (e *Endpoint) DoSomething(c context.Context, in *protos.CallbackRequest) (out *protos.CallbackReply, err error) {

	return
}
