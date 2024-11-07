package db

import (
	"context"

	"github.com/ruohuaii/goddd/infrastructure/db/ent"
	"github.com/ruohuaii/goddd/infrastructure/db/ent/migrate"
)

var instance *ent.Client

type DB struct{}

func (d *DB) Init(dsn string) error {
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return err
	}
	instance = client
	return nil
}

func (d *DB) Close() error {
	if instance != nil {
		return instance.Close()
	}
	return nil
}

func Client() *ent.Client {
	checkInstance()
	return instance
}

func Migrate(ctx context.Context) error {
	checkInstance()
	err := instance.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		return err
	}
	return nil
}

func checkInstance() {
	if instance == nil {
		panic("database is not initialized")
	}
}
