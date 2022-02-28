package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-juno/kit/api/grpc/protos"
	"github.com/go-juno/kit/res"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// DataAccessKey 数据权限上下文字段
const DataAccessKey = "dataAccess"

// DataAccess 数据权限中间件
type DataAccess interface {
	Handler(roles ...string) gin.HandlerFunc
}

// dataAccess 数据权限中间件
type dataAccess struct {
	Client protos.RoleClient
	Config *DataAccessConfig
}

// DataAccessConfig 数据权限中间价配置
type DataAccessConfig struct {
	Host    string
	Port    uint
	SiteKey string
}

// NewDataAccess 创建数据权限中间价实例
func NewDataAccess(cfg DataAccessConfig) (res DataAccess, err error) {
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	res = &dataAccess{
		Client: protos.NewRoleClient(conn),
		Config: &cfg,
	}
	return
}

// Handler 返回中间件处理方法
// roles 角色的唯一 key 值
func (m *dataAccess) Handler(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 先获取domain
		domain := c.GetString(DomainKey)
		in := &protos.GetUserRoleParam{
			Domain:  domain,
			SiteKey: m.Config.SiteKey,
		}
		out, err := m.Client.GetUserRole(c, in)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.SystemErrorRes(c, err)
			return
		}
		find := false
	LOOP:
		for _, v := range out.Role {
			for _, k := range roles {
				if v.RoleKey == k {
					find = true
					break LOOP
				}
			}
		}
		c.Set(DataAccessKey, find)
		c.Next()
	}
}

// GetBindDataAccess 获取数据权限状态
func GetBindDataAccess(c *gin.Context) (access bool) {
	access = c.GetBool(DataAccessKey)
	return
}
