package kafkautils

import (
	apiconfig "dmp-api/api/api_config"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// a simple publishing function
func PublishToKafka(topic string, data []byte) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": apiconfig.GetFullKafkaAddress()})
	if err != nil {
		return err
	}

	defer p.Close()

	deliveryChan := make(chan kafka.Event)

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}
