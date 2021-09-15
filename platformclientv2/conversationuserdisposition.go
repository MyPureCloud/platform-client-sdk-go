package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationuserdisposition
type Conversationuserdisposition struct { 
	// Code - User-defined wrap-up code for the conversation.
	Code *string `json:"code,omitempty"`


	// Notes - Text entered by the user to describe the call or disposition.
	Notes *string `json:"notes,omitempty"`


	// User - The user that wrapped up the conversation.
	User *Addressableentityref `json:"user,omitempty"`

}

func (o *Conversationuserdisposition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationuserdisposition
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		User *Addressableentityref `json:"user,omitempty"`
		*Alias
	}{ 
		Code: o.Code,
		
		Notes: o.Notes,
		
		User: o.User,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationuserdisposition) UnmarshalJSON(b []byte) error {
	var ConversationuserdispositionMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationuserdispositionMap)
	if err != nil {
		return err
	}
	
	if Code, ok := ConversationuserdispositionMap["code"].(string); ok {
		o.Code = &Code
	}
	
	if Notes, ok := ConversationuserdispositionMap["notes"].(string); ok {
		o.Notes = &Notes
	}
	
	if User, ok := ConversationuserdispositionMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationuserdisposition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
