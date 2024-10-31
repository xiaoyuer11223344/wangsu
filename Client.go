package main

import (
	"fmt"
	"github.com/xiaoyuer11223344/wangsu/api/client"
	"github.com/xiaoyuer11223344/wangsu/common/auth"
)

func main() {

	ipInfoServiceRequest := client.IpInfoServiceRequest{}
	var subipInfoServiceRequest0 = "58.220.76.28"
	ipInfoServiceRequest.SetIp([]*string{&subipInfoServiceRequest0})

	var config auth.AkskConfig
	config.AccessKey = "ak"
	config.SecretKey = "sk"
	config.EndPoint = "open.chinanetcenter.com"
	config.Uri = "/api/tools/ip-info"
	config.Method = "POST"
	response := auth.Invoke(config, ipInfoServiceRequest.String())
	fmt.Printf("response body is %#v\n", response)
}
