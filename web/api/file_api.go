package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"web/common"
)

// 初始化方法
func init() {
	v := FileApi{}
	file := common.App.Group("/file")
	file.POST("/upload", v.upload)
}

type FileApi struct {
}

// 上传文件
func (FileApi) upload(c *gin.Context) {
	//单文件
	file, _ := c.FormFile("file")
	// 上传文件至指定的完整文件路径
	err := c.SaveUploadedFile(file, file.Filename)
	if err != nil {
		log.Panic(err)
	}
	c.JSON(http.StatusOK, common.OkMsg(file.Filename+"文件上传成功！"))
}
