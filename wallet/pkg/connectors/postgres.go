package connectors

import (
	"context"
	"database/sql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log/slog"
	"sync"
)

type Postgres struct {
	db   *sql.DB
	DSN  string
	init sync.Once
}

func (p *Postgres) Client(ctx context.Context) *sql.DB {
	var e error
	p.init.Do(func() {
		p.db, e = sql.Open("postgres", p.DSN)
		if e != nil {
			slog.Error("Error connecting to Postgres", slog.String("error", e.Error()))
		}
		e = p.db.PingContext(ctx)
		if e != nil {
			slog.Error("Error connecting to Postgres", slog.String("error", e.Error()))
		}
		slog.Info("Successfully connected to Postgres database")
	})
	return p.db
}

func (p *Postgres) Close(ctx context.Context) {
	if err := p.db.Close(); err != nil {
		slog.Error("postgresClient.Close", slog.String("error", err.Error()))
	}

	slog.Info(
		"postgres disconnected",
		slog.String("database", p.DSN),
	)
}
