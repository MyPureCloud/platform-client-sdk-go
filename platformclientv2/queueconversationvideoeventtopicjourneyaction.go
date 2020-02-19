package platformclientv2
import (
	"encoding/json"
)

// Queueconversationvideoeventtopicjourneyaction
type Queueconversationvideoeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationvideoeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
