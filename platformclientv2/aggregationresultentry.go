package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Aggregationresultentry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Aggregationresultentry

	

	return json.Marshal(&struct { 
		Count *int `json:"count,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Gte *float32 `json:"gte,omitempty"`
		
		Lt *float32 `json:"lt,omitempty"`
		*Alias
	}{ 
		Count: u.Count,
		
		Value: u.Value,
		
		Gte: u.Gte,
		
		Lt: u.Lt,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Aggregationresultentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
