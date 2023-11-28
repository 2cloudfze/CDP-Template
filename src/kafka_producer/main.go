package kafkaProducer

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ProducerManager struct {
	producer     *kafka.Producer
	topic        string
	deliveryChan chan kafka.Event
}

var Producer ProducerManager
var FailedProducer ProducerManager

func init() {
	p, err := CreateKafkaProducer()
	if err != nil {
		panic(err)
	}
	Producer = *NewProducerManager(p, "DATPROC")
	FailedProducer = *NewProducerManager(p, "FAILED")
}

func NewProducerManager(p *kafka.Producer, topic string) *ProducerManager {

	return &ProducerManager{
		producer:     p,
		topic:        topic,
		deliveryChan: make(chan kafka.Event, 10000),
	}
}

func CreateKafkaProducer() (*kafka.Producer, error) {
	hostname := os.Getenv("KAFKA_BROKER")
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": hostname,
		"client.id":         os.Getenv("CLIENT_ID"),
		"acks":              "all",
	})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (pm *ProducerManager) PlaceMessage(body string) error {
	var (
		payload = []byte(body)
	)
	err := pm.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &pm.topic, Partition: kafka.PartitionAny},
		Value:          payload,
	},
		pm.deliveryChan,
	)
	if err != nil {
		return err
	}
	<-pm.deliveryChan
	return nil
}
