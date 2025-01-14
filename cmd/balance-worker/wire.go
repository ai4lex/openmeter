//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"log/slog"

	"github.com/google/wire"

	"github.com/openmeterio/openmeter/app/common"
	"github.com/openmeterio/openmeter/app/config"
)

type Application struct {
	common.GlobalInitializer
	common.Migrator
	common.Runner
}

func initializeApplication(ctx context.Context, conf config.Configuration, logger *slog.Logger) (Application, func(), error) {
	wire.Build(
		metadata,
		common.Config,
		common.Framework,
		common.Telemetry,
		common.NewDefaultTextMapPropagator,
		common.Database,
		common.ClickHouse,
		common.KafkaTopic,
		common.Watermill,
		common.WatermillRouter,
		common.OpenMeter,
		common.BalanceWorkerAdapter,
		common.BalanceWorker,
		wire.Struct(new(Application), "*"),
	)
	return Application{}, nil, nil
}

// TODO: is this necessary? Do we need a logger first?
func initializeLogger(conf config.Configuration) *slog.Logger {
	wire.Build(metadata, common.Config, common.Logger)

	return new(slog.Logger)
}

func metadata(conf config.Configuration) common.Metadata {
	return common.Metadata{
		ServiceName:       "openmeter",
		Version:           version,
		Environment:       conf.Environment,
		OpenTelemetryName: "openmeter.io/balance-worker",
	}
}
