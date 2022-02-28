package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-juno/kit/res"
	"golang.org/x/xerrors"
)

// DomainKey 域控上下文字段
const DomainKey = "domain"

// Claims 令牌数据摘要
type Claims struct {
	LoginName string `json:"loginName"`
}

// ParseToken 解析令牌
func ParseToken(token string) (c *Claims, err error) {
	tokenSplit := strings.Split(token, "Bearer ")
	if len(tokenSplit) < 2 {
		err = xerrors.New("token不正确")
		return
	}

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		err = xerrors.New("token不正确")
		return
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		err = xerrors.New("token不正确")
		return
	}
	var cs Claims
	err = json.Unmarshal(payload, &cs)
	if err != nil {
		err = xerrors.New("token不正确")
		return
	}
	c = &cs
	return
}

// Auth auth中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		claims, err := ParseToken(token)
		if err != nil {
			res.AuthOverdueRes(c)
			return
		}
		c.Set(DomainKey, claims.LoginName)
		c.Next()
	}
}

// GetBindAuth 获取绑定用户
func GetBindAuth(c *gin.Context) (auth string) {
	auth = c.GetString(DomainKey)
	return
}

// Error 处理错误的中间件
func Error(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("err:%+v\n stack:%s", r, string(debug.Stack()))
			response := res.Response{
				Status: res.SystemError,
				Msg:    fmt.Sprintf("%v", r),
				Data:   nil,
			}
			c.AbortWithStatusJSON(http.StatusOK, response)
			return
		}
	}()
	c.Next()
}

//定制日志格式
func LoggerFormate() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		token := param.Request.Header.Get("Authorization")

		claims, _ := ParseToken(token)
		loginName := ""
		if claims != nil {
			loginName = claims.LoginName
		}

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s-[%s]\"\n",
			param.ClientIP,
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
			loginName,
		)
	})
}
