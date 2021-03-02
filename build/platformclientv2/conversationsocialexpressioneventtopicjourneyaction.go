package platformclientv2
import (
	"encoding/json"
)

// Conversationsocialexpressioneventtopicjourneyaction
type Conversationsocialexpressioneventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationsocialexpressioneventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
