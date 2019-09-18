package edisegment

import (
	"testing"
)

func (suite *SegmentSuite) TestValidateL7() {
	validL7Default := L7{}
	validL7WithValues := L7{
		LadingLineItemNumber: 1,
		TariffNumber:         "X",
		TariffItemNumber:     "X",
		TariffDistance:       1,
	}

	suite.T().Run("validate success default", func(t *testing.T) {
		err := suite.validator.Struct(validL7Default)
		suite.NoError(err)
	})

	suite.T().Run("validate success with values", func(t *testing.T) {
		err := suite.validator.Struct(validL7WithValues)
		suite.NoError(err)
	})

	suite.T().Run("validate failure 1", func(t *testing.T) {
		l7 := L7{
			LadingLineItemNumber: -3,                  // min
			TariffNumber:         "XXXXXXXX",          // max
			TariffItemNumber:     "XXXXXXXXXXXXXXXXX", // max
			TariffDistance:       -2,                  // min
		}

		err := suite.validator.Struct(l7)
		suite.ValidateError(err, "LadingLineItemNumber", "min")
		suite.ValidateError(err, "TariffNumber", "max")
		suite.ValidateError(err, "TariffItemNumber", "max")
		suite.ValidateError(err, "TariffDistance", "min")
		suite.ValidateErrorLen(err, 4)
	})

	suite.T().Run("validate failure 2", func(t *testing.T) {
		l7 := L7{
			LadingLineItemNumber: 1000,   // max
			TariffDistance:       100000, // max
		}

		err := suite.validator.Struct(l7)
		suite.ValidateError(err, "LadingLineItemNumber", "max")
		suite.ValidateError(err, "TariffDistance", "max")
		suite.ValidateErrorLen(err, 2)
	})
}
