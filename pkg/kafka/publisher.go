package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Publisher struct {
	Address string
	Topic   string
}

func (kp *Publisher) WriteOnce(key, value []byte) error {
	kw := NewKafkaWriter(kp.Address, kp.Topic)
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}
	return kw.WriteMessages(context.Background(), msg)
}

func NewKafkaPublisher(address string, topic string) *Publisher {
	return &Publisher{
		Address: address,
		Topic:   topic,
	}
}

func NewKafkaWriter(address string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(address),
		Topic:        topic,
		MaxAttempts:  10, // default: 10
		RequiredAcks: kafka.RequireAll,
		Async:        false,
		Logger:       nil,
		ErrorLogger:  nil,
	}
}
