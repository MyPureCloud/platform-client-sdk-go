package platformclientv2
import (
	"encoding/json"
)

// Securesessionentitylisting
type Securesessionentitylisting struct { 
	// Entities
	Entities *[]Securesession `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Securesessionentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
