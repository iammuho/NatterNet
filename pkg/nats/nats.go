// Package nats is a module for nats
package nats

import (
	"github.com/nats-io/nats.go"
)

type server struct {
	conn             *nats.Conn
	jetstreamContext nats.JetStreamContext
}

// NewNats creates a new nats connection
func NewNats(opts ...Option) (NatsContext, error) {
	natsOptions := NatsOptions{}
	for _, o := range opts {
		o(&natsOptions)
	}

	nc, _ := nats.Connect(natsOptions.URL)
	js, err := nc.JetStream() // Returns JetStreamContext

	if err != nil {
		return nil, err
	}

	return &server{
		conn:             nc,
		jetstreamContext: js,
	}, nil
}

// GetConn returns the nats connection
func (s *server) GetConn() *nats.Conn {
	return s.conn
}

// GetJetStreamContext returns the nats jetstream context
func (s *server) GetJetStreamContext() nats.JetStreamContext {
	return s.jetstreamContext
}

// CreateStream creates a new stream
func (s *server) CreateStream(streamName string, subject string) error {
	_, err := s.jetstreamContext.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{subject},
	})

	return err
}

// Subscribe subscribes to a stream
func (s *server) Subscribe(subject string, handler func(msg *nats.Msg) error) error {
	_, err := s.jetstreamContext.Subscribe(subject, func(msg *nats.Msg) {
		handler(msg)
	}, nats.Durable("monitor"), nats.ManualAck())

	return err
}
