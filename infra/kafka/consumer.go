package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
import "log"
import "os"

type KafkaConsumer struct{
	MsgChan chan *ckafka.Message
}

func(k *KafkaConsumer) Consume(){
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv(key: "KafkaBootstrapServers"),
		"group.id": os.Getenv(key: "KafkaConsumerGroupId"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf(format: "error consuming kafka message:" err.Error())
	}

	topics := []string{
		os.Getenv(key: "KafkaReadTopic")
	}
	c.SubscribeTopics(topics, nil)
}