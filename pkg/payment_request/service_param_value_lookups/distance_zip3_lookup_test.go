package serviceparamvaluelookups

import (
	"strconv"
	"testing"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
	"github.com/transcom/mymove/pkg/unit"
)

func (suite *ServiceParamValueLookupsSuite) TestDistanceZip3Lookup() {
	key := models.ServiceItemParamNameDistanceZip3.String()

	suite.T().Run("Calculate zip3 distance", func(t *testing.T) {
		mtoServiceItem := testdatagen.MakeDefaultMTOServiceItem(suite.DB())

		paymentRequest := testdatagen.MakePaymentRequest(suite.DB(),
			testdatagen.Assertions{
				PaymentRequest: models.PaymentRequest{
					MoveTaskOrderID: mtoServiceItem.MoveTaskOrderID,
				},
			})

		paramCache := ServiceParamsCache{}
		paramCache.Initialize(suite.DB())

		paramLookup, err := ServiceParamLookupInitialize(suite.DB(), suite.planner, mtoServiceItem.ID, paymentRequest.ID, paymentRequest.MoveTaskOrderID, &paramCache)
		suite.FatalNoError(err)

		distanceStr, err := paramLookup.ServiceParamValue(key)
		suite.FatalNoError(err)
		expected := strconv.Itoa(defaultZip3Distance)
		suite.Equal(expected, distanceStr)

		var mtoShipment models.MTOShipment
		suite.DB().Find(&mtoShipment, mtoServiceItem.MTOShipmentID)

		suite.Equal(unit.Miles(defaultZip3Distance), *mtoShipment.Distance)

		// Verify value from paramCache
		paramCacheValue := paramCache.ParamValue(*mtoServiceItem.MTOShipmentID, key)
		suite.Equal(expected, *paramCacheValue)
	})
}
