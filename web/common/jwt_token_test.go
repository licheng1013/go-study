package common

import (
	"log"
	"testing"
)

func TestJwtToken(t *testing.T) {
	token, err := CreateToken("1")
	if err != nil {
		log.Println("错误:" + err.Error())
	}
	log.Println(token)
	parseToken, err := ParseToken(token)
	if err != nil {
		log.Println("错误:" + err.Error())
	}
	log.Println(parseToken)
}
