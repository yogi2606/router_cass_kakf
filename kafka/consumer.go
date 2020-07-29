package kafka

import (
	"encoding/json"
	"fmt"
	"practise/router_cass_kakf/model"

	"github.com/Shopify/sarama"
)

//NewConsumer NewConsumer
func NewConsumer() (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer(brokers, config)
	return consumer, err
}

//Subscribe Subscribe
func Subscribe(topic string, consumer sarama.Consumer) {
	partitionList, err := consumer.Partitions(topic) //get all partitions on the given topic
	if err != nil {
		fmt.Println("Error retrieving partitionList ", err)
	}
	initialOffset := sarama.OffsetOldest //get offset for the oldest message on the topic

	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, partition, initialOffset)
		if err != nil {
			fmt.Println("ConsumePartition", err)
		}
		go func(pc sarama.PartitionConsumer) {
			defer func() {
				r := recover()
				if r != nil {
					fmt.Println(r)
				}
			}()
			for {
				select {
				case err := <-pc.Errors():
					fmt.Println("PartitionConsumer", err)
				case msg := <-pc.Messages():
					messageReceived(msg)
				}
			}
		}(pc)
	}
}

func messageReceived(message *sarama.ConsumerMessage) {
	customer := model.Customer{}
	err := json.Unmarshal(message.Value, &customer)
	if err != nil {
		fmt.Println("Unmarshal error", err)
	}
	fmt.Println("data received from kakfka ", customer)
}
