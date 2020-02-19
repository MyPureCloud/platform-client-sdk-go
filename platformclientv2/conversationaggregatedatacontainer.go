package platformclientv2
import (
	"encoding/json"
)

// Conversationaggregatedatacontainer
type Conversationaggregatedatacontainer struct { 
	// Group - A mapping from dimension to value
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Statisticalresponse `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationaggregatedatacontainer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
