package platformclientv2
import (
	"encoding/json"
)

// Nluqualityreportsummary
type Nluqualityreportsummary struct { 
	// Metrics - The list of metrics in the summary
	Metrics *[]Nluqualityreportsummarymetric `json:"metrics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nluqualityreportsummary) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
