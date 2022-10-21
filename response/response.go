package response

import (
	"github.com/go-resty/resty/v2"
	"time"
)

type Response struct {
	UUID          string
	Error         string
	Steps         []string
	Trace         Trace
	Status        string
	StatusCode    int
	Content       string
	ContentType   string
	ContentLength int64
	ReceivedAt    time.Time
}

func ProcessResp(resp *resty.Response, err error) Response {
	processResp := Response{}
	// 处理异常
	if err != nil {
		processResp.Error = err.Error()
		return processResp
	}
	// 处理响应
	processResp.Status = resp.Status()
	processResp.StatusCode = resp.StatusCode()
	processResp.ContentType = resp.Header().Get("content-type")
	processResp.ContentLength = resp.RawResponse.ContentLength
	processResp.Content = string(resp.Body())
	processResp.ReceivedAt = resp.ReceivedAt()
	// 处理 trace
	ti := resp.Request.TraceInfo()
	resTrace := Trace{
		DNSLookup:      ti.DNSLookup,
		ConnTime:       ti.ConnTime,
		TCPConnTime:    ti.TCPConnTime,
		TLSHandshake:   ti.TLSHandshake,
		ServerTime:     ti.ServerTime,
		ResponseTime:   ti.ResponseTime,
		TotalTime:      ti.TotalTime,
		IsConnReused:   ti.IsConnReused,
		IsConnWasIdle:  ti.IsConnWasIdle,
		ConnIdleTime:   ti.ConnIdleTime,
		RequestAttempt: ti.RequestAttempt,
		RemoteAddr:     ti.RemoteAddr,
	}
	processResp.Trace = resTrace
	return processResp
}
