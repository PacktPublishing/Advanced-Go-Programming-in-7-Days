package application

import (
	"github.com/Shopify/sarama"
	"log"
	"encoding/json"
	"fmt"
)

var (
	Topic   = "user-transactions"
	Partition int32 = 0
	queues = []string{"127.0.0.1:9092"}
	topics  = []string{Topic}

)

func newKafkaConfig() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.ChannelBufferSize = 1
	conf.Version = sarama.V0_10_2_1
	return conf
}

func NewKafkaSyncProducer() sarama.SyncProducer {
	kafka, err := sarama.NewSyncProducer(queues, newKafkaConfig())

	if err != nil {
		log.Fatalf("Kafka error: %s\n", err)
	}

	return kafka
}

func NewKafkaConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer(queues, newKafkaConfig())

	if err != nil {
		log.Fatalf("Kafka error: %s\n", err)
	}

	return consumer
}

func SubmitEvent(kafka sarama.SyncProducer, event interface{}) error {
	json, err := json.Marshal(event)

	if err != nil {
		return err
	}

	msgLog := &sarama.ProducerMessage{
		Topic: Topic,
		Value: sarama.StringEncoder(string(json)),
	}

	partition, offset, err := kafka.SendMessage(msgLog)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in partition %d, and offset %d\n", partition, offset)

	return nil
}

func ConsumeUserRegistrationEvents(consumer sarama.PartitionConsumer, handleEventType func(msgVal []byte, e interface{})) {
	var err error
	var msgItem []byte
	var log interface{}
	var logMap map[string]interface{}

	for {
		select {
		case err := <-consumer.Errors():
			fmt.Printf("Kafka error: %s\n", err)
		case msg := <-consumer.Messages():
			msgItem = msg.Value
			err = json.Unmarshal(msgItem, &log)
			if err != nil {
				fmt.Printf("Parsing error: %s", err)
			}
			logMap = log.(map[string]interface{})
			logType := logMap["Type"]
			handleEventType(msgItem, logType)
		}

	}
}
