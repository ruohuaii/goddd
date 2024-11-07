package di

import (
	"context"
	"github.com/ruohuaii/goddd/infrastructure/cfg"
	"github.com/ruohuaii/goddd/infrastructure/db"
	"github.com/ruohuaii/goddd/infrastructure/log"
)

type Container struct {
	cfg *cfg.Config
	log *log.Log
	db  *db.DB
}

func NewContainer() *Container {
	return &Container{
		cfg: new(cfg.Config),
		log: new(log.Log),
		db:  new(db.DB),
	}
}

func (c *Container) Init(ctx context.Context) error {
	err := c.log.Init()
	if err != nil {
		return err
	}
	err = c.cfg.Init()
	if err != nil {
		return err
	}
	dsn := cfg.GetString("DB.DSN")
	err = c.db.Init(dsn)
	if err != nil {
		return err
	}
	_ = db.Migrate(ctx)
	return nil
}

func (c *Container) Close() {
	_ = c.db.Close()
	_ = c.log.Close()
}
