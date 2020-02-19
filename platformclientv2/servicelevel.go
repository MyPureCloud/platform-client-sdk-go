package platformclientv2
import (
	"encoding/json"
)

// Servicelevel
type Servicelevel struct { 
	// Percentage
	Percentage *float64 `json:"percentage,omitempty"`


	// DurationMs
	DurationMs *int64 `json:"durationMs,omitempty"`

}

// String returns a JSON representation of the model
func (o *Servicelevel) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
