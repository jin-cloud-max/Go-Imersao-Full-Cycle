package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jin-cloud-max/gateway/adapter/presenter/transaction"
	"github.com/jin-cloud-max/gateway/domain/entity"
	"github.com/jin-cloud-max/gateway/useCase/process_transaction"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	expectedOutput := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "You dont have the limit for this transaction",
	}

	// outputJson, _ := json.Marshal(expectedOutput)

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}

	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())

	err := producer.Publish(expectedOutput, []byte("1"), "test")

	assert.Nil(t, err)
}
