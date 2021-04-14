package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"net"
	"wechatpro_client/gen-go/wechat"
)

var host = "1.15.72.208"
//var host = "127.0.0.1"

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(10240)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport, err := thrift.NewTSocket(net.JoinHostPort(host, "9876"))
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
	for {
		fmt.Println("下面是所有的群组 请选择要发送的群组序号")
		for i := 0; i < len(groups); i++ {
			fmt.Printf("%v: %v\n", i+1, groups[i].GroupName)
		}
		var num int
		_, err = fmt.Scanln(&num)
		if err != nil {
			panic(err)
		}
		fmt.Println("请输入你要发送的消息")
		var content string
		_, err = fmt.Scanln(&content)
		if err != nil {
			panic(err)
		}
		res, err := client.Send(context.Background(), int32(num-1), content)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
}
