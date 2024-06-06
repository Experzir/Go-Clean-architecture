package usecases

import (
	"userservice/modules/entities"

	"github.com/IBM/sarama"
)

type consumerHandler struct {
	eventHandler entities.EventHandler
}

func NewConsumerHandler(eventHandler entities.EventHandler) sarama.ConsumerGroupHandler {
	return &consumerHandler{eventHandler}
}

func (obj *consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (obj *consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (obj *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		obj.eventHandler.Handle(msg.Topic, msg.Value)
		session.MarkMessage(msg, "")
	}
	return nil
}
