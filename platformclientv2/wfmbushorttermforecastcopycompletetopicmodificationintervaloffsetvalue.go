package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Wfmbushorttermforecastcopycompletetopicmodificationintervaloffsetvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastcopycompletetopicmodificationintervaloffsetvalue
	
	return json.Marshal(&struct { 
		IntervalIndex *int `json:"intervalIndex,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		*Alias
	}{ 
		IntervalIndex: o.IntervalIndex,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbushorttermforecastcopycompletetopicmodificationintervaloffsetvalue) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecastcopycompletetopicmodificationintervaloffsetvalueMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecastcopycompletetopicmodificationintervaloffsetvalueMap)
	if err != nil {
		return err
	}
	
	if IntervalIndex, ok := WfmbushorttermforecastcopycompletetopicmodificationintervaloffsetvalueMap["intervalIndex"].(float64); ok {
		IntervalIndexInt := int(IntervalIndex)
		o.IntervalIndex = &IntervalIndexInt
	}
	
	if Value, ok := WfmbushorttermforecastcopycompletetopicmodificationintervaloffsetvalueMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastcopycompletetopicmodificationintervaloffsetvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
