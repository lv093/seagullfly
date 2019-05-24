package consumer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"sync"
	"seagullfly/services/handler"
	"seagullfly/models"
)

type KafkaConsumer struct {
	consumer sarama.Consumer
}

func NewKafkaConsumer(cfg []string) (*KafkaConsumer, error) {
	svc := new(KafkaConsumer)

	consumer, err := sarama.NewConsumer(cfg, nil)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	svc.consumer = consumer

	return svc, nil
}

func (this *KafkaConsumer) Consume(topic string) {

	partitions, err := this.consumer.Partitions(topic)

	if err != nil {
		logs.Error(err)
		return
	}

	var wg sync.WaitGroup

	for partition := range partitions {
		pc, err := this.consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			logs.Error(err)
			return
		}

		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Topic: %s, Partition:%d, Offset:%d, Key:%s, Value:%s\n", topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				this.Handle(topic, msg.Value)
			}
		}(pc)
		wg.Wait()
	}
	this.consumer.Close()
}

func (this *KafkaConsumer) Handle(topic string, message []byte) bool {
	if topic == "article" {
		handler := handler.GetArticleDataHandler()
		msg := new(models.ArticlesOrm)

		err := handler.Handle(msg)
		if err == nil {
			return true
		}
	}
	return false
}
