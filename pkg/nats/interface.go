package nats

//go:generate mockgen -destination=mocks/mock_nats_contexter.go -package=nats_mock -source=interface.go

import "github.com/nats-io/nats.go"

type NatsContext interface {
	// GetConn returns the nats connection
	GetConn() *nats.Conn

	// GetJetStreamContext returns the nats jetstream context
	GetJetStreamContext() nats.JetStreamContext

	// CreateStream creates a new stream
	CreateStream(streamName string, subject string) error

	// Subscribe subscribes to a stream
	Subscribe(subject string, handler func(msg *nats.Msg) error) error
}
