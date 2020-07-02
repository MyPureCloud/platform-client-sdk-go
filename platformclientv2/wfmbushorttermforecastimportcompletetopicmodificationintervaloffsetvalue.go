package platformclientv2
import (
	"encoding/json"
)

// Wfmbushorttermforecastimportcompletetopicmodificationintervaloffsetvalue
type Wfmbushorttermforecastimportcompletetopicmodificationintervaloffsetvalue struct { 
	// IntervalIndex
	IntervalIndex *int32 `json:"intervalIndex,omitempty"`


	// Value
	Value *float32 `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicmodificationintervaloffsetvalue) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
