package platformclientv2
import (
	"encoding/json"
)

// Queueconversationeventtopicmessagesticker
type Queueconversationeventtopicmessagesticker struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicmessagesticker) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
