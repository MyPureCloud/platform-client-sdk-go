package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
