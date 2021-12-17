package middleware

import (
	"fmt"

	"git.yupaopao.com/ops-public/kit/api/grpc/protos"
	"git.yupaopao.com/ops-public/kit/res"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

// RoleKey 角色上下文字段
const RoleKey = "role"

// Role 角色中间件
type Role interface {
	// Handler 返回中间件处理方法
	Handler() gin.HandlerFunc
	// OneOf 角色认证中间件
	OneOf(keys ...string) gin.HandlerFunc
	// Has 角色认证中间件
	Has(key string) gin.HandlerFunc
	// Both 角色认证中间件
	Both(keys ...string) gin.HandlerFunc
}

type role struct {
	Client protos.RoleClient
	Config *RoleConfig
}

// RoleConfig 角色中间价配置
type RoleConfig struct {
	Host    string
	Port    uint
	SiteKey string
}

// NewRole 创建角色中间价实例
func NewRole(cfg RoleConfig) (res Role, err error) {
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	res = &role{
		Client: protos.NewRoleClient(conn),
		Config: &cfg,
	}
	return
}

// Handler 返回中间件处理方法
func (m *role) Handler() gin.HandlerFunc {
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
		roles := make(roles, len(out.Role))
		for i, v := range out.Role {
			roles[i] = v.RoleKey
		}
		c.Set(RoleKey, roles)
		c.Next()
	}
}

func (m *role) OneOf(keys ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		l, ok := GetBindRole(c)
		if !ok {
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
			s := make(roles, len(out.Role))
			for i, v := range out.Role {
				s[i] = v.RoleKey
			}
			c.Set(RoleKey, s)
			l = s
		}
		if l.OneOf(keys...) {
			c.Next()
			return
		}
		res.NoPermissionRes(c)
	}
}

func (m *role) Has(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		l, ok := GetBindRole(c)
		if !ok {
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
			s := make(roles, len(out.Role))
			for i, v := range out.Role {
				s[i] = v.RoleKey
			}
			c.Set(RoleKey, s)
			l = s
		}
		if l.Has(key) {
			c.Next()
			return
		}
		res.NoPermissionRes(c)
	}
}

func (m *role) Both(keys ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		l, ok := GetBindRole(c)
		if !ok {
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
			s := make(roles, len(out.Role))
			for i, v := range out.Role {
				s[i] = v.RoleKey
			}
			c.Set(RoleKey, s)
			l = s
		}
		if l.Both(keys...) {
			c.Next()
			return
		}
		res.NoPermissionRes(c)
	}
}

// GetBindRole 获取绑定角色
func GetBindRole(c *gin.Context) (role Roles, exists bool) {
	var temp interface{}
	if temp, exists = c.Get(RoleKey); exists {
		role, exists = temp.(Roles)
		return
	}
	return
}

// MustGetBindRole 获取绑定角色
func MustGetBindRole(c *gin.Context) (r Roles) {
	var exists bool
	r, exists = GetBindRole(c)
	if exists {
		return
	}
	return make(roles, 0)
}

// Roles 角色数据
type Roles interface {
	// Has 判断是否具有某个角色
	// key 需要查询的角色唯一 key 值
	Has(key string) bool
	// OneOf 判断为多个角色之一
	// keys 角色的唯一 key 值
	OneOf(keys ...string) bool
	// Both 判断同时为多个角色
	// keys 角色的唯一 key 值
	Both(keys ...string) bool
}

type roles []string

func (r roles) Has(key string) bool {
	for _, v := range r {
		if v == key {
			return true
		}
	}
	return false
}

func (r roles) OneOf(keys ...string) bool {
	for _, v := range r {
		for _, role := range keys {
			if v == role {
				return true
			}
		}
	}
	return false
}

func (r roles) Both(keys ...string) bool {
	if len(r) < len(keys) {
		return false
	}
	for _, role := range keys {
		find := false
		for _, v := range r {
			if v == role {
				find = true
				break
			}
		}
		if !find {
			return false
		}
	}
	return true
}
