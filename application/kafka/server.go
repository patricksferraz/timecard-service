package kafka

import (
	"fmt"

	"github.com/c-4u/timecard-service/domain/service"
	"github.com/c-4u/timecard-service/infrastructure/db"
	"github.com/c-4u/timecard-service/infrastructure/external"
	"github.com/c-4u/timecard-service/infrastructure/repository"
)

func StartKafkaServer(database *db.Postgres, kafkaConsumer *external.KafkaConsumer, kafkaProducer *external.KafkaProducer) {
	repository := repository.NewRepository(database, kafkaProducer)
	service := service.NewService(repository)

	fmt.Println("kafka pocessor has been started")
	processor := NewKafkaProcessor(service, kafkaConsumer, kafkaProducer)
	processor.Consume()
}
