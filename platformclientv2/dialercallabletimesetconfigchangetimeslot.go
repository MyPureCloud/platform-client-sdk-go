package platformclientv2
import (
	"encoding/json"
)

// Dialercallabletimesetconfigchangetimeslot
type Dialercallabletimesetconfigchangetimeslot struct { 
	// StartTime
	StartTime *string `json:"startTime,omitempty"`


	// StopTime
	StopTime *string `json:"stopTime,omitempty"`


	// Day
	Day *int `json:"day,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercallabletimesetconfigchangetimeslot) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
