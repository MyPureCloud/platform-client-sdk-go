package platformclientv2
import (
	"encoding/json"
)

// Availabletopicentitylisting
type Availabletopicentitylisting struct { 
	// Entities
	Entities *[]Availabletopic `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Availabletopicentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
