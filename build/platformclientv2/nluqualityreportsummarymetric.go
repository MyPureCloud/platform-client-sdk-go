package platformclientv2
import (
	"encoding/json"
)

// Nluqualityreportsummarymetric
type Nluqualityreportsummarymetric struct { 
	// Name - The name of the metric. e.g. recall, f1_score
	Name *string `json:"name,omitempty"`


	// Value - The value of the metric
	Value *float32 `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nluqualityreportsummarymetric) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
