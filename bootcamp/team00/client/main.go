package client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"randomaliens/internal/grpc/transmitter"
	"randomaliens/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	service *service.Service
}

func NewClient(service *service.Service) *Client {
	return &Client{
		service: service,
	}
}

func (c Client) StartTransmission(ctx context.Context) error {
	conn, err := grpc.DialContext(ctx, "localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	log.Println("Created grpc client")

	clienttt := transmitter.NewTransmitterClient(conn)
	recordStream, err := clienttt.StartTransmit(ctx, &transmitter.Empty{})

	if err != nil {
		log.Println("Server request failed")
		return err
	}

	log.Println("Started receiving stream...")
	var session *service.AnomalySession

	for {
		select {
		case <-ctx.Done():
			log.Println("Stream force interrupted")
			return c.service.Session.SaveSession(session)
		default:
			res, err := recordStream.Recv()

			if err != nil {
				errSS := c.service.Session.SaveSession(session)
				if errSS != nil {
					return errors.New(err.Error() + " another error occoured during saving: " + errSS.Error())
				}
				return err
			}

			if session == nil {
				_sess, err := c.service.Session.NewSession(*res.SessionId, 2)
				if err != nil {
					return err
				}
				session = _sess
			}

			if session.ReceivedRecordsCount > 50 {
				if session.IsAnomaly(res) {
					log.Printf("------ Found anomaly: id: %v, freq: %v, time: %v\n", *res.SessionId, *res.Frequency, *res.Timestamp)
					err := c.service.Session.SaveAnomaly(res, session)
					if err != nil {
						return err
					}

				}
			} else {
				fmt.Printf("Record:\nid: %s\nfrequency: %v\ntime: %s\n", *res.SessionId, *res.Frequency, *res.Timestamp)
				if err := session.Record(res); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
