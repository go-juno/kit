package middleware

import (
	"fmt"

	"git.yupaopao.com/ops-public/kit/api/grpc/protos"
	"git.yupaopao.com/ops-public/kit/res"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

// Privilege 中间件
type Privilege interface {
	// Handler 返回中间件处理方法
	Handler(category string, shape string, point string) gin.HandlerFunc
}

// privilege 权限中间件
type privilege struct {
	Client protos.PrivilegeClient
	Config *PrivilegeConfig
}

// PrivilegeConfig 权限中间价配置
type PrivilegeConfig struct {
	Host    string
	Port    uint
	SiteKey string
}

// NewPrivilege 创建权限中间价实例
func NewPrivilege(cfg PrivilegeConfig) (res Privilege, err error) {
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	res = &privilege{
		Client: protos.NewPrivilegeClient(conn),
		Config: &cfg,
	}
	return
}

// Handler 返回中间件处理方法
func (m *privilege) Handler(category string, shape string, point string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if category == "" && shape == "" && point == "" {
			c.Next()
			return
		}
		// 先获取domain
		domain := c.GetString(DomainKey)
		in := &protos.GetUserPrivilegeAllParam{
			Domain:  domain,
			SiteKey: m.Config.SiteKey,
		}
		out, err := m.Client.GetUserPrivilegeAll(c, in)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.SystemErrorRes(c, err)
			return
		}
		has := false
		for _, v := range out.Privilege {
			if v.Category == category && v.Shape == shape && v.Point == point {
				has = true
				break
			}
		}
		if !has {
			res.NoPermissionRes(c)
			return
		}
		c.Next()
	}
}
