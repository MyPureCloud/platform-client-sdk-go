package platformclientv2
import (
	"encoding/json"
)

// Queueconversationvideoeventtopicjourneyactionmap
type Queueconversationvideoeventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
