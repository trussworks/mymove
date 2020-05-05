package serviceparamvaluelookups

import (
	"fmt"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/gen/primemessages"
)

// ServiceItemParamKeyData contains service item parameter keys
type ServiceItemParamKeyData struct {
	db                 *pop.Connection
	lookups            map[string]ServiceItemParamKeyLookup
	PayloadServiceItem primemessages.ServiceItem
	MTOServiceItemID   uuid.UUID
	PaymentRequestID   uuid.UUID
	MoveTaskOrderID    uuid.UUID
}

// ServiceItemParamKeyLookup does lookup on service item parameter keys
type ServiceItemParamKeyLookup interface {
	lookup(keyData *ServiceItemParamKeyData) (string, error)
}

// ServiceParamLookupInitialize initializes service parameter lookup
func ServiceParamLookupInitialize(
	db *pop.Connection,
	mtoServiceItemID uuid.UUID,
	paymentRequestID uuid.UUID,
	moveTaskOrderID uuid.UUID,
) *ServiceItemParamKeyData {

	s := ServiceItemParamKeyData{
		db:               db,
		lookups:          make(map[string]ServiceItemParamKeyLookup),
		MTOServiceItemID: mtoServiceItemID,
		PaymentRequestID: paymentRequestID,
		MoveTaskOrderID:  moveTaskOrderID,
	}

	s.lookups["RequestedPickupDate"] = RequestedPickupDateLookup{}
	s.lookups["WeightBilledActual"] = NotImplementedLookup{}
	s.lookups["WeightActual"] = NotImplementedLookup{}
	s.lookups["WeightEstimated"] = NotImplementedLookup{}
	s.lookups["DistanceZip3"] = NotImplementedLookup{}
	s.lookups["ZipPickupAddress"] = NotImplementedLookup{}
	s.lookups["ZipDestAddress"] = NotImplementedLookup{}
	s.lookups["ServiceAreaOrigin"] = NotImplementedLookup{}
	//s.lookups["RequestedPickupDate"] = RequestedPickupDateLookup{}
	//s.lookups["WeightBilledActual"] = WeightBilledActualLookup{}
	//s.lookups["WeightActual"] = WeightActualLookup{}
	//s.lookups["WeightEstimated"] = WeightEstimatedLookup{}
	//s.lookups["DistanceZip3"] = DistanceZip3Lookup{}
	//s.lookups["ZipPickupAddress"] = ZipPickupAddressLookup{}
	//s.lookups["ZipDestAddress"] = ZipDestAddressLookup{}
	//s.lookups["ServiceAreaOrigin"] = ServiceAreaOriginLookup{}

	return &s
}

// ServiceParamValue returns a service parameter value from a key
func (s *ServiceItemParamKeyData) ServiceParamValue(key string) (string, error) {
	if lookup, ok := s.lookups[key]; ok {
		value, err := lookup.lookup(s)
		if err != nil {
			return "", fmt.Errorf(" failed ServiceParamValue %sLookup with error %w", key, err)
		}
		return value, nil
	}
	return "", fmt.Errorf("  ServiceParamValue <%sLookup> does not exist for key: <%s>", key, key)
}
