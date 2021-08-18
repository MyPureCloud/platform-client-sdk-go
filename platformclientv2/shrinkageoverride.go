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

func (u *Shrinkageoverride) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shrinkageoverride

	

	return json.Marshal(&struct { 
		IntervalIndex *int `json:"intervalIndex,omitempty"`
		
		ShrinkagePercent *float64 `json:"shrinkagePercent,omitempty"`
		*Alias
	}{ 
		IntervalIndex: u.IntervalIndex,
		
		ShrinkagePercent: u.ShrinkagePercent,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Shrinkageoverride) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
