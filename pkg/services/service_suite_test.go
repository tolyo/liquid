package services

import (
	"liquid/pkg/conf"
	"liquid/pkg/db"
	"liquid/pkg/utils"
	"liquid/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	// appEntity1 models.AppEntityId
	// paymentAccount1 models.paymentAccountId // seller
	// appEntity2 models.AppEntityId
	// paymentAccount2 models.paymentAccountId // buyer
}

func TestServiceTestSuite(t *testing.T) {
	conf.LoadTestConfig()
	suite.Run(t, &ServiceTestSuite{})
}

func (assert *ServiceTestSuite) SetupSuite() {
	db.SetupInstance()
	sql.MigrateUp()
}

func (assert *ServiceTestSuite) SetupTest() {

}

func (assert *ServiceTestSuite) TearDownTest() {
	utils.Each([]string{
		"app_entity WHERE external_id != 'MASTER'",
	}, db.DeleteAll)
}

func (assert *ServiceTestSuite) TearDownAllSuite() {
	sql.MigrateDown()
	db.Instance().Close()
}
