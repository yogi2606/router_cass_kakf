package event

import (
	"fmt"
	"practise/router_cass_kakf/kafka"

	"github.com/Shopify/sarama"
)

const (
	//CRMTopic CRMTopic
	CRMTopic = "crmtopic"
)

var producer sarama.SyncProducer
var consumer sarama.Consumer
var err error

func init() {
	consumer, err = kafka.NewConsumer()
	if err != nil {
		fmt.Println("error while creating consumer", err)
	}
	producer, err = kafka.NewProducer()
	if err != nil {
		fmt.Println("error while creating producer", err)
	}
}

//KafkaImpl KafkaImpl
type KafkaImpl struct {
}

//SendDataToKafka SendDataToKafka
func (k KafkaImpl) SendDataToKafka(byteArr []byte) {
	/* producer, err := kafka.NewProducer()
	if err != nil {
		fmt.Println("error while creating producer", err)
	} */
	message := kafka.PrepareMessage(CRMTopic, string(byteArr))
	prtn, offs, err := producer.SendMessage(message)
	if err != nil {
		fmt.Println("error while sending message", err)
	}
	fmt.Printf("message sent on partition %v offset %v", prtn, offs)
}

//ReceiveDataFromKafka ReceiveDataFromKafka
func (k KafkaImpl) ReceiveDataFromKafka() {
	/* consumer, err := kafka.NewConsumer()
	if err != nil {
		fmt.Println("error while creating consumer", err)
	} */
	kafka.Subscribe(CRMTopic, consumer)
}
