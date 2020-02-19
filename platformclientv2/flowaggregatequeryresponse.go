package platformclientv2
import (
	"encoding/json"
)

// Flowaggregatequeryresponse
type Flowaggregatequeryresponse struct { 
	// Results
	Results *[]Flowaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
