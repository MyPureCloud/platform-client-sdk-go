package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationbuttoncomponent - Structured template button object.
type Conversationbuttoncomponent struct { 
	// Title - Text to show inside the button.
	Title *string `json:"title,omitempty"`


	// Actions - The button actions.
	Actions *Conversationcontentactions `json:"actions,omitempty"`

}

func (o *Conversationbuttoncomponent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationbuttoncomponent
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Actions *Conversationcontentactions `json:"actions,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationbuttoncomponent) UnmarshalJSON(b []byte) error {
	var ConversationbuttoncomponentMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationbuttoncomponentMap)
	if err != nil {
		return err
	}
	
	if Title, ok := ConversationbuttoncomponentMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Actions, ok := ConversationbuttoncomponentMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationbuttoncomponent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
