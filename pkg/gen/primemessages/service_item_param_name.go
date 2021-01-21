// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ServiceItemParamName service item param name
//
// swagger:model ServiceItemParamName
type ServiceItemParamName string

const (

	// ServiceItemParamNameActualPickupDate captures enum value "ActualPickupDate"
	ServiceItemParamNameActualPickupDate ServiceItemParamName = "ActualPickupDate"

	// ServiceItemParamNameCanStandAlone captures enum value "CanStandAlone"
	ServiceItemParamNameCanStandAlone ServiceItemParamName = "CanStandAlone"

	// ServiceItemParamNameContractCode captures enum value "ContractCode"
	ServiceItemParamNameContractCode ServiceItemParamName = "ContractCode"

	// ServiceItemParamNameCubicFeetBilled captures enum value "CubicFeetBilled"
	ServiceItemParamNameCubicFeetBilled ServiceItemParamName = "CubicFeetBilled"

	// ServiceItemParamNameCubicFeetCrating captures enum value "CubicFeetCrating"
	ServiceItemParamNameCubicFeetCrating ServiceItemParamName = "CubicFeetCrating"

	// ServiceItemParamNameDistanceZip3 captures enum value "DistanceZip3"
	ServiceItemParamNameDistanceZip3 ServiceItemParamName = "DistanceZip3"

	// ServiceItemParamNameDistanceZip5 captures enum value "DistanceZip5"
	ServiceItemParamNameDistanceZip5 ServiceItemParamName = "DistanceZip5"

	// ServiceItemParamNameDistanceZipSITDest captures enum value "DistanceZipSITDest"
	ServiceItemParamNameDistanceZipSITDest ServiceItemParamName = "DistanceZipSITDest"

	// ServiceItemParamNameDistanceZipSITOrigin captures enum value "DistanceZipSITOrigin"
	ServiceItemParamNameDistanceZipSITOrigin ServiceItemParamName = "DistanceZipSITOrigin"

	// ServiceItemParamNameEIAFuelPrice captures enum value "EIAFuelPrice"
	ServiceItemParamNameEIAFuelPrice ServiceItemParamName = "EIAFuelPrice"

	// ServiceItemParamNameFSCWeightBasedDistanceMultiplier captures enum value "FSCWeightBasedDistanceMultiplier"
	ServiceItemParamNameFSCWeightBasedDistanceMultiplier ServiceItemParamName = "FSCWeightBasedDistanceMultiplier"

	// ServiceItemParamNameMarketDest captures enum value "MarketDest"
	ServiceItemParamNameMarketDest ServiceItemParamName = "MarketDest"

	// ServiceItemParamNameMarketOrigin captures enum value "MarketOrigin"
	ServiceItemParamNameMarketOrigin ServiceItemParamName = "MarketOrigin"

	// ServiceItemParamNameMTOAvailableToPrimeAt captures enum value "MTOAvailableToPrimeAt"
	ServiceItemParamNameMTOAvailableToPrimeAt ServiceItemParamName = "MTOAvailableToPrimeAt"

	// ServiceItemParamNameNumberDaysSIT captures enum value "NumberDaysSIT"
	ServiceItemParamNameNumberDaysSIT ServiceItemParamName = "NumberDaysSIT"

	// ServiceItemParamNamePriceAreaDest captures enum value "PriceAreaDest"
	ServiceItemParamNamePriceAreaDest ServiceItemParamName = "PriceAreaDest"

	// ServiceItemParamNamePriceAreaIntlDest captures enum value "PriceAreaIntlDest"
	ServiceItemParamNamePriceAreaIntlDest ServiceItemParamName = "PriceAreaIntlDest"

	// ServiceItemParamNamePriceAreaIntlOrigin captures enum value "PriceAreaIntlOrigin"
	ServiceItemParamNamePriceAreaIntlOrigin ServiceItemParamName = "PriceAreaIntlOrigin"

	// ServiceItemParamNamePriceAreaOrigin captures enum value "PriceAreaOrigin"
	ServiceItemParamNamePriceAreaOrigin ServiceItemParamName = "PriceAreaOrigin"

	// ServiceItemParamNamePSILinehaulDom captures enum value "PSI_LinehaulDom"
	ServiceItemParamNamePSILinehaulDom ServiceItemParamName = "PSI_LinehaulDom"

	// ServiceItemParamNamePSILinehaulDomPrice captures enum value "PSI_LinehaulDomPrice"
	ServiceItemParamNamePSILinehaulDomPrice ServiceItemParamName = "PSI_LinehaulDomPrice"

	// ServiceItemParamNamePSILinehaulShort captures enum value "PSI_LinehaulShort"
	ServiceItemParamNamePSILinehaulShort ServiceItemParamName = "PSI_LinehaulShort"

	// ServiceItemParamNamePSILinehaulShortPrice captures enum value "PSI_LinehaulShortPrice"
	ServiceItemParamNamePSILinehaulShortPrice ServiceItemParamName = "PSI_LinehaulShortPrice"

	// ServiceItemParamNamePSIPackingDom captures enum value "PSI_PackingDom"
	ServiceItemParamNamePSIPackingDom ServiceItemParamName = "PSI_PackingDom"

	// ServiceItemParamNamePSIPackingDomPrice captures enum value "PSI_PackingDomPrice"
	ServiceItemParamNamePSIPackingDomPrice ServiceItemParamName = "PSI_PackingDomPrice"

	// ServiceItemParamNamePSIPackingHHGIntl captures enum value "PSI_PackingHHGIntl"
	ServiceItemParamNamePSIPackingHHGIntl ServiceItemParamName = "PSI_PackingHHGIntl"

	// ServiceItemParamNamePSIPackingHHGIntlPrice captures enum value "PSI_PackingHHGIntlPrice"
	ServiceItemParamNamePSIPackingHHGIntlPrice ServiceItemParamName = "PSI_PackingHHGIntlPrice"

	// ServiceItemParamNamePSIPriceDomDest captures enum value "PSI_PriceDomDest"
	ServiceItemParamNamePSIPriceDomDest ServiceItemParamName = "PSI_PriceDomDest"

	// ServiceItemParamNamePSIPriceDomDestPrice captures enum value "PSI_PriceDomDestPrice"
	ServiceItemParamNamePSIPriceDomDestPrice ServiceItemParamName = "PSI_PriceDomDestPrice"

	// ServiceItemParamNamePSIPriceDomOrigin captures enum value "PSI_PriceDomOrigin"
	ServiceItemParamNamePSIPriceDomOrigin ServiceItemParamName = "PSI_PriceDomOrigin"

	// ServiceItemParamNamePSIPriceDomOriginPrice captures enum value "PSI_PriceDomOriginPrice"
	ServiceItemParamNamePSIPriceDomOriginPrice ServiceItemParamName = "PSI_PriceDomOriginPrice"

	// ServiceItemParamNamePSIShippingLinehaulIntlCO captures enum value "PSI_ShippingLinehaulIntlCO"
	ServiceItemParamNamePSIShippingLinehaulIntlCO ServiceItemParamName = "PSI_ShippingLinehaulIntlCO"

	// ServiceItemParamNamePSIShippingLinehaulIntlCOPrice captures enum value "PSI_ShippingLinehaulIntlCOPrice"
	ServiceItemParamNamePSIShippingLinehaulIntlCOPrice ServiceItemParamName = "PSI_ShippingLinehaulIntlCOPrice"

	// ServiceItemParamNamePSIShippingLinehaulIntlOC captures enum value "PSI_ShippingLinehaulIntlOC"
	ServiceItemParamNamePSIShippingLinehaulIntlOC ServiceItemParamName = "PSI_ShippingLinehaulIntlOC"

	// ServiceItemParamNamePSIShippingLinehaulIntlOCPrice captures enum value "PSI_ShippingLinehaulIntlOCPrice"
	ServiceItemParamNamePSIShippingLinehaulIntlOCPrice ServiceItemParamName = "PSI_ShippingLinehaulIntlOCPrice"

	// ServiceItemParamNamePSIShippingLinehaulIntlOO captures enum value "PSI_ShippingLinehaulIntlOO"
	ServiceItemParamNamePSIShippingLinehaulIntlOO ServiceItemParamName = "PSI_ShippingLinehaulIntlOO"

	// ServiceItemParamNamePSIShippingLinehaulIntlOOPrice captures enum value "PSI_ShippingLinehaulIntlOOPrice"
	ServiceItemParamNamePSIShippingLinehaulIntlOOPrice ServiceItemParamName = "PSI_ShippingLinehaulIntlOOPrice"

	// ServiceItemParamNameRateAreaNonStdDest captures enum value "RateAreaNonStdDest"
	ServiceItemParamNameRateAreaNonStdDest ServiceItemParamName = "RateAreaNonStdDest"

	// ServiceItemParamNameRateAreaNonStdOrigin captures enum value "RateAreaNonStdOrigin"
	ServiceItemParamNameRateAreaNonStdOrigin ServiceItemParamName = "RateAreaNonStdOrigin"

	// ServiceItemParamNameRequestedPickupDate captures enum value "RequestedPickupDate"
	ServiceItemParamNameRequestedPickupDate ServiceItemParamName = "RequestedPickupDate"

	// ServiceItemParamNameServiceAreaDest captures enum value "ServiceAreaDest"
	ServiceItemParamNameServiceAreaDest ServiceItemParamName = "ServiceAreaDest"

	// ServiceItemParamNameServiceAreaOrigin captures enum value "ServiceAreaOrigin"
	ServiceItemParamNameServiceAreaOrigin ServiceItemParamName = "ServiceAreaOrigin"

	// ServiceItemParamNameServicesScheduleDest captures enum value "ServicesScheduleDest"
	ServiceItemParamNameServicesScheduleDest ServiceItemParamName = "ServicesScheduleDest"

	// ServiceItemParamNameServicesScheduleOrigin captures enum value "ServicesScheduleOrigin"
	ServiceItemParamNameServicesScheduleOrigin ServiceItemParamName = "ServicesScheduleOrigin"

	// ServiceItemParamNameSITScheduleDest captures enum value "SITScheduleDest"
	ServiceItemParamNameSITScheduleDest ServiceItemParamName = "SITScheduleDest"

	// ServiceItemParamNameSITScheduleOrigin captures enum value "SITScheduleOrigin"
	ServiceItemParamNameSITScheduleOrigin ServiceItemParamName = "SITScheduleOrigin"

	// ServiceItemParamNameWeightActual captures enum value "WeightActual"
	ServiceItemParamNameWeightActual ServiceItemParamName = "WeightActual"

	// ServiceItemParamNameWeightBilledActual captures enum value "WeightBilledActual"
	ServiceItemParamNameWeightBilledActual ServiceItemParamName = "WeightBilledActual"

	// ServiceItemParamNameWeightEstimated captures enum value "WeightEstimated"
	ServiceItemParamNameWeightEstimated ServiceItemParamName = "WeightEstimated"

	// ServiceItemParamNameZipDestAddress captures enum value "ZipDestAddress"
	ServiceItemParamNameZipDestAddress ServiceItemParamName = "ZipDestAddress"

	// ServiceItemParamNameZipPickupAddress captures enum value "ZipPickupAddress"
	ServiceItemParamNameZipPickupAddress ServiceItemParamName = "ZipPickupAddress"

	// ServiceItemParamNameZipSITAddress captures enum value "ZipSITAddress"
	ServiceItemParamNameZipSITAddress ServiceItemParamName = "ZipSITAddress"
)

// for schema
var serviceItemParamNameEnum []interface{}

func init() {
	var res []ServiceItemParamName
	if err := json.Unmarshal([]byte(`["ActualPickupDate","CanStandAlone","ContractCode","CubicFeetBilled","CubicFeetCrating","DistanceZip3","DistanceZip5","DistanceZipSITDest","DistanceZipSITOrigin","EIAFuelPrice","FSCWeightBasedDistanceMultiplier","MarketDest","MarketOrigin","MTOAvailableToPrimeAt","NumberDaysSIT","PriceAreaDest","PriceAreaIntlDest","PriceAreaIntlOrigin","PriceAreaOrigin","PSI_LinehaulDom","PSI_LinehaulDomPrice","PSI_LinehaulShort","PSI_LinehaulShortPrice","PSI_PackingDom","PSI_PackingDomPrice","PSI_PackingHHGIntl","PSI_PackingHHGIntlPrice","PSI_PriceDomDest","PSI_PriceDomDestPrice","PSI_PriceDomOrigin","PSI_PriceDomOriginPrice","PSI_ShippingLinehaulIntlCO","PSI_ShippingLinehaulIntlCOPrice","PSI_ShippingLinehaulIntlOC","PSI_ShippingLinehaulIntlOCPrice","PSI_ShippingLinehaulIntlOO","PSI_ShippingLinehaulIntlOOPrice","RateAreaNonStdDest","RateAreaNonStdOrigin","RequestedPickupDate","ServiceAreaDest","ServiceAreaOrigin","ServicesScheduleDest","ServicesScheduleOrigin","SITScheduleDest","SITScheduleOrigin","WeightActual","WeightBilledActual","WeightEstimated","ZipDestAddress","ZipPickupAddress","ZipSITAddress"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceItemParamNameEnum = append(serviceItemParamNameEnum, v)
	}
}

func (m ServiceItemParamName) validateServiceItemParamNameEnum(path, location string, value ServiceItemParamName) error {
	if err := validate.EnumCase(path, location, value, serviceItemParamNameEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this service item param name
func (m ServiceItemParamName) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateServiceItemParamNameEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
