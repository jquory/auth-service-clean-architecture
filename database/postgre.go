package database

import (
	"e-rt/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

type postgresDatabase struct {
	DB *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgresDatabase(config *config.Config) Database {
	once.Do(func() {

		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.DB.Host,
			config.DB.Port,
			config.DB.User,
			config.DB.Password,
			config.DB.DBName,
			config.DB.SslMode,
		)

		loggerDb := logger.New(
			log.New(os.Stdout, "\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			})

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: loggerDb,
		})
		if err != nil {

			panic(err)
		}

		sqlDb, err := db.DB()
		if err != nil {
			panic(err)
		}

		sqlDb.SetMaxOpenConns(config.DB.MaxConnectionPool)
		sqlDb.SetMaxIdleConns(config.DB.MaxIdleTime)
		sqlDb.SetConnMaxLifetime(time.Duration(config.DB.MaxLifetimePool) * time.Second)
		dbInstance = &postgresDatabase{DB: db}
	})

	return dbInstance
}

func (postgres *postgresDatabase) GetDb() *gorm.DB {
	return dbInstance.DB
}
