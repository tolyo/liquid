package services

import (
	"liquid/pkg/db"
	"liquid/pkg/models"

	"github.com/google/uuid"
)

func (assert *ServiceTestSuite) TestCreateAppEntity() {
	// when
	count := db.GetCount("app_entity")
	res := models.CreateClientEntity("test")

	// then
	assert.NotNil(res)
	_, err := uuid.Parse(string(res))
	assert.Nil(err)
	assert.Equal(count+1, db.GetCount("app_entity"))

	// when
	res2 := models.CreateClientEntity("test1")

	// then
	assert.NotEqual(res, res2)
	assert.Equal(count+2, db.GetCount("app_entity"))

	// when same id
	res3 := models.CreateClientEntity("test1")

	// then should not create anything
	assert.Equal(res3, res2)
	assert.Equal(count+2, db.GetCount("app_entity"))
}
