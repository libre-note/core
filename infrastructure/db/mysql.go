package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"librenote/infrastructure/config"
)

// must call once before server boot to Get() the db instance
func connectMysql() (err error) {
	if dbc.DB != nil {
		logrus.Info("db already initialized")
		return nil
	}

	cfg := config.Get().Database
	dbURL := fmt.Sprintf("%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v\n", err)
	}

	if cfg.MaxOpenConn > 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConn)
	}

	if cfg.MaxIdleConn > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConn)
	}

	if cfg.MaxLifeTime > 0 {
		db.SetConnMaxLifetime(cfg.MaxLifeTime)
	}

	dbc.DB = db
	return nil
}
