package kafka

import (
	"testing"
)

func TestKafkaPublisher_WriteOnce(t *testing.T) {
	kp := NewKafkaPublisher("localhost:19092", "topic-1")

	err := kp.WriteOnce([]byte("key-1"), []byte("value-1"))

	if err != nil {
		t.Errorf("failed to write message: %v", err)
	}

}
