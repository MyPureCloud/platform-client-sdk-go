package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Reportinginterval
type Reportinginterval struct { 
	// IntervalType - The granularity of the reporting interval period
	IntervalType *string `json:"intervalType,omitempty"`


	// IntervalValue - The value of the reporting interval period for a given interval type
	IntervalValue *int `json:"intervalValue,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reportinginterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
