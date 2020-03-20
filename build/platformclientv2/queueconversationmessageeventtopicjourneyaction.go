package platformclientv2
import (
	"encoding/json"
)

// Queueconversationmessageeventtopicjourneyaction
type Queueconversationmessageeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationmessageeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
