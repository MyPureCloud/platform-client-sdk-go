package platformclientv2
import (
	"encoding/json"
)

// Conversationscreenshareeventtopicscreenshareconversation
type Conversationscreenshareeventtopicscreenshareconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Conversationscreenshareeventtopicscreensharemediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicscreenshareconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
