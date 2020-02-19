package platformclientv2
import (
	"encoding/json"
)

// Queueconversationemaileventtopicemailconversation
type Queueconversationemaileventtopicemailconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Queueconversationemaileventtopicemailmediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicemailconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
