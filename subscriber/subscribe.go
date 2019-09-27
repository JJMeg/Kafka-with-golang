package subscriber

import (
	"../error_const"
	"../reporter"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	Consumer sarama.Consumer
	logger   *logrus.Logger
}

func NewSubscriber(cfg *reporter.KafkaCfg, log *logrus.Logger) *Subscriber {
	subscriber := &Subscriber{
		logger: log,
	}
	subscriber.setConsumer(cfg)
	return subscriber
}

func (subscriber *Subscriber) setConsumer(cfg *reporter.KafkaCfg) {
	consumer, err := sarama.NewConsumer([]string{cfg.Host}, nil)
	if err != nil {
		panic(err)
	}
	subscriber.Consumer = consumer
}

func (subscriber *Subscriber) Consume(topic string, ch chan string) {
	defer func() {
		if err := subscriber.Consumer.Close(); err != nil {
			subscriber.logger.Errorf(error_const.SubcriberCloseConsumerError, err)
		}
	}()

	partitionList, err := subscriber.Consumer.Partitions(topic)
	if err != nil {
		subscriber.logger.Errorf(error_const.SubScriberGetPartitionsError, err, topic)
	}

	for _, partition := range partitionList {
		pc, _ := subscriber.Consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				subscriber.logger.Infof(messageReceived(message))
				ch <- messageReceived(message)
			}
		}(pc)
	}
}

func messageReceived(message *sarama.ConsumerMessage) string {
	return string(message.Value)
}
