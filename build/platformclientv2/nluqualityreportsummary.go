package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Nluqualityreportsummary
type Nluqualityreportsummary struct { 
	// Metrics - The list of metrics in the summary
	Metrics *[]Nluqualityreportsummarymetric `json:"metrics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nluqualityreportsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
