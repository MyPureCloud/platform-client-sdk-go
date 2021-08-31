package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shrinkageoverride
type Shrinkageoverride struct { 
	// IntervalIndex - Index of shrinkage override interval. Starting index is 0 and indexes are based on 15 minute intervals for a 7 day week
	IntervalIndex *int `json:"intervalIndex,omitempty"`


	// ShrinkagePercent - Shrinkage override percent. Setting a null value will reset the interval to the default
	ShrinkagePercent *float64 `json:"shrinkagePercent,omitempty"`

}

func (o *Shrinkageoverride) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shrinkageoverride
	
	return json.Marshal(&struct { 
		IntervalIndex *int `json:"intervalIndex,omitempty"`
		
		ShrinkagePercent *float64 `json:"shrinkagePercent,omitempty"`
		*Alias
	}{ 
		IntervalIndex: o.IntervalIndex,
		
		ShrinkagePercent: o.ShrinkagePercent,
		Alias:    (*Alias)(o),
	})
}

func (o *Shrinkageoverride) UnmarshalJSON(b []byte) error {
	var ShrinkageoverrideMap map[string]interface{}
	err := json.Unmarshal(b, &ShrinkageoverrideMap)
	if err != nil {
		return err
	}
	
	if IntervalIndex, ok := ShrinkageoverrideMap["intervalIndex"].(float64); ok {
		IntervalIndexInt := int(IntervalIndex)
		o.IntervalIndex = &IntervalIndexInt
	}
	
	if ShrinkagePercent, ok := ShrinkageoverrideMap["shrinkagePercent"].(float64); ok {
		o.ShrinkagePercent = &ShrinkagePercent
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shrinkageoverride) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
