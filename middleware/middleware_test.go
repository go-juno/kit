package middleware_test

import (
	"testing"

	"git.yupaopao.com/ops-public/kit/middleware"
	"github.com/gin-gonic/gin"
)

func TestRole(t *testing.T) {
	config := &middleware.RoleConfig{
		Host:    "test-auth.yupaopao.com",
		Port:    8088,
		SiteKey: "53c9014f9afc4b378ed3bda86876d67a",
	}

	roleMiddleware, err := middleware.NewRole(*config)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/ping", roleMiddleware.Handler(), func(c *gin.Context) {
		// 获取中间件绑定的角色上下文
		roles := middleware.MustGetBindRole(c)
		// 判断为某一角色
		if roles.Has("880e0a7c37fc41229e6505b535dd757f") {
			c.JSON(200, gin.H{
				"message": "pong 角色 A ",
			})
			return
		}
		// 判断为多个角色之一
		if roles.OneOf("880e0a7c37fc41229e6505b535dd757f",
			"b6de126a99504a78be9bc90480d84b52",
			"cd26497a22b4400b90ab8533ede63537") {
			c.JSON(200, gin.H{
				"message": "角色 A|B|C",
			})
			return
		}
		// 判断同时为多个角色
		if roles.Both("880e0a7c37fc41229e6505b535dd757f",
			"b6de126a99504a78be9bc90480d84b52",
			"cd26497a22b4400b90ab8533ede63537") {
			c.JSON(200, gin.H{
				"message": "角色 A&B&C",
			})
			return
		}
		c.JSON(403, nil)
	})
	_ = r.Run()
}

func TestPrivilege(t *testing.T) {
	config := &middleware.PrivilegeConfig{
		Host:    "test-auth.yupaopao.com",
		Port:    8088,
		SiteKey: "53c9014f9afc4b378ed3bda86876d67a",
	}

	privilegeMiddleware, err := middleware.NewPrivilege(*config)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/ping", privilegeMiddleware.Handler("category", "", ""), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run()
}

func TestDataAccess(t *testing.T) {
	config := &middleware.DataAccessConfig{
		Host:    "test-auth.yupaopao.com",
		Port:    8088,
		SiteKey: "53c9014f9afc4b378ed3bda86876d67a",
	}

	dataAccessMiddleware, err := middleware.NewDataAccess(*config)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/ping", dataAccessMiddleware.Handler(
		"880e0a7c37fc41229e6505b535dd757f",
		"880e0a7c37fc41229e6505b535dd757f"),
		func(c *gin.Context) {
			if middleware.GetBindDataAccess(c) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
				return
			}
			c.JSON(403, nil)
		})
	_ = r.Run()
}

func TestAuth(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", middleware.Auth(),
		func(c *gin.Context) {
			// 获取绑定上下文用户域控
			auth := middleware.GetBindAuth(c)
			if middleware.GetBindDataAccess(c) {
				c.JSON(200, gin.H{
					"message": auth,
				})
				return
			}
			c.JSON(403, nil)
		})
	_ = r.Run()
}
