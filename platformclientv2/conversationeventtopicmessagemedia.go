package platformclientv2
import (
	"encoding/json"
)

// Conversationeventtopicmessagemedia
type Conversationeventtopicmessagemedia struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// MediaType
	MediaType *string `json:"mediaType,omitempty"`


	// ContentLengthBytes
	ContentLengthBytes *int32 `json:"contentLengthBytes,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicmessagemedia) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
