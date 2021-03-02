package platformclientv2
import (
	"encoding/json"
)

// Conversationscreenshareeventtopicurireference
type Conversationscreenshareeventtopicurireference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicurireference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
