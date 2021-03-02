package platformclientv2
import (
	"encoding/json"
)

// Wfmforecastmodificationintervaloffsetvalue - Override the value of a single interval in a forecast
type Wfmforecastmodificationintervaloffsetvalue struct { 
	// IntervalIndex - The number of 15 minute intervals past referenceStartDate to which to apply this modification
	IntervalIndex *int `json:"intervalIndex,omitempty"`


	// Value - The value to set for the given interval
	Value *float64 `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmforecastmodificationintervaloffsetvalue) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
