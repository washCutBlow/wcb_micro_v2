package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	wcb_micro_v2 "wcb_micro_v2/proto/wcb_micro_v2"
)

type Wcb_micro_v2 struct{}

func (e *Wcb_micro_v2) Handle(ctx context.Context, msg *wcb_micro_v2.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *wcb_micro_v2.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
