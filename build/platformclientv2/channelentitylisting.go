package platformclientv2
import (
	"encoding/json"
)

// Channelentitylisting
type Channelentitylisting struct { 
	// Entities
	Entities *[]Channel `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Channelentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
