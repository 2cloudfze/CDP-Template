package messageManager

import (
	"encoding/json"
	"fmt"
)

type TopicBody struct {
}

func BytesToStruct(message []byte) (TopicBody, error) {
	topic := TopicBody{}
	err := json.Unmarshal([]byte(message), &topic)
	if err != nil {
		fmt.Println("Error:", err)
		return topic, err
	}
	return topic, nil
}

func (tb *TopicBody) ProcessTopic() {

}
