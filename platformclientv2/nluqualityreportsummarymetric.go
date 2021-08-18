package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluqualityreportsummarymetric
type Nluqualityreportsummarymetric struct { 
	// Name - The name of the metric. e.g. recall, f1_score
	Name *string `json:"name,omitempty"`


	// Value - The value of the metric
	Value *float32 `json:"value,omitempty"`

}

func (u *Nluqualityreportsummarymetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluqualityreportsummarymetric

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nluqualityreportsummarymetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
