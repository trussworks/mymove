package ghcrateengine

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/unit"
)

type DomesticServicePricingData struct {
	MoveDate      time.Time
	ServiceAreaID uuid.UUID
	Distance      unit.Miles
	Weight        unit.CWTFloat // record this here as 5.00 if actualWt less than minimum of 5.00 cwt (500lb)
	IsPeakPeriod  bool
	ContractCode  string
}

func lookupDomesticLinehaulRate(db *pop.Connection, d DomesticServicePricingData) unit.Millicents {
	// TODO: check/correct syntax && implement when models are created
	// query := db.Where(
	// 	"is_peak_period = d.IsPeakPeriod").Join(
	// 	"serviceAreas sa", "sa.id=re_domestic_linehaul_prices.service_area_id").Join(
	// 	"re_contracts c", "re_contract.id=re_domestic_linehaul_prices.contract_id").Join(
	// 	"re_contract_years cy", "re_contract.id=cy.contract_id").Where(
	// 		"sa.id=d.ServiceAreaID").Where(
	// 		"weight gte weight_lower").Where(
	// 			"weight lte weight_upper").Where(
	// 				"distance gte miles_lower").Where(
	// 					"distance lte miles_upper").Where(
	// 										"re_contracts.code=?", d.ContractCode).Where(
	// 											"d.MoveDate gte cy.start_date").Where(
	// 												"d.MoveDate lte cy.end_date")

	// domesticLinehaulPricing := models.DomesticLinehaulPricing{}
	// err := db.query.All(&domesticLinehaulPricing)
	// if err != nil {
	// 	return err
	// }

	var stubPrice unit.Millicents = 272700
	return stubPrice
}

// Calculation Functions
// CalculateBaseDomesticLinehaul calculates the cost domestic linehaul and returns the cost in millicents
func (gre *GHCRateEngine) CalculateBaseDomesticLinehaul(d DomesticServicePricingData) unit.Millicents {
	rate := lookupDomesticLinehaulRate(gre.db, d)

	return rate.MultiplyFloat64(float64(d.Weight))
}
