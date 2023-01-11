package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeconversationcontext
type Knowledgeconversationcontext struct { 
	// ConversationId - The unique identifier of the conversation.
	ConversationId *string `json:"conversationId,omitempty"`


	// MediaType - The media type of the conversation.
	MediaType *string `json:"mediaType,omitempty"`


	// QueueId - The unique identifier of the queue used to assign the interaction to the user.
	QueueId *string `json:"queueId,omitempty"`


	// ExternalContactId - The external contact identifier of the end-user participant.
	ExternalContactId *string `json:"externalContactId,omitempty"`

}

func (o *Knowledgeconversationcontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeconversationcontext
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		*Alias
	}{ 
		ConversationId: o.ConversationId,
		
		MediaType: o.MediaType,
		
		QueueId: o.QueueId,
		
		ExternalContactId: o.ExternalContactId,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeconversationcontext) UnmarshalJSON(b []byte) error {
	var KnowledgeconversationcontextMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeconversationcontextMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := KnowledgeconversationcontextMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if MediaType, ok := KnowledgeconversationcontextMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if QueueId, ok := KnowledgeconversationcontextMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if ExternalContactId, ok := KnowledgeconversationcontextMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeconversationcontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
