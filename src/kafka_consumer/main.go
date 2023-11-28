package kafkaConsumer

import (
	messageManager "cdpTemplate/src/message_manager"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func consumeToProducer() (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
		"group.id":          os.Getenv("GROUP_ID"),
		"auto.offset.reset": "smallest",
	},
	)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func SubscribeToSingleTopic(topic string) error {
	consumer, err := consumeToProducer()
	if err != nil {
		return err
	}
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		return err
	}
	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			topic, err := messageManager.BytesToStruct(e.Value)

			if err != nil {
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
			topic.ProcessTopic()

		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
		}
	}

}

func SubscribeToMultipleTopic(topics []string) error {
	consumer, err := consumeToProducer()
	if err != nil {
		return err
	}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}
	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			topic, err := messageManager.BytesToStruct(e.Value)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
			topic.ProcessTopic()
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
		}
	}

}
