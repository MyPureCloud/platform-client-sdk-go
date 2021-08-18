package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluqualityreportsummary
type Nluqualityreportsummary struct { 
	// Metrics - The list of metrics in the summary
	Metrics *[]Nluqualityreportsummarymetric `json:"metrics,omitempty"`

}

func (u *Nluqualityreportsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluqualityreportsummary

	

	return json.Marshal(&struct { 
		Metrics *[]Nluqualityreportsummarymetric `json:"metrics,omitempty"`
		*Alias
	}{ 
		Metrics: u.Metrics,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nluqualityreportsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
