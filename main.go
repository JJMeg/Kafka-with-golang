package main

import (
	"./reporter"
	"./subscriber"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	config = `{
	"host": "localhost:9092",
    "topic": "kafka_test"
  }`
)

func main() {
	var cfg reporter.KafkaCfg
	json.Unmarshal([]byte(config), &cfg)
	log := &logrus.Logger{Out: os.Stdout}

	//reporter
	producer := reporter.NewReporter(&cfg, log)
	//subscriber
	consumer := subscriber.NewSubscriber(&cfg, log)
	message := "Hello Kafka World."

	ch := make(chan string)
	consumer.Consume(cfg.Topic, ch)
	producer.DoReport(cfg.Topic, []byte(message))

	select {
	case msg := <-ch:
		fmt.Println("Got msg: ", msg)
		break
	}
}
