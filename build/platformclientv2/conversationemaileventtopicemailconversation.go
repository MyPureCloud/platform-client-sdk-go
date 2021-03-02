package platformclientv2
import (
	"encoding/json"
)

// Conversationemaileventtopicemailconversation
type Conversationemaileventtopicemailconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Conversationemaileventtopicemailmediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicemailconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
