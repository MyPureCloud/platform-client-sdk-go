package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessagemetadata - Metadata information about a message.
type Conversationmessagemetadata struct { 
	// VarType - Message type.
	VarType *string `json:"type,omitempty"`


	// Events - List of events metadata
	Events *[]Conversationmessagemetadataevent `json:"events,omitempty"`


	// Content - List of message content
	Content *[]Conversationmessagemetadatacontent `json:"content,omitempty"`

}

func (o *Conversationmessagemetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessagemetadata
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Events *[]Conversationmessagemetadataevent `json:"events,omitempty"`
		
		Content *[]Conversationmessagemetadatacontent `json:"content,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Events: o.Events,
		
		Content: o.Content,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessagemetadata) UnmarshalJSON(b []byte) error {
	var ConversationmessagemetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessagemetadataMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationmessagemetadataMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Events, ok := ConversationmessagemetadataMap["events"].([]interface{}); ok {
		EventsString, _ := json.Marshal(Events)
		json.Unmarshal(EventsString, &o.Events)
	}
	
	if Content, ok := ConversationmessagemetadataMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessagemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
