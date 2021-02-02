package platformclientv2
import (
	"encoding/json"
)

// Wfmbuintradaydataupdatetopiclocaldate
type Wfmbuintradaydataupdatetopiclocaldate struct { 
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
func (o *Wfmbuintradaydataupdatetopiclocaldate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
