# Kafka With Golang
分布式消息平台，用Scala实现，本项目使用Go语言实现简单的消息生产和消费

# 常用的Kafka客户端

| project | feature | weakness |
| --- | --- | --- |
| Shopify/sarama | 最受欢迎 | 集群式消费难实现，不支持Context |
| bsm/sarama-cluster | 基于sarama补充集群式消费 | 不支持Context |
| confluentinc/confluent-kafka-go | | 依赖C语言库，不支持Context |
| lovoo/goka | 依赖于sarama | 不支持Context |
| segmentio/kafka-go | 同时支持集群模式，易与软件交互 | 未正式发布，支持Context |
