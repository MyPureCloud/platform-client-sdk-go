package platformclientv2
import (
	"encoding/json"
)

// Surveyaggregatequeryresponse
type Surveyaggregatequeryresponse struct { 
	// Results
	Results *[]Surveyaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Surveyaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
