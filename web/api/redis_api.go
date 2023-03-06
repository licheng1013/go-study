package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web/common"
	"web/config"
)

// 初始化方法
func init() {
	//fmt.Println("初始化user")
	redis := common.App.Group("/redis")
	redis.GET("/test", redisTest)
}

// 测试redis锁
func redisTest(c *gin.Context) {
	lock := config.RedisLock("test", 10*time.Second)
	if lock {
		k := "goods1"
		getInt := config.RedisGetInt(k)
		if getInt <= 0 {
			fmt.Println("已卖光")
			c.JSON(http.StatusOK, common.OkMsg("已卖光"))
		} else {
			fmt.Println("剩余=", getInt)
			getInt -= 1
			config.RedisSet(k, getInt)
		}
		config.RedisDel("test") //释放锁
	} else {
		fmt.Println("上锁失败")
	}
	c.JSON(http.StatusOK, common.OkMsg("测试"))
}
