package models

func (assert *ModelsTestSuite) TestGenerateId() {
	res := GenerateTigerId()
	assert.NotNil(res)
	assert.Positive(res)
}
