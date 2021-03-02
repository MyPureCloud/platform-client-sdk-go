package platformclientv2
import (
	"encoding/json"
)

// Queueconversationsocialexpressioneventtopicmessagemedia
type Queueconversationsocialexpressioneventtopicmessagemedia struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// MediaType
	MediaType *string `json:"mediaType,omitempty"`


	// ContentLengthBytes
	ContentLengthBytes *int `json:"contentLengthBytes,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicmessagemedia) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
