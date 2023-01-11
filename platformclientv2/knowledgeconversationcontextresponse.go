package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeconversationcontextresponse
type Knowledgeconversationcontextresponse struct { 
	// Conversation - The conversation.
	Conversation *Addressableentityref `json:"conversation,omitempty"`


	// Queue - The queue used to assign the interaction to the user.
	Queue *Addressableentityref `json:"queue,omitempty"`


	// ExternalContact - The end-user participant of the conversation.
	ExternalContact *Addressableentityref `json:"externalContact,omitempty"`


	// MediaType - The media type of the conversation.
	MediaType *string `json:"mediaType,omitempty"`

}

func (o *Knowledgeconversationcontextresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeconversationcontextresponse
	
	return json.Marshal(&struct { 
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		
		Queue *Addressableentityref `json:"queue,omitempty"`
		
		ExternalContact *Addressableentityref `json:"externalContact,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		*Alias
	}{ 
		Conversation: o.Conversation,
		
		Queue: o.Queue,
		
		ExternalContact: o.ExternalContact,
		
		MediaType: o.MediaType,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeconversationcontextresponse) UnmarshalJSON(b []byte) error {
	var KnowledgeconversationcontextresponseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeconversationcontextresponseMap)
	if err != nil {
		return err
	}
	
	if Conversation, ok := KnowledgeconversationcontextresponseMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if Queue, ok := KnowledgeconversationcontextresponseMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if ExternalContact, ok := KnowledgeconversationcontextresponseMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if MediaType, ok := KnowledgeconversationcontextresponseMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeconversationcontextresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
