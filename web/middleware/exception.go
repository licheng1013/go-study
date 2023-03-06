package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"web/common"
)

func init() {
	common.App.Use(Recover)
}

// Recover 全局异常处理器
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack() //打印错误堆栈信息
			c.JSON(http.StatusOK, common.Fail(errorToString(r)))
			c.Abort() //终止后续接口调用，不加的话,还会继续接口后续代码
		}
	}()
	c.Next()
}

// NewServiceError 使用关键字 panic 抛出异常,否则无法识别到
func NewServiceError(text string) error {
	return &ServiceError{text}
}

type ServiceError struct {
	s string
}

func (e *ServiceError) Error() string {
	return e.s
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case *ServiceError:
		return v.Error()
	default:
		return fmt.Sprint(v)
	}
}
