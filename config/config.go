package config

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"time"
)

func InitConfig(client *resty.Client) {
	client.SetAllowGetMethodPayload(true)
	client.SetDebug(true)
	client.SetRetryCount(3)
	client.SetTimeout(30 * time.Second)
	client.SetRetryMaxWaitTime(5 * time.Second)
	client.EnableTrace()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
}
