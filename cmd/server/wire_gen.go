// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/openmeterio/openmeter/config"
	"github.com/openmeterio/openmeter/openmeter/app"
	"github.com/openmeterio/openmeter/openmeter/ent/db"
	"github.com/openmeterio/openmeter/openmeter/ingest"
	"github.com/openmeterio/openmeter/openmeter/meter"
	"github.com/openmeterio/openmeter/openmeter/namespace"
	"github.com/openmeterio/openmeter/openmeter/streaming"
	"github.com/openmeterio/openmeter/openmeter/watermill/eventbus"
	"github.com/openmeterio/openmeter/pkg/kafka/metrics"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/semconv/v1.20.0"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
)

// Injectors from wire.go:

func initializeApplication(ctx context.Context, conf config.Configuration, logger *slog.Logger) (Application, func(), error) {
	telemetryConfig := conf.Telemetry
	metricsTelemetryConfig := telemetryConfig.Metrics
	resource := NewOtelResource(conf)
	meterProvider, cleanup, err := app.NewMeterProvider(ctx, metricsTelemetryConfig, resource, logger)
	if err != nil {
		return Application{}, nil, err
	}
	traceTelemetryConfig := telemetryConfig.Trace
	tracerProvider, cleanup2, err := app.NewTracerProvider(ctx, traceTelemetryConfig, resource, logger)
	if err != nil {
		cleanup()
		return Application{}, nil, err
	}
	textMapPropagator := NewTextMapPropagator()
	globalInitializer := app.GlobalInitializer{
		Logger:            logger,
		MeterProvider:     meterProvider,
		TracerProvider:    tracerProvider,
		TextMapPropagator: textMapPropagator,
	}
	aggregationConfiguration := conf.Aggregation
	clickHouseAggregationConfiguration := aggregationConfiguration.ClickHouse
	v, err := app.NewClickHouse(clickHouseAggregationConfiguration)
	if err != nil {
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	v2 := conf.Meters
	inMemoryRepository := app.NewMeterRepository(v2)
	clickhouseConnector, err := app.NewClickHouseStreamingConnector(aggregationConfiguration, v, inMemoryRepository, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	postgresConfig := conf.Postgres
	meter := NewMeter(meterProvider)
	driver, cleanup3, err := app.NewPostgresDriver(ctx, postgresConfig, meterProvider, meter, tracerProvider, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	db := app.NewDB(driver)
	entPostgresDriver, cleanup4 := app.NewEntPostgresDriver(db, logger)
	client := app.NewEntClient(entPostgresDriver)
	health := app.NewHealthChecker(logger)
	telemetryHandler := app.NewTelemetryHandler(metricsTelemetryConfig, health)
	v3, cleanup5 := app.NewTelemetryServer(telemetryConfig, telemetryHandler)
	producer, err := app.NewKafkaProducer(conf, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	metrics, err := app.NewKafkaMetrics(meter)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	watermillClientID := _wireWatermillClientIDValue
	ingestConfiguration := conf.Ingest
	topicProvisioner, err := app.NewKafkaTopicProvisioner(ingestConfiguration, logger, meter)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	publisher, cleanup6, err := app.NewPublisher(ctx, conf, watermillClientID, logger, meter, topicProvisioner)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	eventbusPublisher, err := app.NewEventBusPublisher(publisher, conf, logger)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	namespacedTopicResolver, err := app.NewNamespacedTopicResolver(conf)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	collector, err := app.NewKafkaIngestCollector(producer, namespacedTopicResolver)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	ingestCollector, cleanup7, err := app.NewIngestCollector(conf, collector, logger, meter)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	namespaceHandler, err := app.NewKafkaNamespaceHandler(namespacedTopicResolver, topicProvisioner, conf)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	v4 := app.NewNamespaceHandlers(namespaceHandler, clickhouseConnector)
	namespaceConfiguration := conf.Namespace
	manager, err := app.NewNamespaceManager(v4, namespaceConfiguration)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	application := Application{
		GlobalInitializer:  globalInitializer,
		StreamingConnector: clickhouseConnector,
		MeterRepository:    inMemoryRepository,
		EntClient:          client,
		TelemetryServer:    v3,
		KafkaProducer:      producer,
		KafkaMetrics:       metrics,
		EventPublisher:     eventbusPublisher,
		IngestCollector:    ingestCollector,
		NamespaceHandlers:  v4,
		NamespaceManager:   manager,
		Meter:              meter,
		TracerProvider:     tracerProvider,
		MeterProvider:      meterProvider,
	}
	return application, func() {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireWatermillClientIDValue = app.WatermillClientID(otelName)
)

// TODO: is this necessary? Do we need a logger first?
func initializeLogger(conf config.Configuration) *slog.Logger {
	telemetryConfig := conf.Telemetry
	logTelemetryConfiguration := telemetryConfig.Log
	resource := NewOtelResource(conf)
	logger := app.NewLogger(logTelemetryConfiguration, resource)
	return logger
}

// wire.go:

type Application struct {
	app.GlobalInitializer

	StreamingConnector streaming.Connector
	MeterRepository    meter.Repository
	EntClient          *db.Client
	TelemetryServer    app.TelemetryServer
	KafkaProducer      *kafka.Producer
	KafkaMetrics       *metrics.Metrics
	EventPublisher     eventbus.Publisher

	IngestCollector ingest.Collector

	NamespaceHandlers []namespace.Handler
	NamespaceManager  *namespace.Manager

	Meter metric.Meter

	// TODO: move to global setter
	TracerProvider trace.TracerProvider
	MeterProvider  metric.MeterProvider
}

// TODO: consider moving this to a separate package
// TODO: make sure this doesn't generate any random IDs, because it is called twice
func NewOtelResource(conf config.Configuration) *resource.Resource {
	extraResources, _ := resource.New(context.Background(), resource.WithContainer(), resource.WithAttributes(semconv.ServiceName("openmeter"), semconv.ServiceVersion(version), semconv.DeploymentEnvironment(conf.Environment)),
	)

	res, _ := resource.Merge(resource.Default(), extraResources)

	return res
}

// TODO: consider moving this to a separate package
func NewMeter(meterProvider metric.MeterProvider) metric.Meter {
	return meterProvider.Meter(otelName)
}

// TODO: consider moving this to a separate package
func NewTextMapPropagator() propagation.TextMapPropagator {
	return propagation.TraceContext{}
}
