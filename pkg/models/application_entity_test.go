package models

import "liquid/pkg/db"

func (assert *ModelsTestSuite) TestMasterEntity() {
	// expect master entity to exist
	assert.LessOrEqual(1, db.GetCount("app_entity"))
	res, err := FindAppEntityExternalId("MASTER")
	assert.NotNil(res)
	assert.Nil(err)
}

func (assert *ModelsTestSuite) TestEmptyEntity() {
	// expect master entity to exist
	res, err := FindAppEntityExternalId("FAIL")
	assert.Equal(AppEntityId(""), res)
	assert.NotNil(err)
}
