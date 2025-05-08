package kafka

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/timecard-service/application/kafka/schema"
	"github.com/patricksferraz/timecard-service/domain/entity"
	"github.com/patricksferraz/timecard-service/domain/service"
	"github.com/patricksferraz/timecard-service/infrastructure/external"
	"github.com/patricksferraz/timecard-service/infrastructure/external/topic"
)

type KafkaProcessor struct {
	Service *service.Service
	Kc      *external.KafkaConsumer
}

func NewKafkaProcessor(service *service.Service, kafkaConsumer *external.KafkaConsumer) *KafkaProcessor {
	return &KafkaProcessor{
		Service: service,
		Kc:      kafkaConsumer,
	}
}

func (p *KafkaProcessor) Consume() {
	p.Kc.Consumer.SubscribeTopics(p.Kc.ConsumerTopics, nil)
	for {
		msg, err := p.Kc.Consumer.ReadMessage(-1)
		if err == nil {
			// fmt.Println(string(msg.Value))
			err := p.processMessage(msg)
			if err != nil {
				fmt.Println(err)
			}
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

func (p *KafkaProcessor) processMessage(msg *ckafka.Message) error {

	event, err := p.processEvent(msg)
	if err != nil {
		return fmt.Errorf("event processing error %s", err)
	}

	switch _topic := *msg.TopicPartition.Topic; _topic {
	case topic.NEW_EMPLOYEE:
		// TODO: add fault tolerance
		err := p.createEmployee(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("creation employee error %s", err)
		}
	case topic.NEW_COMPANY:
		err := p.createCompany(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("creation company error %s", err)
		}
	case topic.ADD_EMPLOYEE_TO_COMPANY:
		err := p.addEmployeeToCompany(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("addition employee to company error %s", err)
		}
	case topic.NEW_WORK_SCALE:
		err := p.createWorkScale(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("creation work scale error %s", err)
		}
	case topic.ADD_WORK_SCALE_TO_EMPLOYEE:
		err := p.addWorkScaleToEmployee(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("addition work scale to employee error %s", err)
		}
	case topic.ADD_CLOCK_TO_WORK_SCALE:
		err := p.addClockToWorkScale(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("addition clock to work scale error %s", err)
		}
	case topic.UPDATE_CLOCK:
		err := p.updateClock(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("update clock error %s", err)
		}
	case topic.DELETE_CLOCK:
		err := p.deleteClock(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("delete clock error %s", err)
		}
	default:
		return fmt.Errorf("not a valid topic %s", string(msg.Value))
	}

	err = p.completeEvent(event)
	if err != nil {
		return fmt.Errorf("event completion error %s", err)
	}

	return nil
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

func (p *KafkaProcessor) createWorkScale(msg *ckafka.Message) error {
	workScaleEvent := schema.NewWorkScaleEvent()
	err := workScaleEvent.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	err = p.Service.CreateWorkScale(context.TODO(), workScaleEvent.WorkScale.ID, workScaleEvent.WorkScale.CompanyID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) addWorkScaleToEmployee(msg *ckafka.Message) error {
	workScaleEmployeeEvent := schema.NewWorkScaleEmployeeEvent()
	err := workScaleEmployeeEvent.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	err = p.Service.AddWorkScaleToEmployee(context.TODO(), workScaleEmployeeEvent.CompanyID, workScaleEmployeeEvent.EmployeeID, workScaleEmployeeEvent.WorkScaleID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) addClockToWorkScale(msg *ckafka.Message) error {
	clockEvent := schema.NewClockEvent()
	err := clockEvent.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	err = p.Service.AddClockToWorkScale(context.TODO(), clockEvent.Clock.ID, clockEvent.Clock.Type, clockEvent.Clock.Clock, clockEvent.Clock.Timezone, clockEvent.Clock.WorkScaleID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) updateClock(msg *ckafka.Message) error {
	clockEvent := schema.NewClockEvent()
	err := clockEvent.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	err = p.Service.UpdateClock(context.TODO(), clockEvent.Clock.Type, clockEvent.Clock.Clock, clockEvent.Clock.Timezone, clockEvent.Clock.WorkScaleID, clockEvent.Clock.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) deleteClock(msg *ckafka.Message) error {
	deleteClockEvent := schema.NewDeleteClockEvent()
	err := deleteClockEvent.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	err = p.Service.DeleteClock(context.TODO(), deleteClockEvent.WorkScaleID, deleteClockEvent.ClockID)
	if err != nil {
		return err
	}

	return nil
}
