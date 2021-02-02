package platformclientv2
import (
	"encoding/json"
)

// Queueconversationmessageeventtopicjourneyactionmap
type Queueconversationmessageeventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
