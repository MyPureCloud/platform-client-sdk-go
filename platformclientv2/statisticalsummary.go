package platformclientv2
import (
	"encoding/json"
)

// Statisticalsummary
type Statisticalsummary struct { 
	// Max
	Max *float32 `json:"max,omitempty"`


	// Min
	Min *float32 `json:"min,omitempty"`


	// Count
	Count *int64 `json:"count,omitempty"`


	// Sum
	Sum *float32 `json:"sum,omitempty"`


	// Current
	Current *float32 `json:"current,omitempty"`


	// Ratio
	Ratio *float32 `json:"ratio,omitempty"`


	// Numerator
	Numerator *float32 `json:"numerator,omitempty"`


	// Denominator
	Denominator *float32 `json:"denominator,omitempty"`


	// Target
	Target *float32 `json:"target,omitempty"`

}

// String returns a JSON representation of the model
func (o *Statisticalsummary) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
