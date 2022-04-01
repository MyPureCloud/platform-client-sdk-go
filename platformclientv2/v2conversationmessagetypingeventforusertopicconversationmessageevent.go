package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationmessageevent
type V2conversationmessagetypingeventforusertopicconversationmessageevent struct { 
	// EventType
	EventType *string `json:"eventType,omitempty"`


	// CoBrowse
	CoBrowse *V2conversationmessagetypingeventforusertopicconversationeventcobrowse `json:"coBrowse,omitempty"`


	// Typing
	Typing *V2conversationmessagetypingeventforusertopicconversationeventtyping `json:"typing,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationmessageevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationmessageevent
	
	return json.Marshal(&struct { 
		EventType *string `json:"eventType,omitempty"`
		
		CoBrowse *V2conversationmessagetypingeventforusertopicconversationeventcobrowse `json:"coBrowse,omitempty"`
		
		Typing *V2conversationmessagetypingeventforusertopicconversationeventtyping `json:"typing,omitempty"`
		*Alias
	}{ 
		EventType: o.EventType,
		
		CoBrowse: o.CoBrowse,
		
		Typing: o.Typing,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicconversationmessageevent) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationmessageeventMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationmessageeventMap)
	if err != nil {
		return err
	}
	
	if EventType, ok := V2conversationmessagetypingeventforusertopicconversationmessageeventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
	
	if CoBrowse, ok := V2conversationmessagetypingeventforusertopicconversationmessageeventMap["coBrowse"].(map[string]interface{}); ok {
		CoBrowseString, _ := json.Marshal(CoBrowse)
		json.Unmarshal(CoBrowseString, &o.CoBrowse)
	}
	
	if Typing, ok := V2conversationmessagetypingeventforusertopicconversationmessageeventMap["typing"].(map[string]interface{}); ok {
		TypingString, _ := json.Marshal(Typing)
		json.Unmarshal(TypingString, &o.Typing)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationmessageevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
