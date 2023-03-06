package middleware

import (
	"log"
	"reflect"
	"runtime/debug"
	"testing"
)

func TestError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			debug.PrintStack()

			log.Println("类型: ", reflect.TypeOf(r))

			switch v := r.(type) {
			case *ServiceError:
				log.Println(v.Error())
			default:
				log.Println("no")
			}
		}
	}()
	panic(NewServiceError("错误信息 ok！"))
}
