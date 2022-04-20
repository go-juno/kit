package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ker "github.com/go-juno/kit/error"
	"golang.org/x/xerrors"
)

type HttpResponse interface {
	Success(ctx *gin.Context, data any)
	Unauthorized(c *gin.Context)
	Forbidden(c *gin.Context)
	BadRequest(ctx *gin.Context, err error)
	InternalServerError(ctx *gin.Context, err error)
}

type Status string

const (
	Success             Status = "Success"
	Unauthorized        Status = "Unauthorized"
	Forbidden           Status = "Forbidden"
	BadRequest          Status = "BadRequest"
	InternalServerError Status = "InternalServerError"
)

type ResponseMsg struct {
	Status Status `json:"sys_status"`
	Msg    any    `json:"message"`
	Data   any    `json:"data"`
}

type response struct{}

func (r *response) Success(c *gin.Context, data any) {
	res := ResponseMsg{
		Status: Success,
		Msg:    nil,
		Data:   data,
	}
	c.JSON(http.StatusForbidden, res)
}

func (r *response) InternalServerError(c *gin.Context, err error) {
	log.Printf("Internal Server Error! error:%+v", err)
	res := ResponseMsg{
		Status: InternalServerError,
		Msg:    ker.Unwrap(err).Error(),
		Data:   nil,
	}
	c.AbortWithStatusJSON(200, res)
}

func (r *response) BadRequest(c *gin.Context, err error) {
	log.Printf("Bad Request! error:%+v", err)
	res := ResponseMsg{
		Status: BadRequest,
		Msg:    ker.Unwrap(err).Error(),
		Data:   nil,
	}
	c.AbortWithStatusJSON(200, res)
}

func (r *response) Unauthorized(c *gin.Context) {
	res := ResponseMsg{
		Status: Unauthorized,
		Msg:    "未登录",
		Data:   nil,
	}
	c.AbortWithStatusJSON(200, res)
}

func (r *response) Forbidden(c *gin.Context) {
	res := ResponseMsg{
		Status: Forbidden,
		Msg:    "无权限",
		Data:   nil,
	}
	c.AbortWithStatusJSON(200, res)
}

func NewHttpResponse() HttpResponse {
	return &response{}
}

func HttpHandleFunc[HttpSchema, HttpSerialize, EndpointRequest, EndpointResponse any](res HttpResponse, httpSchemaTransform func(*gin.Context, HttpSchema) EndpointRequest, endpoint func(context.Context, EndpointRequest) (EndpointResponse, error), endpointResponseTransform func(*gin.Context, EndpointResponse) HttpSerialize) gin.HandlerFunc {
	return func(c *gin.Context) {
		var schema HttpSchema
		err := c.ShouldBind(&schema)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.BadRequest(c, err)
			return
		}
		ereq := httpSchemaTransform(c, schema)
		eres, err := endpoint(c, ereq)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.InternalServerError(c, err)
			return
		}
		res.Success(c, endpointResponseTransform(c, eres))
	}
}
