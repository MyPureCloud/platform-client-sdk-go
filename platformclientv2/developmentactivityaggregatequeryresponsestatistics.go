package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivityaggregatequeryresponsestatistics
type Developmentactivityaggregatequeryresponsestatistics struct { 
	// Count - The count for this metric
	Count *int `json:"count,omitempty"`


	// Min - The minimum value in this metric
	Min *int `json:"min,omitempty"`


	// Max - The maximum value in this metric
	Max *int `json:"max,omitempty"`


	// Sum - The total of the values for this metric
	Sum *int `json:"sum,omitempty"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsestatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
