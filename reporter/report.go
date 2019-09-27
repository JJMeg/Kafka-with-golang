package reporter

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"../error_const"
)

type KafkaCfg struct {
	Host  string `json:"host"`
	Topic string `json:"topic"`
}

type Reporter struct {
	Producer sarama.SyncProducer
	logger   *logrus.Logger
}

func NewReporter(cfg *KafkaCfg, log *logrus.Logger) *Reporter {
	reporter := &Reporter{
		logger: log,
	}

	reporter.setProducer(cfg)
	return reporter
}

func (reporter *Reporter) setProducer(cfg *KafkaCfg) {
	var broker = []string{cfg.Host}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(broker, config)
	if err != nil {
		reporter.logger.Errorf(error_const.InitProducerError, err)
	}

	reporter.Producer = producer
}

func (reporter *Reporter) DoReport(topic string, msg []byte) {
	reporter.do(topic, msg)
}

func (reporter *Reporter) do(topic string, msg []byte) {
	kafkaMsg := generateProducerMessage(topic, msg)
	_, _, err := reporter.Producer.SendMessage(kafkaMsg)
	if err != nil {
		reporter.logger.Errorf(error_const.ReportKafkaMsgError, err, string(msg))
	}
	reporter.logger.Infof(error_const.ReportKafkaMsgSuccess, string(msg))
}

func generateProducerMessage(topic string, message []byte) *sarama.ProducerMessage {
	return &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}
}
