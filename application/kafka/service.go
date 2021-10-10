package kafka

import (
	"context"
	"fmt"

	"github.com/c-4u/timecard-service/application/kafka/schema"
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

func (p *KafkaProcessor) processMessage(msg *ckafka.Message) {
	switch _topic := *msg.TopicPartition.Topic; _topic {
	case topic.NEW_EMPLOYEE:
		// TODO: add fault tolerance
		err := p.createEmployee(msg)
		if err != nil {
			fmt.Println("creation error ", err)
		}
	case topic.NEW_COMPANY:
		err := p.createCompany(msg)
		if err != nil {
			fmt.Println("creation error ", err)
		}
	case topic.ADD_EMPLOYEE_TO_COMPANY:
		err := p.addEmployeeToCompany(msg)
		if err != nil {
			fmt.Println("addition error ", err)
		}
	default:
		fmt.Println("not a valid topic", string(msg.Value))
	}
}

func (p *KafkaProcessor) createEmployee(msg *ckafka.Message) error {
	employeeEvent := &schema.EmployeeEvent{}
	err := employeeEvent.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	event, err := p.Service.ProcessEvent(context.TODO(), employeeEvent.ID, *msg.TopicPartition.Topic)
	if err != nil {
		return err
	}

	err = p.Service.CreateEmployee(context.TODO(), employeeEvent.Employee.ID)
	if err != nil {
		p.Service.Repository.PublishEvent(context.TODO(), string(msg.Value), *msg.TopicPartition.Topic, employeeEvent.Employee.ID)
		return err
	}

	err = p.Service.CompleteEvent(context.TODO(), event.ID)
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

	event, err := p.Service.ProcessEvent(context.TODO(), companyEvent.ID, *msg.TopicPartition.Topic)
	if err != nil {
		return err
	}

	err = p.Service.CreateCompany(context.TODO(), companyEvent.Company.ID)
	if err != nil {
		p.Service.Repository.PublishEvent(context.TODO(), string(msg.Value), *msg.TopicPartition.Topic, companyEvent.Company.ID)
		return err
	}

	err = p.Service.CompleteEvent(context.TODO(), event.ID)
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

	event, err := p.Service.ProcessEvent(context.TODO(), companyEmployeeEvent.ID, *msg.TopicPartition.Topic)
	if err != nil {
		return err
	}

	err = p.Service.AddEmployeeToCompany(context.TODO(), companyEmployeeEvent.CompanyID, companyEmployeeEvent.EmployeeID)
	if err != nil {
		p.Service.Repository.PublishEvent(context.TODO(), string(msg.Value), *msg.TopicPartition.Topic, companyEmployeeEvent.CompanyID)
		return err
	}

	err = p.Service.CompleteEvent(context.TODO(), event.ID)
	if err != nil {
		return err
	}

	return nil
}
