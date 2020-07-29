package event

import "sync"

//KafkaInterface KafkaInterface
type KafkaInterface interface {
	SendDataToKafka(byteArr []byte)
	ReceiveDataFromKafka()
}

var (
	kafkaImp KafkaImpl
	once     sync.Once
)

//New returns singleton object of UploadService
func New() KafkaImpl {
	once.Do(func() {
		kafkaImp = KafkaImpl{}
	})
	return kafkaImp
}
