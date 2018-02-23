package awardqueue

import (
	"fmt"
	"math"

	"github.com/markbates/pop"

	"github.com/transcom/mymove/pkg/models"
)

var db *pop.Connection

const numQualBands = 4

type qualityBand []models.TSPWithBVSCount
type qualityBands []qualityBand

func findAllUnawardedShipments() ([]models.PossiblyAwardedShipment, error) {
	shipments, err := models.FetchAwardedShipments(db)
	return shipments, err
}

// AttemptShipmentAward will attempt to take the given Shipment and award it to
// a TSP.
func AttemptShipmentAward(shipment models.PossiblyAwardedShipment) (*models.ShipmentAward, error) {
	fmt.Printf("Attempting to award shipment: %v\n", shipment.ID)

	// Query the shipment's TDL
	tdl := models.TrafficDistributionList{}
	err := db.Find(&tdl, shipment.TrafficDistributionListID)

	// Find TSPs in that TDL sorted by shipment_awards[asc] and bvs[desc]
	// tspssba stands for TSPs sorted by award
	tspsba, err := models.FetchTSPsInTDLSortByAward(db, tdl.ID)

	if len(tspsba) == 0 {
		return nil, fmt.Errorf("Cannot award. No TSPs found in TDL (%v)", tdl.ID)
	}

	var shipmentAward *models.ShipmentAward

	for _, consideredTSP := range tspsba {
		fmt.Printf("\tConsidering TSP: %s\n", consideredTSP.Name)

		tsp := models.TransportationServiceProvider{}
		if err := db.Find(&tsp, consideredTSP.ID); err == nil {
			// We found a valid TSP to award to!
			shipmentAward, err = models.CreateShipmentAward(db, shipment.ID, tsp.ID, false)
			if err == nil {
				fmt.Print("\tShipment awarded to TSP!\n")
				break
			} else {
				fmt.Printf("\tFailed to award to TSP: %v\n", err)
			}
		} else {
			fmt.Printf("\tFailed to award to TSP: %v\n", err)
		}
	}

	return shipmentAward, err
}

// getTSPsPerBand detemines how many TSPs should be assigned to each Quality Band
// If the number of TSPs in the TDL does not divide evenly into 4 bands, the remainder
// is divided from the top band down. Function takes length of TSPs array (tsp count) as arg.
func getTSPsPerBand(tspc int) []int {
	tspPerBandList := make([]int, numQualBands)
	// tspp is TSP per band
	tsppb := int(math.Floor(float64(tspc) / float64(numQualBands)))
	// assign tsppb to each band in tspPerBandList to refer to when assigning TSPs
	for i := range tspPerBandList {
		tspPerBandList[i] = tsppb
	}

	for i := 0; i < tspc%numQualBands; i++ {
		tspPerBandList[i]++
	}
	return tspPerBandList
}

func assignTSPsToBands(tsps []models.TSPWithBVSCount) qualityBands {
	tspIndex := 0
	qbs := make(qualityBands, numQualBands)
	// Determine how many TSPs should be in each band
	tsppbl := getTSPsPerBand(len(tsps))

	for i, tsppb := range tsppbl {
		for j := tspIndex; j < tspIndex+tsppb; j++ {
			qbs[i] = append(qbs[i], tsps[j])
		}
		tspIndex += tsppb
	}
	return qbs
}

func assignQualityBands() (qualityBands, error) {
	// Find TSPs in that TDL sorted bvs[desc]
	fmt.Printf("Assigning TSPs quality bands")
	// Query the shipment's TDL
	tdl := models.TrafficDistributionList{}
	// tspsbb stands for TSPs sorted by BVS
	tspsbb, err := models.FetchTSPsInTDLSortByBVS(db, tdl.ID)
	// Assign TSPs to bands and return slice of TSP slices divided by band
	return assignTSPsToBands(tspsbb), err
}

// Run will execute the award queue algorithm.
func Run(db *pop.Connection) {
	fmt.Println("TSP Award Queue running.")

	shipments, err := findAllUnawardedShipments()
	if err == nil {
		count := 0
		for _, shipment := range shipments {
			_, err = AttemptShipmentAward(shipment)
			if err != nil {
				fmt.Printf("\tFailed to award shipment: %s\n", err)
			} else {
				count++
			}
		}
		fmt.Printf("Awarded %d shipments.\n", count)
	} else {
		fmt.Printf("Failed to query for shipments: %s", err)
	}
}
