package platformclientv2
import (
	"encoding/json"
)

// Conversationeventtopicjourneyaction
type Conversationeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
