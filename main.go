package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"net"
	"time"
	"wechatpro_client/gen-go/wechat"
)

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(10240)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport, err := thrift.NewTSocket(net.JoinHostPort("1.15.72.208", "9876"))
	if err != nil {
		panic(err)
	}
	useTransport, err := transportFactory.GetTransport(transport)
	if useTransport == nil || err != nil {
		panic(err)
	}
	client := wechat.NewWechatClientFactory(useTransport, protocolFactory)
	if err := useTransport.Open(); err != nil {
		panic(err)
	}
	defer useTransport.Close()

	groups, err :=client.FetchGroups(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(groups)

	res, err := client.Send(context.Background(), 1, fmt.Sprintf("现在是北京时间%v", time.Now().String()))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
