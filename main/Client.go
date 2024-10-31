package main

import (
	"fmt"
	_ "github.com/xiaoyuer11223344/wangsu/api/client"
	"github.com/xiaoyuer11223344/wangsu/common/auth"
)

func main() {

	var config auth.AkskConfig
	config.AccessKey = "{accessKey}"
	config.SecretKey = "{secretKey}"
	config.EndPoint = "{endPoint}"
	config.Uri = "/api/test/example"
	config.Method = "POST"
	var requestBody string = "120.79.66.58"
	response := auth.Invoke(config, requestBody)
	fmt.Printf("response body is %#v\n", response)
}
