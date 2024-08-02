package database

import (
	"auth-service/config"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPostgresDatabase(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cfg := config.GetConfig()

		if cfg == nil {
			fmt.Println("config is nil")
		}

		database := NewPostgresDatabase(config.GetConfig())

		db := Database.GetDb(database)

		if db == nil {
			t.Fatal("gorm db is nil")
		}

		assert.NotNil(t, db)
	})
}
