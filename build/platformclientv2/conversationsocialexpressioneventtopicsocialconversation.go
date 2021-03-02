package platformclientv2
import (
	"encoding/json"
)

// Conversationsocialexpressioneventtopicsocialconversation
type Conversationsocialexpressioneventtopicsocialconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Conversationsocialexpressioneventtopicsocialmediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicsocialconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
