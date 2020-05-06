package hive_go

import (
	"github.com/golang/protobuf/proto"
	"github.com/nsqio/go-nsq"
	uuid "github.com/satori/go.uuid"
)

type Subscription struct {
	repository *Repository
	consumer   *nsq.Consumer
}

func (subscription *Subscription) HandleMessage(message *nsq.Message) error {
	if len(message.Body) <= 0 {
		return nil
	}

	var secretCreated SecretCreatedV1
	err := proto.Unmarshal(message.Body, &secretCreated)
	if err != nil {
		return err
	}

	uuidID, err := uuid.FromBytes(secretCreated.Id)
	if err != nil {
		return err
	}
	uuidValue, err := uuid.FromBytes(secretCreated.Value)
	if err != nil {
		return nil
	}
	go subscription.repository.SetSecret(&Secret{
		Id:      uuidID,
		Created: secretCreated.Created,
		Value:   uuidValue,
	})

	return nil
}

func (subscription *Subscription) Stop() {
	subscription.consumer.Stop()
}

func InitHiveSubscription(serviceName string, topic string, nsqLookupAddress string, repository *Repository) *Subscription {

	if topic == "" {
		topic = "auth-secret-1"
	}

	subscription := &Subscription{repository: repository}
	consumer, err := nsq.NewConsumer(topic, serviceName, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	consumer.AddHandler(subscription)
	err = consumer.ConnectToNSQLookupd(nsqLookupAddress)
	if err != nil {
		panic(err)
	}
	subscription.consumer = consumer
	return subscription
}
