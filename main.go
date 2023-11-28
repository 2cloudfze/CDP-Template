package main

import (
	kafkaConsumer "cdpTemplate/src/kafka_consumer"
)

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	topic := "MONGO"

	err := kafkaConsumer.SubscribeToSingleTopic(topic)

	errorHandler(err)

}
