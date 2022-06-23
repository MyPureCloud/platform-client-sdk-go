package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicmessagemetadata - Metadata information about a message.
type Conversationeventtopicmessagemetadata struct { 
	// VarType - Message type.
	VarType *string `json:"type,omitempty"`


	// Events - List of message events, if any
	Events *[]Conversationeventtopicmessagemetadataevent `json:"events,omitempty"`


	// Content - List of message content, if any
	Content *[]Conversationeventtopicmessagemetadatacontent `json:"content,omitempty"`

}

func (o *Conversationeventtopicmessagemetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicmessagemetadata
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Events *[]Conversationeventtopicmessagemetadataevent `json:"events,omitempty"`
		
		Content *[]Conversationeventtopicmessagemetadatacontent `json:"content,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Events: o.Events,
		
		Content: o.Content,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationeventtopicmessagemetadata) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicmessagemetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicmessagemetadataMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationeventtopicmessagemetadataMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Events, ok := ConversationeventtopicmessagemetadataMap["events"].([]interface{}); ok {
		EventsString, _ := json.Marshal(Events)
		json.Unmarshal(EventsString, &o.Events)
	}
	
	if Content, ok := ConversationeventtopicmessagemetadataMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicmessagemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
