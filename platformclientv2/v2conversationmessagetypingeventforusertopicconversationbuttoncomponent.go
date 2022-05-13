package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationbuttoncomponent
type V2conversationmessagetypingeventforusertopicconversationbuttoncomponent struct { 
	// Title
	Title *string `json:"title,omitempty"`


	// Actions
	Actions *V2conversationmessagetypingeventforusertopicconversationcontentactions `json:"actions,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationbuttoncomponent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationbuttoncomponent
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Actions *V2conversationmessagetypingeventforusertopicconversationcontentactions `json:"actions,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicconversationbuttoncomponent) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationbuttoncomponentMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationbuttoncomponentMap)
	if err != nil {
		return err
	}
	
	if Title, ok := V2conversationmessagetypingeventforusertopicconversationbuttoncomponentMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Actions, ok := V2conversationmessagetypingeventforusertopicconversationbuttoncomponentMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationbuttoncomponent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
