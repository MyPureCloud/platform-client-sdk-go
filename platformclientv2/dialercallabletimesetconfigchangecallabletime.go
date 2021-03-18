package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
