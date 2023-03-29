package sender

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"route256/loms/internal/models"
	"time"

	"github.com/Shopify/sarama"
)

type Repository interface {
	GetOutboxOrders(ctx context.Context) ([]models.OutboxOrder, error)
	UpdateOutboxOrder(ctx context.Context, outboxID int64) error
}

type OrderSender struct {
	ctx        context.Context
	producer   sarama.SyncProducer
	topic      string
	batch      []*sarama.ProducerMessage
	batchSize  int
	repository Repository
}

func NewOrderSender(ctx context.Context, producer sarama.SyncProducer, topic string, batchSize int, repository Repository) *OrderSender {
	s := &OrderSender{
		ctx:        ctx,
		producer:   producer,
		topic:      topic,
		batch:      make([]*sarama.ProducerMessage, 0, batchSize),
		batchSize:  batchSize,
		repository: repository,
	}
	return s
}

func (s *OrderSender) SendOrderID(outboxOrder models.OutboxOrder) error {

	jb, err := json.Marshal(outboxOrder)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     s.topic,
		Partition: -1,
		Value:     sarama.StringEncoder(jb),
		Key:       sarama.StringEncoder(fmt.Sprint(outboxOrder.OrderID)),
	}

	s.batch = append(s.batch, msg)

	if len(s.batch) >= s.batchSize {
		err = s.producer.SendMessages(s.batch)
		if err != nil {
			log.Println("Kafka SendMessages err:", err.Error())
			return err
		}
		s.batch = s.batch[0:0]
	}

	return nil
}

func (s *OrderSender) Run() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			outboxOrders, err := s.repository.GetOutboxOrders(s.ctx)
			if err != nil {
				log.Println("Err GetOutboxOrders:", err.Error())
				continue
			}

			for _, order := range outboxOrders {
				err = s.SendOrderID(order)
				if err != nil {
					log.Println("Err SendOrderID:", err.Error())
					continue
				}
				err = s.repository.UpdateOutboxOrder(s.ctx, order.ID)
				for err != nil {
					err = s.repository.UpdateOutboxOrder(s.ctx, order.ID)
					if err != nil {
						log.Println("Err UpdateOutboxOrder:", err.Error())
					}
				}
			}
		}
	}()
}
