package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"wcb_micro_v2/handler"
	"wcb_micro_v2/proto/wcb_micro_v2"
	"wcb_micro_v2/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.wcb_micro_v2"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	wcb_micro_v2.RegisterWcbMicroV2Handler(service.Server(), new(handler.Wcb_micro_v2))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.wcb_micro_v2", service.Server(), new(subscriber.Wcb_micro_v2))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
