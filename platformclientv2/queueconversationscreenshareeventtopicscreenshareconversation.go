package platformclientv2
import (
	"encoding/json"
)

// Queueconversationscreenshareeventtopicscreenshareconversation
type Queueconversationscreenshareeventtopicscreenshareconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Queueconversationscreenshareeventtopicscreensharemediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicscreenshareconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
