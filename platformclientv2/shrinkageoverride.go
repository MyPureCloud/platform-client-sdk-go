package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Shrinkageoverride) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
