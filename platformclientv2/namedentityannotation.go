package platformclientv2
import (
	"encoding/json"
)

// Namedentityannotation
type Namedentityannotation struct { 
	// Name - The name of the annotated named entity.
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Namedentityannotation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
