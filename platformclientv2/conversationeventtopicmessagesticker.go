package platformclientv2
import (
	"encoding/json"
)

// Conversationeventtopicmessagesticker
type Conversationeventtopicmessagesticker struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicmessagesticker) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
