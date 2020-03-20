package platformclientv2
import (
	"encoding/json"
)

// Conversationcallbackeventtopicjourneyaction
type Conversationcallbackeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationcallbackeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcallbackeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
