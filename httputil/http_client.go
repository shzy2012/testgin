package httputil

import (
	"net"
	"net/http"
	"time"
)

func NewHttpClient(idleConns, perHost int, timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: time.Minute * 1, //设置超时时间,默认0不设置超时时间
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second, //限制建立TCP连接的时间
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second, //限制 TLS握手的时间
			ResponseHeaderTimeout: 10 * time.Second, //限制读取response header的时间
			ExpectContinueTimeout: 1 * time.Second,  //限制client在发送包含 Expect: 100-continue的header到收到继续发送body的response之间的时间等待。
			MaxIdleConns:          idleConns,        //连接池对所有host的最大连接数量，默认无穷大
			MaxConnsPerHost:       perHost,          //连接池对每个host的最大连接数量。
			IdleConnTimeout:       timeout,          //how long an idle connection is kept in the connection pool.
		},
	}
}
