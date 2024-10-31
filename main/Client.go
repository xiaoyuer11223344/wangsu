package main

import (
	"fmt"
	_ "github.com/xiaoyuer11223344/wangsu/api/client"
	"github.com/xiaoyuer11223344/wangsu/common/auth"
	"log"
)

func main() {

	var config auth.AkskConfig
	config.AccessKey = "{accessKey}"
	config.SecretKey = "{secretKey}"
	config.EndPoint = "{endPoint}"
	config.Uri = "/api/test/example"
	config.Method = "POST"
	var requestBody string = "120.79.66.58"
	response, err := auth.Invoke(config, requestBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response body is %#v\n", response)
}
