package common

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ThreeDotsLabs/watermill/message"
	"go.opentelemetry.io/otel/metric"

	"github.com/openmeterio/openmeter/app/config"
	watermillkafka "github.com/openmeterio/openmeter/openmeter/watermill/driver/kafka"
	"github.com/openmeterio/openmeter/openmeter/watermill/eventbus"
)

func NewBrokerConfiguration(
	kafkaConfig config.KafkaConfiguration,
	logConfig config.LogTelemetryConfig,
	appMetadata Metadata,
	logger *slog.Logger,
	meter metric.Meter,
) watermillkafka.BrokerOptions {
	return watermillkafka.BrokerOptions{
		KafkaConfig:  kafkaConfig,
		ClientID:     appMetadata.OpenTelemetryName, // TODO: use a better name or rename otel name
		Logger:       logger,
		MetricMeter:  meter,
		DebugLogging: logConfig.Level == slog.LevelDebug,
	}
}

// NOTE: this is also used by the sink-worker that requires control over how the publisher is closed
func NewPublisher(
	ctx context.Context,
	options watermillkafka.PublisherOptions,
	logger *slog.Logger,
) (message.Publisher, func(), error) {
	publisher, err := watermillkafka.NewPublisher(ctx, options)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize event publisher: %w", err)
	}

	return publisher, func() {
		// TODO: isn't this logged by the publisher itself?
		logger.Info("closing event publisher")

		if err = publisher.Close(); err != nil {
			logger.Error("failed to close event publisher", "error", err)
		}
	}, nil
}

func NewEventBusPublisher(
	publisher message.Publisher,
	conf config.EventsConfiguration,
	logger *slog.Logger,
) (eventbus.Publisher, error) {
	eventBusPublisher, err := eventbus.New(eventbus.Options{
		Publisher:              publisher,
		Config:                 conf,
		Logger:                 logger,
		MarshalerTransformFunc: watermillkafka.AddPartitionKeyFromSubject,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize event bus publisher: %w", err)
	}

	return eventBusPublisher, nil
}
