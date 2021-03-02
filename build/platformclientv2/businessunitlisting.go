package platformclientv2
import (
	"encoding/json"
)

// Businessunitlisting
type Businessunitlisting struct { 
	// Entities
	Entities *[]Businessunitlistitem `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Businessunitlisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
