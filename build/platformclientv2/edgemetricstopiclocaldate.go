package platformclientv2
import (
	"encoding/json"
)

// Edgemetricstopiclocaldate
type Edgemetricstopiclocaldate struct { 
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
func (o *Edgemetricstopiclocaldate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
