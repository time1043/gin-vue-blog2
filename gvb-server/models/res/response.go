package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time1043/gvb-server/utils"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

const (
	Success = 0
	Error   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: code, Data: data, Msg: msg})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}
func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "成功", c)
}
func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}
func OkWith(c *gin.Context) {
	Result(Success, map[string]any{}, "成功", c)
}
func OkWithList(count int64, list any, ctx *gin.Context) {
	OkWithData(ListResponse[any]{Count: count, List: list}, ctx)
}

func Fail(data any, msg string, ctx *gin.Context) {
	Result(Error, data, msg, ctx)
}
func FailWithMessage(msg string, ctx *gin.Context) {
	Result(Error, map[string]any{}, msg, ctx)
}
func FailWithError(err error, obj any, ctx *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, ctx)
}
func FailWithCode(code ErrorCode, ctx *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, ctx)
		return
	}
	Result(Error, map[string]any{}, "未知错误", ctx)
}
