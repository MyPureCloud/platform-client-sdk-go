package platformclientv2
import (
	"encoding/json"
)

// Conversationassociation
type Conversationassociation struct { 
	// ExternalContactId - External Contact ID
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
