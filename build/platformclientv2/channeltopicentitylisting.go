package platformclientv2
import (
	"encoding/json"
)

// Channeltopicentitylisting
type Channeltopicentitylisting struct { 
	// Entities
	Entities *[]Channeltopic `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Channeltopicentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
