package platformclientv2
import (
	"encoding/json"
)

// Queueconversationchateventtopicjourneyaction
type Queueconversationchateventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationchateventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationchateventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
