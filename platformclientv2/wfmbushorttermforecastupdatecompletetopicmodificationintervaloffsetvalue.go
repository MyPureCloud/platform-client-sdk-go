package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalue
type Wfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalue struct { 
	// IntervalIndex
	IntervalIndex *int `json:"intervalIndex,omitempty"`


	// Value
	Value *float32 `json:"value,omitempty"`

}

func (o *Wfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalue
	
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

func (o *Wfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalue) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalueMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalueMap)
	if err != nil {
		return err
	}
	
	if IntervalIndex, ok := WfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalueMap["intervalIndex"].(float64); ok {
		IntervalIndexInt := int(IntervalIndex)
		o.IntervalIndex = &IntervalIndexInt
	}
	
	if Value, ok := WfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalueMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastupdatecompletetopicmodificationintervaloffsetvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
