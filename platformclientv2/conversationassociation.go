package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Conversationassociation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationassociation
	
	return json.Marshal(&struct { 
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		*Alias
	}{ 
		ExternalContactId: o.ExternalContactId,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		MediaType: o.MediaType,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationassociation) UnmarshalJSON(b []byte) error {
	var ConversationassociationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationassociationMap)
	if err != nil {
		return err
	}
	
	if ExternalContactId, ok := ConversationassociationMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if ConversationId, ok := ConversationassociationMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := ConversationassociationMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if MediaType, ok := ConversationassociationMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationassociation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
