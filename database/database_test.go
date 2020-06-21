package database

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	assert := assert.New(t)
	db, err := GetInstance()
	assert.Nil(err)
	assert.NotNil(db)

	// Make sure the same database instance is returned
	db2, err := GetInstance()
	assert.Nil(err)
	assert.Equal(&db, &db2)
}

func TestInitializeDatabase(t *testing.T) {
	assert := assert.New(t)

	// Make sure the database instance singleton is nil before initializing the database
	DestroyInstance()
	assert.Nil(db)
	InitializeDatabase()
	assert.NotNil(db)
}
