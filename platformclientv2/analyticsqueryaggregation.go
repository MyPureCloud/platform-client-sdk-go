package platformclientv2
import (
	"encoding/json"
)

// Analyticsqueryaggregation
type Analyticsqueryaggregation struct { 
	// VarType - Optional type, can usually be inferred
	VarType *string `json:"type,omitempty"`


	// Dimension - For use with termFrequency aggregations
	Dimension *string `json:"dimension,omitempty"`


	// Metric - For use with numericRange aggregations
	Metric *string `json:"metric,omitempty"`


	// Size - For use with termFrequency aggregations
	Size *int `json:"size,omitempty"`


	// Ranges - For use with numericRange aggregations
	Ranges *[]Aggregationrange `json:"ranges,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsqueryaggregation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
