package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Attributedetaileventtopicattributeupdateevent
type Attributedetaileventtopicattributeupdateevent struct { 
	// EventTime
	EventTime *int `json:"eventTime,omitempty"`


	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// ParticipantId
	ParticipantId *string `json:"participantId,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`

}

func (o *Attributedetaileventtopicattributeupdateevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Attributedetaileventtopicattributeupdateevent
	
	return json.Marshal(&struct { 
		EventTime *int `json:"eventTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		*Alias
	}{ 
		EventTime: o.EventTime,
		
		ConversationId: o.ConversationId,
		
		ParticipantId: o.ParticipantId,
		
		Attributes: o.Attributes,
		Alias:    (*Alias)(o),
	})
}

func (o *Attributedetaileventtopicattributeupdateevent) UnmarshalJSON(b []byte) error {
	var AttributedetaileventtopicattributeupdateeventMap map[string]interface{}
	err := json.Unmarshal(b, &AttributedetaileventtopicattributeupdateeventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := AttributedetaileventtopicattributeupdateeventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := AttributedetaileventtopicattributeupdateeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := AttributedetaileventtopicattributeupdateeventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if Attributes, ok := AttributedetaileventtopicattributeupdateeventMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Attributedetaileventtopicattributeupdateevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
