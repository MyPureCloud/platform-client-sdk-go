package platformclientv2
import (
	"encoding/json"
)

// Queueconversationsocialexpressioneventtopicjourneyaction
type Queueconversationsocialexpressioneventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationsocialexpressioneventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
