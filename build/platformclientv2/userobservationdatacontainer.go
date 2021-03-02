package platformclientv2
import (
	"encoding/json"
)

// Userobservationdatacontainer
type Userobservationdatacontainer struct { 
	// Group - A mapping from dimension to value
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Observationmetricdata `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userobservationdatacontainer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
