package platformclientv2
import (
	"encoding/json"
)

// Journeyaggregatedatacontainer
type Journeyaggregatedatacontainer struct { 
	// Group - A mapping from dimension to value
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Statisticalresponse `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Journeyaggregatedatacontainer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
