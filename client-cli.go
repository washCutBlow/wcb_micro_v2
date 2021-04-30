package main

import (
	"context"
	cli "github.com/micro/go-micro/v2/client/grpc"
	"log"
	"wcb_micro_v2/proto/wcb_micro_v2"
)

func main() {
	testCallFunc()
}

func testCallFunc() {
	wcbService := wcb_micro_v2.NewWcbMicroV2Service("go.micro.service.wcb_micro_v2", cli.NewClient())
	//  默认生成的服务中自带三个接口
	resp,err := wcbService.Call(context.Background(),&wcb_micro_v2.Request{
		Name:"hahhhh",
	})
	if err != nil {
		log.Panic("cal func",err)
	}
	log.Println("call func success", resp.Msg)
}