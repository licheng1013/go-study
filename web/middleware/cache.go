package middleware

//import (
//	"github.com/gin-gonic/gin"
//	"go-util/util"
//	"log"
//	"net/http"
//	"time"
//	"web/common"
//	"web/config"
//)
//
//var CachePath []string
//
//// 需要配合Redis进行使用！
//func init() {
//	//initCache()
//}
//
//// Cache Redis缓存中间件
//func Cache() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		path := util.GinUrlPath()
//		for _, i := range CachePath {
//			//需要缓存该返回结果！
//			if i == path {
//				cacheStr := config.RedisGetString(path)
//				log.Println("触发缓存:" + path)
//				if cacheStr == "" {
//					c.Next()
//					return
//				}
//				var r common.Result
//				util.JsonToObj(cacheStr, &r)
//				c.JSON(http.StatusOK, r)
//				c.Abort() //终止访问
//				break
//			}
//		}
//	}
//}
//
//// CacheData 添加需要缓存的路径,默认30秒
//func CacheData(k, v string) {
//	isExistPath := false
//	for _, s := range CachePath {
//		if s == k {
//			log.Println("触发缓存：" + k)
//			isExistPath = true
//			break
//		}
//		isExistPath = false
//	}
//	if !isExistPath {
//		CachePath = append(CachePath, k)
//	}
//	log.Println(CachePath)
//	config.RedisSetExpiration(k, v, 30*time.Second)
//}
//
//// CachePathUrlData 根据路径进行缓存 /file/list?xx=aa => /file/list
//func CachePathUrlData(v string) {
//	CacheData(util.GinUrlPath(), v)
//}
//
//// GetParamsUrlPahData 根据路径+参数进行缓存 /file/list?xx=aa => /file/list?xx=aa
//func GetParamsUrlPahData(v string) {
//	CacheData(util.GinParamsUrlPah(), v)
//}
