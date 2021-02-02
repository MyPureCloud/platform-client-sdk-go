package platformclientv2
import (
	"encoding/json"
)

// Aggregationresult
type Aggregationresult struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Dimension - For termFrequency aggregations
	Dimension *string `json:"dimension,omitempty"`


	// Metric - For numericRange aggregations
	Metric *string `json:"metric,omitempty"`


	// Count
	Count *int `json:"count,omitempty"`


	// Results
	Results *[]Aggregationresultentry `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Aggregationresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
