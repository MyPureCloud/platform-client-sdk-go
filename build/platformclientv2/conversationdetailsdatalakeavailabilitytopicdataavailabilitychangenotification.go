package platformclientv2
import (
	"encoding/json"
)

// Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification
type Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification struct { 
	// DataAvailabilityDate
	DataAvailabilityDate *Conversationdetailsdatalakeavailabilitytopicdatetime `json:"dataAvailabilityDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
