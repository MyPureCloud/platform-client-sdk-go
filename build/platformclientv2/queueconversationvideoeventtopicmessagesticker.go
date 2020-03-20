package platformclientv2
import (
	"encoding/json"
)

// Queueconversationvideoeventtopicmessagesticker
type Queueconversationvideoeventtopicmessagesticker struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicmessagesticker) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
