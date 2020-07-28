package event

import (
	"fmt"
	"practise/router_cass_kakf/kafka"
)

const (
	//CRMTopic CRMTopic
	CRMTopic = "crmtopic"
)

/* var producer sarama.SyncProducer
var err error

func init() {
	producer, err = kafka.NewProducer()
	if err != nil {
		log.Fatal("error while creating producer", err)
	}
} */

//SendDataToKafka SendDataToKafka
func SendDataToKafka(byteArr []byte) {
	producer, err := kafka.NewProducer()
	if err != nil {
		fmt.Println("error while creating producer", err)
	}
	message := kafka.PrepareMessage(CRMTopic, string(byteArr))
	prtn, offs, err := producer.SendMessage(message)
	if err != nil {
		fmt.Println("error while sending message", err)
	}
	fmt.Printf("message sent on partition %v offset %v", prtn, offs)
}

//ReceiveDataFromKafka ReceiveDataFromKafka
func ReceiveDataFromKafka() {
	consumer, err := kafka.NewConsumer()
	if err != nil {
		fmt.Println("error while creating consumer", err)
	}
	kafka.Subscribe(CRMTopic, consumer)
}
