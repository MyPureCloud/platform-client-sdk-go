package platformclientv2
import (
	"encoding/json"
)

// Conversationaggregatequeryresponse
type Conversationaggregatequeryresponse struct { 
	// Results
	Results *[]Conversationaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
