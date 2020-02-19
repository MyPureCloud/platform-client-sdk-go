package platformclientv2
import (
	"encoding/json"
)

// Copyshorttermforecastrequest - Create a new short term forecast by copying an existing forecast
type Copyshorttermforecastrequest struct { 
	// WeekDate - The first day of the short term forecast in yyyy-MM-dd format.  Must be the management unit's start day of week
	WeekDate *string `json:"weekDate,omitempty"`


	// Description - Description for the new forecast
	Description *string `json:"description,omitempty"`

}

// String returns a JSON representation of the model
func (o *Copyshorttermforecastrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
