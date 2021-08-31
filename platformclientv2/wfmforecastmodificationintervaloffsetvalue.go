package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmforecastmodificationintervaloffsetvalue - Override the value of a single interval in a forecast
type Wfmforecastmodificationintervaloffsetvalue struct { 
	// IntervalIndex - The number of 15 minute intervals past referenceStartDate to which to apply this modification
	IntervalIndex *int `json:"intervalIndex,omitempty"`


	// Value - The value to set for the given interval
	Value *float64 `json:"value,omitempty"`

}

func (o *Wfmforecastmodificationintervaloffsetvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmforecastmodificationintervaloffsetvalue
	
	return json.Marshal(&struct { 
		IntervalIndex *int `json:"intervalIndex,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		*Alias
	}{ 
		IntervalIndex: o.IntervalIndex,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmforecastmodificationintervaloffsetvalue) UnmarshalJSON(b []byte) error {
	var WfmforecastmodificationintervaloffsetvalueMap map[string]interface{}
	err := json.Unmarshal(b, &WfmforecastmodificationintervaloffsetvalueMap)
	if err != nil {
		return err
	}
	
	if IntervalIndex, ok := WfmforecastmodificationintervaloffsetvalueMap["intervalIndex"].(float64); ok {
		IntervalIndexInt := int(IntervalIndex)
		o.IntervalIndex = &IntervalIndexInt
	}
	
	if Value, ok := WfmforecastmodificationintervaloffsetvalueMap["value"].(float64); ok {
		o.Value = &Value
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmforecastmodificationintervaloffsetvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
