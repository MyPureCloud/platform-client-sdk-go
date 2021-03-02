package platformclientv2
import (
	"encoding/json"
)

// Conversationvideoeventtopicjourneyaction
type Conversationvideoeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationvideoeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationvideoeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
