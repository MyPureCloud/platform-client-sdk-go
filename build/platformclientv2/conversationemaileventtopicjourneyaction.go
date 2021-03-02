package platformclientv2
import (
	"encoding/json"
)

// Conversationemaileventtopicjourneyaction
type Conversationemaileventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationemaileventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
