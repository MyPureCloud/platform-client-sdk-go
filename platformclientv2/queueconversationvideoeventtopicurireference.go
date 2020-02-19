package platformclientv2
import (
	"encoding/json"
)

// Queueconversationvideoeventtopicurireference
type Queueconversationvideoeventtopicurireference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicurireference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
