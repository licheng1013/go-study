package api

import (
	"log"
	"web/middleware"
)

func Init() {
	middleware.Init()
	log.Println("初始化所有Api")
}
