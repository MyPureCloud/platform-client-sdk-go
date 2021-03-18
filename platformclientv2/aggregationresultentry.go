package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregationresultentry
type Aggregationresultentry struct { 
	// Count
	Count *int `json:"count,omitempty"`


	// Value - For termFrequency aggregations
	Value *string `json:"value,omitempty"`


	// Gte - For numericRange aggregations
	Gte *float32 `json:"gte,omitempty"`


	// Lt - For numericRange aggregations
	Lt *float32 `json:"lt,omitempty"`

}

// String returns a JSON representation of the model
func (o *Aggregationresultentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
