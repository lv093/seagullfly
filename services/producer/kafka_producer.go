package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

var kafkaProducer *KafkaProducer

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func GetKafkaProducer() *KafkaProducer {
	if kafkaProducer == nil {
		svc := new(KafkaProducer)
		config := sarama.NewConfig()
		config.Producer.Return.Successes = true
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Partitioner = sarama.NewRandomPartitioner

		address := beego.AppConfig.String("producer.soccer.kafka.address")
		producer, err := sarama.NewSyncProducer(strings.Split(address, ","), config)

		if err != nil {
			logs.Error(err)
			return nil
		}

		svc.producer = producer

		kafkaProducer = svc
	}

	return kafkaProducer
}

func (this *KafkaProducer) Produce(topic string, message string) {

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	msg.Value = sarama.ByteEncoder(message)
	paritition, offset, err := this.producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Send Message Fail", topic, " | ", message)
	}

	fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
}
