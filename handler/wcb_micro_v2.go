package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	wcb_micro_v2 "wcb_micro_v2/proto/wcb_micro_v2"
)

type Wcb_micro_v2 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Wcb_micro_v2) Call(ctx context.Context, req *wcb_micro_v2.Request, rsp *wcb_micro_v2.Response) error {
	log.Info("Received Wcb_micro_v2.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Wcb_micro_v2) Stream(ctx context.Context, req *wcb_micro_v2.StreamingRequest, stream wcb_micro_v2.WcbMicroV2_StreamStream) error {
	log.Infof("Received Wcb_micro_v2.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&wcb_micro_v2.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Wcb_micro_v2) PingPong(ctx context.Context, stream wcb_micro_v2.WcbMicroV2_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&wcb_micro_v2.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
