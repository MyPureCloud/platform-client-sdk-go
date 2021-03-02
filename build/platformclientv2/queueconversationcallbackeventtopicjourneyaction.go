package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcallbackeventtopicjourneyaction
type Queueconversationcallbackeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationcallbackeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
