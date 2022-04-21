package middleware

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	h "github.com/go-juno/kit/http"
)

// Error 处理错误的中间件
func Error(res h.HttpResponse) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("err:%+v\n stack:%s", r, string(debug.Stack()))
				res.InternalServerError(c, fmt.Errorf("%v", r))
				return
			}
		}()
		c.Next()
	}
}
