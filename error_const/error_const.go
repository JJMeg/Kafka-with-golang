package error_const

const (
	InitReporterSuccess = "Init reporter successfully."
	InitProducerError   = "Init producer error occured: %v"

	ReportKafkaMsgError   = "Reporting KafkaMessage error occured: %v , the message is: [ %s ]"
	ReportKafkaMsgSuccess = "Reporting KafkaMessage successfully, the message is: [ %s ]"

	SubcriberCloseConsumerError  = "Subscriber close consumer error occured: %v"
	SubScriberGetPartitionsError = "Subscriber get topic partitions error occured: %v, the topic is [ %s ]"
)
