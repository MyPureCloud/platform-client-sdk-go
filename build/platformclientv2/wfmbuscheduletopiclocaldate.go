package platformclientv2
import (
	"encoding/json"
)

// Wfmbuscheduletopiclocaldate
type Wfmbuscheduletopiclocaldate struct { 
	// Year
	Year *int `json:"year,omitempty"`


	// Month
	Month *int `json:"month,omitempty"`


	// Day
	Day *int `json:"day,omitempty"`


	// LeapYear
	LeapYear *bool `json:"leapYear,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopiclocaldate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
