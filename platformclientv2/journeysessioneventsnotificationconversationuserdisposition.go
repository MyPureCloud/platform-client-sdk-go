package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationconversationuserdisposition
type Journeysessioneventsnotificationconversationuserdisposition struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Notes
	Notes *string `json:"notes,omitempty"`


	// User
	User *Journeysessioneventsnotificationuser `json:"user,omitempty"`

}

func (o *Journeysessioneventsnotificationconversationuserdisposition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationconversationuserdisposition
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		User *Journeysessioneventsnotificationuser `json:"user,omitempty"`
		*Alias
	}{ 
		Code: o.Code,
		
		Notes: o.Notes,
		
		User: o.User,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationconversationuserdisposition) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationconversationuserdispositionMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationconversationuserdispositionMap)
	if err != nil {
		return err
	}
	
	if Code, ok := JourneysessioneventsnotificationconversationuserdispositionMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Notes, ok := JourneysessioneventsnotificationconversationuserdispositionMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if User, ok := JourneysessioneventsnotificationconversationuserdispositionMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationconversationuserdisposition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
