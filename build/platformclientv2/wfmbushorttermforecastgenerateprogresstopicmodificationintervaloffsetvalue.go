package platformclientv2
import (
	"encoding/json"
)

// Wfmbushorttermforecastgenerateprogresstopicmodificationintervaloffsetvalue
type Wfmbushorttermforecastgenerateprogresstopicmodificationintervaloffsetvalue struct { 
	// IntervalIndex
	IntervalIndex *int `json:"intervalIndex,omitempty"`


	// Value
	Value *float32 `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastgenerateprogresstopicmodificationintervaloffsetvalue) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
