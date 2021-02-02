package platformclientv2
import (
	"encoding/json"
)

// Conversationchateventtopicjourneyactionmap
type Conversationchateventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationchateventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
