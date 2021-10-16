package kafka

import (
	"context"
	"fmt"

	"github.com/c-4u/timecard-service/application/kafka/schema"
	"github.com/c-4u/timecard-service/domain/entity"
	"github.com/c-4u/timecard-service/domain/service"
	"github.com/c-4u/timecard-service/infrastructure/external"
	"github.com/c-4u/timecard-service/infrastructure/external/topic"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProcessor struct {
	Service *service.Service
	Kc      *external.KafkaConsumer
	Kp      *external.KafkaProducer
}

func NewKafkaProcessor(service *service.Service, kafkaConsumer *external.KafkaConsumer, kafkaProducer *external.KafkaProducer) *KafkaProcessor {
	return &KafkaProcessor{
		Service: service,
		Kc:      kafkaConsumer,
		Kp:      kafkaProducer,
	}
}

func (p *KafkaProcessor) Consume() {
	p.Kc.Consumer.SubscribeTopics(p.Kc.ConsumerTopics, nil)
	for {
		msg, err := p.Kc.Consumer.ReadMessage(-1)
		if err == nil {
			// fmt.Println(string(msg.Value))
			p.processMessage(msg)
		}
	}
}

func (p *KafkaProcessor) processEvent(msg *ckafka.Message) (*entity.Event, error) {
	event := &schema.Event{}
	err := event.ParseJson(msg.Value)
	if err != nil {
		return nil, err
	}

	e, err := p.Service.ProcessEvent(context.TODO(), event.ID, msg.String())
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (p *KafkaProcessor) completeEvent(event *entity.Event) error {
	err := p.Service.CompleteEvent(context.TODO(), event.ID)
	return err
}

func (p *KafkaProcessor) processMessage(msg *ckafka.Message) {

	event, err := p.processEvent(msg)
	if err != nil {
		fmt.Println("event processing error ", err)
	}

	switch _topic := *msg.TopicPartition.Topic; _topic {
	case topic.NEW_EMPLOYEE:
		// TODO: add fault tolerance
		err := p.createEmployee(msg)
		if err != nil {
			fmt.Println("creation error ", err)
			p.retry(msg)
		}
	case topic.NEW_COMPANY:
		err := p.createCompany(msg)
		if err != nil {
			fmt.Println("creation error ", err)
			p.retry(msg)
		}
	case topic.ADD_EMPLOYEE_TO_COMPANY:
		err := p.addEmployeeToCompany(msg)
		if err != nil {
			fmt.Println("addition error ", err)
			p.retry(msg)
		}
	default:
		fmt.Println("not a valid topic", string(msg.Value))
	}

	err = p.completeEvent(event)
	if err != nil {
		fmt.Println("event completion error ", err)
	}
}

func (p *KafkaProcessor) retry(msg *ckafka.Message) error {
	err := p.Kc.Consumer.Seek(ckafka.TopicPartition{
		Topic:     msg.TopicPartition.Topic,
		Partition: msg.TopicPartition.Partition,
		Offset:    msg.TopicPartition.Offset,
	}, -1)

	return err
}

func (p *KafkaProcessor) createEmployee(msg *ckafka.Message) error {
	employeeEvent := &schema.EmployeeEvent{}
	err := employeeEvent.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	err = p.Service.CreateEmployee(context.TODO(), employeeEvent.Employee.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) createCompany(msg *ckafka.Message) error {
	companyEvent := &schema.CompanyEvent{}
	err := companyEvent.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	err = p.Service.CreateCompany(context.TODO(), companyEvent.Company.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) addEmployeeToCompany(msg *ckafka.Message) error {
	companyEmployeeEvent := schema.NewCompanyEmployeeEvent()
	err := companyEmployeeEvent.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	err = p.Service.AddEmployeeToCompany(context.TODO(), companyEmployeeEvent.CompanyID, companyEmployeeEvent.EmployeeID)
	if err != nil {
		return err
	}

	return nil
}
