package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jin-cloud-max/gateway/adapter/broker/kafka"
	"github.com/jin-cloud-max/gateway/adapter/factory"
	"github.com/jin-cloud-max/gateway/adapter/presenter/transaction"
	"github.com/jin-cloud-max/gateway/useCase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)

	repository := repositoryFactory.CreateTransactionRepository()

	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.server": "kafka:9092",
	}

	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	var msgChan = make(chan *ckafka.Message)

	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.server": "kafka:9092",
		"client_id":        "goapp",
		"group_id":         "goapp",
	}

	topics := []string{"transactions"}

	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go consumer.Consume(msgChan)

	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)

		usecase.Execute(input)
	}

}
