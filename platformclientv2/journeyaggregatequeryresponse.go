package platformclientv2
import (
	"encoding/json"
)

// Journeyaggregatequeryresponse
type Journeyaggregatequeryresponse struct { 
	// Results
	Results *[]Journeyaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Journeyaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
