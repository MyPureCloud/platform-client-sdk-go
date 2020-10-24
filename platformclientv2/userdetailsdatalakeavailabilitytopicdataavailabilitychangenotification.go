package platformclientv2
import (
	"encoding/json"
)

// Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification
type Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification struct { 
	// DataAvailabilityDate
	DataAvailabilityDate *Userdetailsdatalakeavailabilitytopicdatetime `json:"dataAvailabilityDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
