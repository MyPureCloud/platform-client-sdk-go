package platformclientv2
import (
	"encoding/json"
)

// Conversationcalleventtopicjourneyaction
type Conversationcalleventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationcalleventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
