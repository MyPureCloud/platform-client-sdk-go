package platformclientv2
import (
	"encoding/json"
)

// Queueconversationchateventtopicjourneyactionmap
type Queueconversationchateventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int32 `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationchateventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
