package platformclientv2
import (
	"encoding/json"
)

// Conversationeventtopicjourneyactionmap
type Conversationeventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
