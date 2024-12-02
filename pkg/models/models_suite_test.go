package models

import (
	"liquid/pkg/conf"
	"liquid/pkg/db"
	"liquid/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ModelsTestSuite struct {
	suite.Suite
}

func TestModelSuite(t *testing.T) {
	conf.LoadTestConfig()
	suite.Run(t, &ModelsTestSuite{})
}

func (assert *ModelsTestSuite) SetupTest() {
	db.SetupInstance()
	sql.MigrateUp()
}

func (assert *ModelsTestSuite) TearDownTest() {
	sql.MigrateDown()
	db.Instance().Close()
}
