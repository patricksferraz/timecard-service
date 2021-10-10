package repository

import (
	"context"
	"fmt"

	"github.com/c-4u/timecard-service/domain/entity"
	"github.com/c-4u/timecard-service/infrastructure/db"
	"github.com/c-4u/timecard-service/infrastructure/external"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Repository struct {
	P *db.Postgres
	K *external.KafkaProducer
}

func NewRepository(db *db.Postgres, kafkaProducer *external.KafkaProducer) *Repository {
	return &Repository{
		P: db,
		K: kafkaProducer,
	}
}

func (r *Repository) SaveEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.P.Db.Save(employee).Error
	return err
}

func (r *Repository) CreateEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.P.Db.Create(employee).Error
	return err
}

func (r *Repository) FindEmployee(ctx context.Context, id string) (*entity.Employee, error) {
	var employee entity.Employee
	r.P.Db.Preload("Companies").First(&employee, "id = ?", id)

	if employee.ID == "" {
		return nil, fmt.Errorf("no employee found")
	}

	return &employee, nil
}

func (r *Repository) CreateCompany(ctx context.Context, company *entity.Company) error {
	err := r.P.Db.Create(company).Error
	return err
}

func (r *Repository) FindCompany(ctx context.Context, id string) (*entity.Company, error) {
	var company entity.Company
	r.P.Db.First(&company, "id = ?", id)

	if company.ID == "" {
		return nil, fmt.Errorf("no company found")
	}

	return &company, nil
}

func (r *Repository) CreateEvent(ctx context.Context, event *entity.Event) error {
	err := r.P.Db.Create(event).Error
	return err
}

func (r *Repository) FindEvent(ctx context.Context, id string) (*entity.Event, error) {
	var event entity.Event
	r.P.Db.First(&event, "id = ?", id)

	if event.ID == "" {
		return nil, fmt.Errorf("no event found")
	}

	return &event, nil
}

func (r *Repository) SaveEvent(ctx context.Context, event *entity.Event) error {
	err := r.P.Db.Save(event).Error
	return err
}

func (r *Repository) PublishEvent(ctx context.Context, msg, topic, key string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
		Key:            []byte(key),
	}
	err := r.K.Producer.Produce(message, r.K.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}
