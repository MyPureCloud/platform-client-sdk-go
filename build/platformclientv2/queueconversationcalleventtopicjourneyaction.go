package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcalleventtopicjourneyaction
type Queueconversationcalleventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationcalleventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
