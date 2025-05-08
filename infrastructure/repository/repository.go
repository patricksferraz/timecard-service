package repository

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/timecard-service/domain/entity"
	"github.com/patricksferraz/timecard-service/infrastructure/db"
	"github.com/patricksferraz/timecard-service/infrastructure/external"
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

func (r *Repository) RegisterTimeRecord(ctx context.Context, timeRecord *entity.TimeRecord) error {
	err := r.P.Db.Create(timeRecord).Error
	return err
}

func (r *Repository) SaveTimeRecord(ctx context.Context, timeRecord *entity.TimeRecord) error {
	err := r.P.Db.Save(timeRecord).Error
	return err
}

func (r *Repository) FindTimeRecord(ctx context.Context, id string) (*entity.TimeRecord, error) {
	var timeRecord entity.TimeRecord
	r.P.Db.First(&timeRecord, "id = ?", id)

	if timeRecord.ID == "" {
		return nil, fmt.Errorf("no time record found")
	}

	return &timeRecord, nil
}

func (r *Repository) CreateEpoch(ctx context.Context, epoch *entity.Epoch) error {
	err := r.P.Db.Create(epoch).Error
	return err
}

func (r *Repository) FindEpoch(ctx context.Context, id string) (*entity.Epoch, error) {
	var epoch entity.Epoch
	r.P.Db.First(&epoch, "id = ?", id)

	if epoch.ID == "" {
		return nil, fmt.Errorf("no epoch found")
	}

	return &epoch, nil
}

func (r *Repository) SaveEpoch(ctx context.Context, epoch *entity.Epoch) error {
	err := r.P.Db.Save(epoch).Error
	return err
}

// TODO: add search epochs

func (r *Repository) AddEmployeeToCompany(ctx context.Context, companyEmployee *entity.CompaniesEmployee) error {
	err := r.P.Db.Create(companyEmployee).Error
	return err
}

func (r *Repository) FindCompanyEmployee(ctx context.Context, companyID, employeeID string) (*entity.CompaniesEmployee, error) {
	var companyEmployee entity.CompaniesEmployee
	r.P.Db.First(&companyEmployee, "company_id = ? AND employee_id = ?", companyID, employeeID)
	if companyEmployee.CompanyID == "" {
		return nil, fmt.Errorf("company-employee relationship not found")
	}

	return &companyEmployee, nil
}

func (r *Repository) SaveCompanyEmployee(ctx context.Context, companyEmployee *entity.CompaniesEmployee) error {
	err := r.P.Db.Model(entity.CompaniesEmployee{}).Where("company_id = ? AND employee_id = ?", companyEmployee.CompanyID, companyEmployee.EmployeeID).Update("work_scale_id", companyEmployee.WorkScaleID).Error
	return err
}

func (r *Repository) CreateWorkScale(ctx context.Context, workScale *entity.WorkScale) error {
	err := r.P.Db.Create(workScale).Error
	return err
}

func (r *Repository) FindWorkScale(ctx context.Context, workScaleID string) (*entity.WorkScale, error) {
	var workScale entity.WorkScale
	r.P.Db.Preload("Clocks").First(&workScale, "id = ?", workScaleID)

	if workScale.ID == "" {
		return nil, fmt.Errorf("no work scale found")
	}

	return &workScale, nil
}

func (r *Repository) SaveWorkScale(ctx context.Context, workScale *entity.WorkScale) error {
	err := r.P.Db.Save(workScale).Error
	return err
}

func (r *Repository) CreateClock(ctx context.Context, clock *entity.Clock) error {
	err := r.P.Db.Create(clock).Error
	return err
}

func (r *Repository) FindClock(ctx context.Context, workScaleID, clockID string) (*entity.Clock, error) {
	var clock entity.Clock
	r.P.Db.First(&clock, "id = ? AND work_scale_id = ?", clockID, workScaleID)

	if clock.ID == "" {
		return nil, fmt.Errorf("no clock found")
	}

	return &clock, nil
}

func (r *Repository) DeleteClock(ctx context.Context, workScaleID, clockID string) error {
	err := r.P.Db.Where("id = ? AND work_scale_id = ?", clockID, workScaleID).Delete(entity.Clock{}).Error
	return err
}

func (r *Repository) SaveClock(ctx context.Context, clock *entity.Clock) error {
	err := r.P.Db.Save(clock).Error
	return err
}
