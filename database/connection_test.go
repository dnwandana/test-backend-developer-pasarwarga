package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConnection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	assert.NotNil(t, db)

}
