package pubsub

import (
	"context"
	"fmt"
	"go-template-wire/configs"
	"go-template-wire/pkg/failure"
	"go-template-wire/pkg/logger"

	"cloud.google.com/go/pubsub"
)

type PubSub struct {
	log    *logger.Logger
	client *pubsub.Client
}

func New(cfg *configs.Config, log *logger.Logger) (*PubSub, error) {
	client, err := pubsub.NewClient(context.Background(), cfg.GCP.ProjectID)
	if err != nil {
		return nil, err
	}

	return &PubSub{
		log:    log,
		client: client,
	}, nil
}

func (p *PubSub) Publish(dto *PubSubPublishDTO) error {
	topic := p.getPubSubTopic(dto)
	message := getPubSubMessage(dto)

	ctx := context.Background()
	result := topic.Publish(ctx, message)
	id, err := result.Get(ctx)
	if err != nil {
		return failure.ErrWithTrace(fmt.Errorf(
			"Failed to publish message=%s topic=%s: %w", dto.TopicID, string(dto.Data), err,
		))
	}

	p.log.Infof("Published message=%s to topic=%s", id, dto.TopicID)
	return nil
}

func (p *PubSub) getPubSubTopic(dto *PubSubPublishDTO) *pubsub.Topic {
	topic := p.client.Topic(dto.TopicID)
	if dto.OrderingKey != nil {
		topic.EnableMessageOrdering = true
	}
	return topic
}

func getPubSubMessage(dto *PubSubPublishDTO) *pubsub.Message {
	message := &pubsub.Message{
		Data: dto.Data,
		Attributes: map[string]string{
			"schemaencoding": "JSON",
		},
	}
	if dto.OrderingKey != nil {
		message.OrderingKey = *dto.OrderingKey
	}
	return message
}
