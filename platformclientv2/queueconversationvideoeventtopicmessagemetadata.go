package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicmessagemetadata - Metadata information about a message.
type Queueconversationvideoeventtopicmessagemetadata struct { 
	// VarType - Message type.
	VarType *string `json:"type,omitempty"`


	// Events - List of message events, if any
	Events *[]Queueconversationvideoeventtopicmessagemetadataevent `json:"events,omitempty"`


	// Content - List of message content, if any
	Content *[]Queueconversationvideoeventtopicmessagemetadatacontent `json:"content,omitempty"`

}

func (o *Queueconversationvideoeventtopicmessagemetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicmessagemetadata
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Events *[]Queueconversationvideoeventtopicmessagemetadataevent `json:"events,omitempty"`
		
		Content *[]Queueconversationvideoeventtopicmessagemetadatacontent `json:"content,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Events: o.Events,
		
		Content: o.Content,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopicmessagemetadata) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicmessagemetadataMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicmessagemetadataMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := QueueconversationvideoeventtopicmessagemetadataMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Events, ok := QueueconversationvideoeventtopicmessagemetadataMap["events"].([]interface{}); ok {
		EventsString, _ := json.Marshal(Events)
		json.Unmarshal(EventsString, &o.Events)
	}
	
	if Content, ok := QueueconversationvideoeventtopicmessagemetadataMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicmessagemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
