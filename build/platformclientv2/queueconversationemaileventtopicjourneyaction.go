package platformclientv2
import (
	"encoding/json"
)

// Queueconversationemaileventtopicjourneyaction
type Queueconversationemaileventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationemaileventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
