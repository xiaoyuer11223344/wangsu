// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/tea"
)

type IpInfoServiceRequest struct {
  // {'en':'The list of IP that needs to be querying is 20 times a single time.', 'zh_CN':'需要查询的IP列表，单次最大20个（联系技术支持可调上限）'}
  Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" require:"true" type:"Repeated"`
}

func (s IpInfoServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceRequest) GoString() string {
  return s.String()
}

func (s *IpInfoServiceRequest) SetIp(v []*string) *IpInfoServiceRequest {
  s.Ip = v
  return s
}

type IpInfoServiceResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result []*IpInfoServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s IpInfoServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceResponse) GoString() string {
  return s.String()
}

func (s *IpInfoServiceResponse) SetResult(v []*IpInfoServiceResponseResult) *IpInfoServiceResponse {
  s.Result = v
  return s
}

type IpInfoServiceResponseResult struct     {
  // {'en':'IP addresses', 'zh_CN':'IP地址'}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {'en':'Whether to network the our IP
  // 
  //         1.true is the node IP of our CDN
  // 
  //         2.false is not the node IP of the CDN', 'zh_CN':'是否我司CDN的IP
  //         1.true 是我司CDN的节点IP
  //         2.false &nbsp; 不是我司CDN的节点IP'}
  IsCdnIp *bool `json:"isCdnIp,omitempty" xml:"isCdnIp,omitempty" require:"true"`
  // {'en':'If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.', 'zh_CN':'归属国家地区；不是我司CDN的节点，不返回；如未规划的则返回未知。'}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {'en':'If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.', 'zh_CN':'归属省份；不是我司CDN的节点，不返回；如未规划的则返回未知；'}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {'en':'If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.', 'zh_CN':'归属城市；不是我司CDN的节点，不返回；如未规划的则返回未知；'}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {'en':'If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.', 'zh_CN':'归属运营商；不是我司CDN的节点，不返回；如未规划的则返回未知。'}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
}

func (s IpInfoServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceResponseResult) GoString() string {
  return s.String()
}

func (s *IpInfoServiceResponseResult) SetIp(v string) *IpInfoServiceResponseResult {
  s.Ip = &v
  return s
}

func (s *IpInfoServiceResponseResult) SetIsCdnIp(v bool) *IpInfoServiceResponseResult {
  s.IsCdnIp = &v
  return s
}

func (s *IpInfoServiceResponseResult) SetCountry(v string) *IpInfoServiceResponseResult {
  s.Country = &v
  return s
}

func (s *IpInfoServiceResponseResult) SetProvince(v string) *IpInfoServiceResponseResult {
  s.Province = &v
  return s
}

func (s *IpInfoServiceResponseResult) SetCity(v string) *IpInfoServiceResponseResult {
  s.City = &v
  return s
}

func (s *IpInfoServiceResponseResult) SetIsp(v string) *IpInfoServiceResponseResult {
  s.Isp = &v
  return s
}

type Paths struct {
}

func (s Paths) String() string {
  return tea.Prettify(s)
}

func (s Paths) GoString() string {
  return s.String()
}

type Parameters struct {
}

func (s Parameters) String() string {
  return tea.Prettify(s)
}

func (s Parameters) GoString() string {
  return s.String()
}

type RequestHeader struct {
}

func (s RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RequestHeader) GoString() string {
  return s.String()
}

type ResponseHeader struct {
}

func (s ResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ResponseHeader) GoString() string {
  return s.String()
}


