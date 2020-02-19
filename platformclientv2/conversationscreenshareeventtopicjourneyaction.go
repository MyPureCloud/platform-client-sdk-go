package platformclientv2
import (
	"encoding/json"
)

// Conversationscreenshareeventtopicjourneyaction
type Conversationscreenshareeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationscreenshareeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
