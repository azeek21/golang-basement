package handler

import (
	"log"
	"randomaliens/internal/grpc/transmitter"
	"time"
)

func (h Handler) StartTransmit(_ *transmitter.Empty, srv transmitter.Transmitter_StartTransmitServer) error {
	log.Println("Transmitter StartTransmit: started")
	mean := h.service.Generator.Mean(-10, 10)
	sd := h.service.Generator.Deviation(0.3, 1.5)
	sessionId := h.service.Generator.Uuid()

	for {
		select {
		case <-srv.Context().Done():
			log.Println("Transmitter StartTransmit: Context cancelled, returning...")
			return nil
		case <-time.After(100 * time.Millisecond):
			{
				record := transmitter.Record{}
				tstamp := time.Now().String()
				frequency := h.service.Generator.Frequency(mean, sd)
				record.Frequency = &frequency
				record.SessionId = &sessionId
				record.Timestamp = &tstamp
				if err := srv.Send(&record); err != nil {
					log.Printf("Transmitter StartTransmit: send error %v", err)
					return err
				}
				log.Printf("Transmitter StartTransmit: record sent. Frequency: %v\n", record.GetFrequency())
			}
		}

	}

	log.Println("Transmitter StartTransmit: ended")
	return nil
}
