package platformclientv2
import (
	"encoding/json"
)

// Timeinterval
type Timeinterval struct { 
	// Months
	Months *int32 `json:"months,omitempty"`


	// Weeks
	Weeks *int32 `json:"weeks,omitempty"`


	// Days
	Days *int32 `json:"days,omitempty"`


	// Hours
	Hours *int32 `json:"hours,omitempty"`

}

// String returns a JSON representation of the model
func (o *Timeinterval) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
