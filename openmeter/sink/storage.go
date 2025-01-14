package sink

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/huandu/go-sqlbuilder"

	sinkmodels "github.com/openmeterio/openmeter/openmeter/sink/models"
	"github.com/openmeterio/openmeter/openmeter/streaming/clickhouse_connector"
)

type Storage interface {
	BatchInsert(ctx context.Context, messages []sinkmodels.SinkMessage) error
}

type ClickHouseStorageConfig struct {
	ClickHouse      clickhouse.Conn
	Database        string
	AsyncInsert     bool
	AsyncInsertWait bool
	QuerySettings   map[string]string
}

func (c ClickHouseStorageConfig) Validate() error {
	if c.ClickHouse == nil {
		return fmt.Errorf("clickhouse connection is required")
	}

	if c.Database == "" {
		return fmt.Errorf("database is required")
	}

	return nil
}

func NewClickhouseStorage(config ClickHouseStorageConfig) (*ClickHouseStorage, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &ClickHouseStorage{
		config: config,
	}, nil
}

type ClickHouseStorage struct {
	config ClickHouseStorageConfig
}

func (c *ClickHouseStorage) BatchInsert(ctx context.Context, messages []sinkmodels.SinkMessage) error {
	query := InsertEventsQuery{
		Clock:         realClock{},
		Database:      c.config.Database,
		Messages:      messages,
		QuerySettings: c.config.QuerySettings,
	}
	sql, args, err := query.ToSQL()
	if err != nil {
		return err
	}

	// By default, ClickHouse is writing data synchronously.
	// See https://clickhouse.com/docs/en/cloud/bestpractices/asynchronous-inserts
	if c.config.AsyncInsert {
		// With the `wait_for_async_insert` setting, you can configure
		// if you want an insert statement to return with an acknowledgment
		// either immediately after the data got inserted into the buffer.
		err = c.config.ClickHouse.AsyncInsert(ctx, sql, c.config.AsyncInsertWait, args...)
	} else {
		err = c.config.ClickHouse.Exec(ctx, sql, args...)
	}

	if err != nil {
		return fmt.Errorf("failed to batch insert events: %w", err)
	}

	return nil
}

type InsertEventsQuery struct {
	Clock         Clock
	Database      string
	Messages      []sinkmodels.SinkMessage
	QuerySettings map[string]string
}

func (q InsertEventsQuery) ToSQL() (string, []interface{}, error) {
	tableName := clickhouse_connector.GetEventsTableName(q.Database)

	query := sqlbuilder.ClickHouse.NewInsertBuilder()
	query.InsertInto(tableName)
	query.Cols("namespace", "validation_error", "id", "type", "source", "subject", "time", "data", "ingested_at", "stored_at")

	// Add settings
	var settings []string
	for key, value := range q.QuerySettings {
		settings = append(settings, fmt.Sprintf("%s = %s", key, value))
	}

	if len(settings) > 0 {
		query.SQL(fmt.Sprintf("SETTINGS %s", strings.Join(settings, ", ")))
	}

	for _, message := range q.Messages {
		var eventErr string
		if message.Status.Error != nil {
			eventErr = message.Status.Error.Error()
		}

		storedAt := q.Clock.Now()
		ingestedAt := storedAt

		if message.KafkaMessage != nil {
			for _, header := range message.KafkaMessage.Headers {
				// Parse ingested_at header
				if header.Key == "ingested_at" {
					var err error

					ingestedAt, err = time.Parse(time.RFC3339, string(header.Value))
					if err != nil {
						eventErr = fmt.Sprintf("failed to parse ingested_at header: %s", err)
					}
				}
			}
		}

		query.Values(
			message.Namespace,
			eventErr,
			message.Serialized.Id,
			message.Serialized.Type,
			message.Serialized.Source,
			message.Serialized.Subject,
			message.Serialized.Time,
			message.Serialized.Data,
			ingestedAt,
			storedAt,
		)
	}

	sql, args := query.Build()
	return sql, args, nil
}

// Clock is an interface for getting the current time.
// It is used to make the code testable.
type Clock interface {
	Now() time.Time
}

// realClock implements Clock using the system clock.
type realClock struct{}

func (realClock) Now() time.Time {
	return time.Now()
}
