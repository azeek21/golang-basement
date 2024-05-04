package handlers

import (
	context "golang.org/x/net/context"
	"randomaliens/server/pkgs/services/transmitter"
	"time"
)

type Handlers struct {
}

func (h *Handlers) StartTransmit(ctx context.Context, in *transmitter.Empty) (*transmitter.Record, error) {
	id := "session_id"
	frequency := 1.23
	curTime := time.Now().Format("DD.MM.YYYY-HH:mm:ss")
	return &transmitter.Record{
		SessionId: &id,
		Frequency: &frequency,
		Timestamp: &curTime,
	}, nil
}
