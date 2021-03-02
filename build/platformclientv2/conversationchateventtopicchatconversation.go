package platformclientv2
import (
	"encoding/json"
)

// Conversationchateventtopicchatconversation
type Conversationchateventtopicchatconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Conversationchateventtopicchatmediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationchateventtopicchatconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
