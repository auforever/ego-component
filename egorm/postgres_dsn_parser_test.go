package egorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresDsnParser_ParseDSN(t *testing.T) {
	dsn := "user=user password=password dbname=dbname port=9920 host=localhost sslmode=disable"

	cfg, err := PostgresDsnParser(dsn).ParseDSN()
	assert.NoError(t, err)
	assert.Equal(t, "user", cfg.User)
	assert.Equal(t, "password", cfg.Password)
	assert.Equal(t, "dbname", cfg.DBName)
	assert.Equal(t, "localhost:9920", cfg.Addr)
	assert.Equal(t, "disable", cfg.Params["sslmode"])
}
