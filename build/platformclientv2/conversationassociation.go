package platformclientv2
import (
	"encoding/json"
)

// Conversationassociation
type Conversationassociation struct { 
	// ExternalContactId - An external contact ID.  If not supplied, implies the conversation should be disassociated with any external contact.
	ExternalContactId *string `json:"externalContactId,omitempty"`


	// ConversationId - Conversation ID
	ConversationId *string `json:"conversationId,omitempty"`


	// CommunicationId - Communication ID
	CommunicationId *string `json:"communicationId,omitempty"`


	// MediaType - Media type
	MediaType *string `json:"mediaType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationassociation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
