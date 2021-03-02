package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcobrowseeventtopicjourneyaction
type Queueconversationcobrowseeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationcobrowseeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
