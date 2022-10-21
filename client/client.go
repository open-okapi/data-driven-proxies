package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/open-okapi/data-driven-proxies/config"
	"github.com/open-okapi/data-driven-proxies/request"
	"github.com/open-okapi/data-driven-proxies/response"
	"sync"
)

var client *resty.Client

var once sync.Once

func getClient() *resty.Client {
	once.Do(func() {
		client = resty.New()
		config.InitConfig(client)
	})
	return client
}

func Execute(meta request.Meta) response.Response {
	// 获取 client
	var r = getClient().R()
	// 组装请求
	request.Assemble(r, meta)
	// 执行请求
	resp, err := r.Send()
	// 处理响应
	processResp := response.ProcessResp(resp, err)
	// 关联响应 UUID
	processResp.UUID = meta.UUID
	// 响应结果
	return processResp
}
