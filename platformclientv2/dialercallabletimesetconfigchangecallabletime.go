package platformclientv2
import (
	"encoding/json"
)

// Dialercallabletimesetconfigchangecallabletime
type Dialercallabletimesetconfigchangecallabletime struct { 
	// TimeSlots
	TimeSlots *[]Dialercallabletimesetconfigchangetimeslot `json:"timeSlots,omitempty"`


	// TimeZoneId
	TimeZoneId *string `json:"timeZoneId,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercallabletimesetconfigchangecallabletime) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
