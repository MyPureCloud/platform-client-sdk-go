package platformclientv2
import (
	"encoding/json"
)

// Generateweekschedulerequest - Request to generate a week schedule
type Generateweekschedulerequest struct { 
	// Description - Description for the generated week schedule
	Description *string `json:"description,omitempty"`


	// ShortTermForecastId - ID of short term forecast used for schedule generation
	ShortTermForecastId *string `json:"shortTermForecastId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Generateweekschedulerequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
