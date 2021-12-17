package res

import (
	"log"

	kerror "git.yupaopao.com/ops-public/kit/error"
	"git.yupaopao.com/ops-public/kit/kerrors"
	"github.com/gin-gonic/gin"
)

// Status 响应状态
type Status string

//  响应状态枚举
const (
	// Success 成功
	Success Status = "success"
	// Failure 失败
	Failure Status = "failure"
	// ParamCheck 参数错误
	ParamCheck Status = "param check"
	// AuthOverdue 无效令牌
	AuthOverdue Status = "auth overdue"
	// NoPermission 没有权限
	NoPermission Status = "no permission"
	// SystemError 系统内部错误
	SystemError Status = "system error"
)

// Response 响应封装
type Response struct {
	Status Status      `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

// SuccessRes 成功响应
func SuccessRes(c *gin.Context, data interface{}) {
	res := Response{
		Status: Success,
		Msg:    nil,
		Data:   data,
	}
	c.JSON(200, res)
}

// FailureRes 失败响应
func FailureRes(c *gin.Context, err error) {
	log.Printf("failure response! error:%+v", err)
	res := Response{
		Status: "",
		Msg:    nil,
		Data:   nil,
	}
	err = kerror.Unwrap(err)
	if _, ok := err.(*kerrors.ErrBussiness); ok {
		res.Status = Success
		res.Msg = err.Error()
	} else {
		res.Msg = err.Error()
		res.Status = Failure
	}
	c.AbortWithStatusJSON(200, res)
}

// ParamCheckRes 参数错误响应
func ParamCheckRes(c *gin.Context, err error) {
	log.Printf("param check! error:%+v", err)
	res := Response{
		Status: ParamCheck,
		Msg:    kerror.Unwrap(err).Error(),
		Data:   nil,
	}
	c.AbortWithStatusJSON(200, res)
}

// AuthOverdueRes 无效令牌响应
func AuthOverdueRes(c *gin.Context) {
	res := Response{
		Status: AuthOverdue,
		Msg:    "未登录",
		Data:   nil,
	}
	c.AbortWithStatusJSON(200, res)
}

// NoPermissionRes 权限不足响应
func NoPermissionRes(c *gin.Context) {
	res := Response{
		Status: NoPermission,
		Msg:    "无权限",
		Data:   nil,
	}
	c.AbortWithStatusJSON(200, res)
}

// SystemErrorRes 系统错误响应
func SystemErrorRes(c *gin.Context, err error) {
	res := Response{
		Status: "",
		Msg:    nil,
		Data:   nil,
	}
	err = kerror.Unwrap(err)
	if _, ok := err.(*kerrors.ErrBussiness); ok {
		res.Status = Success
		res.Msg = err.Error()
	} else {
		res.Msg = err.Error()
		res.Status = SystemError
	}
	c.AbortWithStatusJSON(200, res)
}

// CommonRes 响应
func CommonRes(c *gin.Context, status Status, msg string, data interface{}) {
	res := Response{
		Status: status,
		Msg:    msg,
		Data:   data,
	}
	c.JSON(200, res)
}

type List struct {
	Items interface{} `json:"items"`
	Total int64       `json:"total"`
}
