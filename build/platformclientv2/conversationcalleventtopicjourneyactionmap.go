package platformclientv2
import (
	"encoding/json"
)

// Conversationcalleventtopicjourneyactionmap
type Conversationcalleventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int32 `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
