package platformclientv2
import (
	"encoding/json"
)

// Generatebuforecastrequest
type Generatebuforecastrequest struct { 
	// Description - The description for the forecast
	Description *string `json:"description,omitempty"`


	// WeekCount - The number of weeks this forecast covers
	WeekCount *int `json:"weekCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Generatebuforecastrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
