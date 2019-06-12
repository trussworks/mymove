package storageintransit

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/rateengine"
	"github.com/transcom/mymove/pkg/route"
	"github.com/transcom/mymove/pkg/unit"
)

type CreateStorageInTransitLineItems struct {
	DB      *pop.Connection
	Planner route.Planner
}

func (c CreateStorageInTransitLineItems) storageInTransitDistance(storageInTransit models.StorageInTransit, shipment models.Shipment) (*models.DistanceCalculation, error) {

	var origin models.Address
	var destination models.Address

	if storageInTransit.Location == models.StorageInTransitLocationDESTINATION {
		origin = storageInTransit.WarehouseAddress
		destination = shipment.Move.Orders.NewDutyStation.Address
		if shipment.DestinationAddressOnAcceptance != nil {
			destination = *shipment.DestinationAddressOnAcceptance
		}

	} else if storageInTransit.Location == models.StorageInTransitLocationORIGIN {
		if shipment.PickupAddress != nil {
			origin = *shipment.PickupAddress
		} else {
			return nil, errors.New("StorageInTransit PickupAddress not provided")
		}

		destination = storageInTransit.WarehouseAddress
	}

	if origin.ID == uuid.Nil {
		return nil, errors.New("StorageInTransit PickupAddress not provided")
	}

	if destination.ID == uuid.Nil {
		return nil, errors.New("StorageInTransit Destination address not provided")
	}

	distanceCalculation, err := models.NewDistanceCalculation(c.Planner, origin, destination)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating StorageInTransit DistanceCalculation model")
	}

	verrs, err := c.DB.ValidateAndSave(distanceCalculation)
	if verrs.HasAny() || err != nil {
		saveError := errors.Wrapf(err, "Error saving storage in transit distance %s", verrs.Error())
		return nil, saveError
	}

	return &distanceCalculation, nil
}

func (c CreateStorageInTransitLineItems) shipmentItemLocation(location models.StorageInTransitLocation) models.ShipmentLineItemLocation {
	if location == models.StorageInTransitLocationORIGIN {
		return models.ShipmentLineItemLocationORIGIN
	}

	if location == models.StorageInTransitLocationDESTINATION {
		return models.ShipmentLineItemLocationDESTINATION
	}

	return models.ShipmentLineItemLocationNEITHER
}

func (c CreateStorageInTransitLineItems) CreateStorageInTransitLineItems(costByShipment rateengine.CostByShipment) ([]models.ShipmentLineItem, error) {

	var lineItems []models.ShipmentLineItem
	shipment := costByShipment.Shipment
	cost := costByShipment.Cost
	now := time.Now()
	storageInTransits, err := models.FetchStorageInTransitsOnShipment(c.DB, shipment.ID)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating StorageInTransit line items")
	}

	for _, sit := range storageInTransits {

		// Only add line items for storage in transits that have been released or delivered
		if sit.Status != models.StorageInTransitStatusDELIVERED && sit.Status != models.StorageInTransitStatusRELEASED {
			continue
		}

		// Calculate distance for storage in transit
		distanceCalculation, err := c.storageInTransitDistance(sit, shipment)
		if distanceCalculation == nil || err != nil {
			return nil, err
		}
		sit.StorageInTransitDistance = *distanceCalculation
		sit.StorageInTransitDistanceID = &(*distanceCalculation).ID

		/****************************************************************************
				 * Add 210A, 210B, and 210C Shipment Line Items
				 * 210A-E = Additional flat rate costs based on distance to/from the SIT facility. These vary based on geographical schedules.
		         *
				 *
				 * Up to 30 miles: Item 210A --  SIT Pup/Del - 30 or Less Miles
				 * Up to 50 miles: Item 201A & 210B -- SIT Pup/Del 31 - 50 Miles
				 * Over 50 miles : Item 210C (Use the linehaul tables for computation of charges) -- SIT Pup/Del - Over 50 Miles
				 * Over 50 miles (Alaska only) : Item 210F (Use linehaul tables section 7 Intra-AK)
				 ****************************************************************************/

		if sit.StorageInTransitDistance.DistanceMiles > 50 {
			additionalFlateRateCItem, err := models.FetchTariff400ngItemByCode(c.DB, "210C")
			if err != nil {
				return nil, err
			}
			additionalFlateRateC := models.ShipmentLineItem{
				ShipmentID:        shipment.ID,
				Tariff400ngItemID: additionalFlateRateCItem.ID,
				Tariff400ngItem:   additionalFlateRateCItem,
				Location:          c.shipmentItemLocation(sit.Location),
				Quantity1:         unit.BaseQuantityFromInt(sit.StorageInTransitDistance.DistanceMiles),
				Status:            models.ShipmentLineItemStatusSUBMITTED,
				AmountCents:       &cost.NonLinehaulCostComputation.DestinationService.Fee,
				AppliedRate:       &cost.NonLinehaulCostComputation.DestinationService.Rate,
				SubmittedDate:     now,
			}
			lineItems = append(lineItems, additionalFlateRateC)
		} else {
			if sit.StorageInTransitDistance.DistanceMiles > 30 {
				additionalFlateRateBItem, err := models.FetchTariff400ngItemByCode(c.DB, "210B")
				if err != nil {
					return nil, err
				}
				additionalFlateRateB := models.ShipmentLineItem{
					ShipmentID:        shipment.ID,
					Tariff400ngItemID: additionalFlateRateBItem.ID,
					Tariff400ngItem:   additionalFlateRateBItem,
					Location:          c.shipmentItemLocation(sit.Location),
					Quantity1:         unit.BaseQuantityFromInt(sit.StorageInTransitDistance.DistanceMiles),
					Status:            models.ShipmentLineItemStatusSUBMITTED,
					AmountCents:       &cost.NonLinehaulCostComputation.DestinationService.Fee,
					AppliedRate:       &cost.NonLinehaulCostComputation.DestinationService.Rate,
					SubmittedDate:     now,
				}
				lineItems = append(lineItems, additionalFlateRateB)
			}

			additionalFlateRateAItem, err := models.FetchTariff400ngItemByCode(c.DB, "210A")
			if err != nil {
				return nil, err
			}
			additionalFlateRateA := models.ShipmentLineItem{
				ShipmentID:        shipment.ID,
				Tariff400ngItemID: additionalFlateRateAItem.ID,
				Tariff400ngItem:   additionalFlateRateAItem,
				Location:          c.shipmentItemLocation(sit.Location),
				Quantity1:         unit.BaseQuantityFromInt(sit.StorageInTransitDistance.DistanceMiles),
				Status:            models.ShipmentLineItemStatusSUBMITTED,
				AmountCents:       &cost.NonLinehaulCostComputation.DestinationService.Fee,
				AppliedRate:       &cost.NonLinehaulCostComputation.DestinationService.Rate,
				SubmittedDate:     now,
			}
			lineItems = append(lineItems, additionalFlateRateA)
		}

	}

	return lineItems, nil
}
