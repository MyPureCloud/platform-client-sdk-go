package platformclientv2
import (
	"encoding/json"
)

// Surveyaggregationview
type Surveyaggregationview struct { 
	// Target - Target metric name
	Target *string `json:"target,omitempty"`


	// Name - A unique name for this view. Must be distinct from other views and built-in metric names.
	Name *string `json:"name,omitempty"`


	// Function - Type of view you wish to create
	Function *string `json:"function,omitempty"`


	// VarRange - Range of numbers for slicing up data
	VarRange *Aggregationrange `json:"range,omitempty"`

}

// String returns a JSON representation of the model
func (o *Surveyaggregationview) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
