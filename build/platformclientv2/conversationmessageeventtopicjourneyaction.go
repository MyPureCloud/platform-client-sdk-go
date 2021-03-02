package platformclientv2
import (
	"encoding/json"
)

// Conversationmessageeventtopicjourneyaction
type Conversationmessageeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationmessageeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
