package request

import (
	"github.com/go-resty/resty/v2"
	"github.com/open-okapi/data-driven-proxies/enums"
	"github.com/open-okapi/data-driven-proxies/sub"
	http2 "net/http"
	"strconv"
)

// Assemble 整合 http 请求
func Assemble(r *resty.Request, meta Meta) {
	// 处理 http 基础数据
	var http = meta.Http
	r.URL = http.Schema + "://" + http.Host + ":" + strconv.Itoa(http.Port) + http.Path
	r.Method = http.Method
	// base auth
	if http.BaseAuth != (sub.AuthPair{}) {
		r.SetBasicAuth(http.BaseAuth.Username, http.BaseAuth.Password)
	}
	// 处理 header
	for _, v := range meta.Header {
		r.SetHeader(v.Key, v.Value)
	}
	// 处理 query parameter
	for _, v := range meta.Query {
		r.SetQueryParam(v.Key, v.Value)
	}
	// 处理 body
	var entity = meta.Entity
	var pairs = meta.Pairs
	if entity != (sub.Entity{}) {
		switch entity.Type {
		case enums.Form:
			var pairTextMap = map[string]string{}
			for _, v := range pairs {
				// 判断值或者文件
				switch v.ValueType {
				case enums.PairTypeFile:
					var fileResp, _ = http2.Get(v.Value)
					r.SetFileReader(v.Key, v.Key, fileResp.Body)
				case enums.PairTypeText:
					pairTextMap[v.Key] = v.Value
				}
			}
		case enums.RawJson:
			r.SetBody(entity.Data)
		default:
			r.SetBody(entity.Data)
		}
	}
}
