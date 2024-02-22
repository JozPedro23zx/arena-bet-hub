package kafka

import (
	"testing"

	createtournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-usecase/tournament/create-tournament"
	tournamentkafka "github.com/JozPedro23zx/arena-bet-hub/infrastructure/presenter/tournament-kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	expectOutput := createtournament.TournamentOutputDto{
		ID:       "tournament123",
		Finished: false,
	}

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}

	producer := NewKafkaProducer(&configMap, tournamentkafka.NewCreateTournamentPresenter())
	err := producer.Publish(expectOutput, []byte("1"), "test")
	assert.Nil(t, err)
}
