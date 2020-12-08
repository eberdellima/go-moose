package database

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
)

func setEnvVariablesForTesting(envVariables *DBConfig) {

	_ = os.Setenv("DB_ALIAS", envVariables.Alias)
	_ = os.Setenv("DB_PORT", envVariables.port)
	_ = os.Setenv("DB_HOST", envVariables.host)
	_ = os.Setenv("DB_NAME", envVariables.name)
	_ = os.Setenv("DB_USER", envVariables.user)
	_ = os.Setenv("DB_PASS", envVariables.password)
}

func TestInitializeDBConnection(t *testing.T) {

	assert.Panics(t, func() { InitializeDBConnection() })

	// Presume that we have already created the test_db
	envVariables := &DBConfig{
		Alias:    "test_db",
		Driver:   "mysql",
		host:     "localhost",
		name:     "test_db",
		user:     "root",
		password: "test_pass",
		port:     "3306",
	}

	setEnvVariablesForTesting(envVariables)

	db := InitializeDBConnection()

	assert.IsType(t, &gorm.DB{
		RWMutex:      sync.RWMutex{},
		Value:        nil,
		Error:        nil,
		RowsAffected: 0,
	}, db)
}