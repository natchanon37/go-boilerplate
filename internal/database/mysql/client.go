package database_mysql

import (
	"fmt"
	"go-boilerplate/configs"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type MySqlDb struct {
	db *gorm.DB
}

func (ms *MySqlDb) Connect(cfg configs.Database) error {
	cfgDB := &gorm.Config{
		Logger:  logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time { return time.Now().UTC() },
	}

	fmt.Println(ms.buildDBConnection(cfg.Write))
	// set up write and read connection
	db, err := gorm.Open(mysql.Open(ms.buildDBConnection(cfg.Write)), cfgDB)
	if err != nil {
		return err
	}

	// set up plugin for read and write
	err = db.Use(dbresolver.Register(
		dbresolver.Config{
			Replicas:          []gorm.Dialector{mysql.Open(ms.buildDBConnection(cfg.Read))},
			Sources:           []gorm.Dialector{mysql.Open(ms.buildDBConnection(cfg.Write))},
			Policy:            dbresolver.RandomPolicy{},
			TraceResolverMode: true,
		}).
		SetConnMaxLifetime(5 * time.Minute).
		SetMaxIdleConns(10).
		SetMaxOpenConns(100))
	if err != nil {
		return err
	}

	ms.db = db

	return nil
}

func (m *MySqlDb) buildDBConnection(db configs.DatabaseConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		db.User, db.Pass, db.Host, db.Port, db.Name,
	)
}

func (ms *MySqlDb) Close() error {
	if ms.db != nil {
		sqlDb, _ := ms.db.DB()
		sqlDb.Close()
	}

	return nil
}

func (ms *MySqlDb) GetDB() *gorm.DB {
	return ms.db
}
