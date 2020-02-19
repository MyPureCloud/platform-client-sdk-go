package platformclientv2
import (
	"encoding/json"
)

// Evaluationaggregatequeryresponse
type Evaluationaggregatequeryresponse struct { 
	// Results
	Results *[]Evaluationaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluationaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
