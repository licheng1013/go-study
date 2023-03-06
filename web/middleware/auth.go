package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"web/common"
)

// Authorization 请求头key
var Authorization = "Authorization"

// Auth 认证拦截器
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//保存到上下文中
		path := c.Request.RequestURI // TODO 此处可能需要测试
		//排除路径处理
		for _, item := range common.AppConfig.ExcludePath {
			// 排除路径例如： /test/** , 排除匹配 /test/ 的所有拦截
			var ePath = item
			//检查是否存在**号
			if strings.HasSuffix(ePath, "**") || path == ePath {
				var prefix = strings.ReplaceAll(ePath, "*", "")
				if strings.HasPrefix(path, prefix) {
					return
				}
			}
		}
		//登入认证
		query := c.GetHeader(Authorization)
		if query == "" {
			c.JSON(http.StatusOK, common.Fail("没有登入!"))
			c.Abort() //终止访问
			return
		}
		c.Next()
	}
}
