package platformclientv2
import (
	"encoding/json"
)

// Conversationmessageeventtopicmessagesticker
type Conversationmessageeventtopicmessagesticker struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicmessagesticker) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
