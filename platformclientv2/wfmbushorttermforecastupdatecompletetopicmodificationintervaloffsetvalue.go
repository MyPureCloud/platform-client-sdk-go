package platformclientv2
import (
	"encoding/json"
)

// Wfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalue
type Wfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalue struct { 
	// IntervalIndex
	IntervalIndex *int `json:"intervalIndex,omitempty"`


	// Value
	Value *float32 `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalue) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
