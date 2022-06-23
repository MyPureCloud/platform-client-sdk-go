package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationmessageeventtopicmessagemetadata
type Queueconversationmessageeventtopicmessagemetadata struct { 
	// VarType - Message type.
	VarType *string `json:"type,omitempty"`


	// Events - List of message events, if any
	Events *[]Queueconversationmessageeventtopicmessagemetadataevent `json:"events,omitempty"`


	// Content - List of message content, if any
	Content *[]Queueconversationmessageeventtopicmessagemetadatacontent `json:"content,omitempty"`

}

func (o *Queueconversationmessageeventtopicmessagemetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationmessageeventtopicmessagemetadata
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Events *[]Queueconversationmessageeventtopicmessagemetadataevent `json:"events,omitempty"`
		
		Content *[]Queueconversationmessageeventtopicmessagemetadatacontent `json:"content,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Events: o.Events,
		
		Content: o.Content,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationmessageeventtopicmessagemetadata) UnmarshalJSON(b []byte) error {
	var QueueconversationmessageeventtopicmessagemetadataMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationmessageeventtopicmessagemetadataMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := QueueconversationmessageeventtopicmessagemetadataMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Events, ok := QueueconversationmessageeventtopicmessagemetadataMap["events"].([]interface{}); ok {
		EventsString, _ := json.Marshal(Events)
		json.Unmarshal(EventsString, &o.Events)
	}
	
	if Content, ok := QueueconversationmessageeventtopicmessagemetadataMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicmessagemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
