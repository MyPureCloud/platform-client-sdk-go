package platformclientv2
import (
	"encoding/json"
)

// Wfmbulkshifttradestateupdatenotificationtopiclocaldate
type Wfmbulkshifttradestateupdatenotificationtopiclocaldate struct { 
	// Year
	Year *int32 `json:"year,omitempty"`


	// Month
	Month *int32 `json:"month,omitempty"`


	// Day
	Day *int32 `json:"day,omitempty"`


	// LeapYear
	LeapYear *bool `json:"leapYear,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopiclocaldate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
