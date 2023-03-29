package consumer

import (
	"encoding/json"
	"log"
	"route256/notifications/cmd/internal/models"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	ready chan bool
}

func NewConsumerGroup() Consumer {
	return Consumer{
		ready: make(chan bool),
	}
}

func (consumer *Consumer) Ready() <-chan bool {
	return consumer.ready
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			var order models.Order

			err := json.Unmarshal(message.Value, &order)
			if err != nil {
				log.Println("Unmarshall err:", err.Error())
			} else {
				log.Printf("Order: %d. New status: %s", order.OrderID, models.GetName(order.Status))
				session.MarkMessage(message, "")
			}
		case <-session.Context().Done():
			return nil
		}
	}
}
