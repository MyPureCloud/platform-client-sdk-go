package platformclientv2
import (
	"encoding/json"
)

// Routingconversationattributes
type Routingconversationattributes struct { 
	// Priority
	Priority *int `json:"priority,omitempty"`

}

// String returns a JSON representation of the model
func (o *Routingconversationattributes) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
