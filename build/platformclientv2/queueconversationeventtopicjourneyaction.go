package platformclientv2
import (
	"encoding/json"
)

// Queueconversationeventtopicjourneyaction
type Queueconversationeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
