package platformclientv2
import (
	"encoding/json"
)

// Conversationchateventtopicjourneyaction
type Conversationchateventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationchateventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationchateventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
