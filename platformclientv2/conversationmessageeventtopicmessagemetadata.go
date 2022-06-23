package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicmessagemetadata
type Conversationmessageeventtopicmessagemetadata struct { 
	// VarType - Message type.
	VarType *string `json:"type,omitempty"`


	// Events - List of message events, if any
	Events *[]Conversationmessageeventtopicmessagemetadataevent `json:"events,omitempty"`


	// Content - List of message content, if any
	Content *[]Conversationmessageeventtopicmessagemetadatacontent `json:"content,omitempty"`

}

func (o *Conversationmessageeventtopicmessagemetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicmessagemetadata
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Events *[]Conversationmessageeventtopicmessagemetadataevent `json:"events,omitempty"`
		
		Content *[]Conversationmessageeventtopicmessagemetadatacontent `json:"content,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Events: o.Events,
		
		Content: o.Content,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessageeventtopicmessagemetadata) UnmarshalJSON(b []byte) error {
	var ConversationmessageeventtopicmessagemetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessageeventtopicmessagemetadataMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationmessageeventtopicmessagemetadataMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Events, ok := ConversationmessageeventtopicmessagemetadataMap["events"].([]interface{}); ok {
		EventsString, _ := json.Marshal(Events)
		json.Unmarshal(EventsString, &o.Events)
	}
	
	if Content, ok := ConversationmessageeventtopicmessagemetadataMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicmessagemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
