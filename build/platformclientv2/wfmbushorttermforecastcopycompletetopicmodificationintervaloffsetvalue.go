package platformclientv2
import (
	"encoding/json"
)

// Wfmbushorttermforecastcopycompletetopicmodificationintervaloffsetvalue
type Wfmbushorttermforecastcopycompletetopicmodificationintervaloffsetvalue struct { 
	// IntervalIndex
	IntervalIndex *int `json:"intervalIndex,omitempty"`


	// Value
	Value *float32 `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastcopycompletetopicmodificationintervaloffsetvalue) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
