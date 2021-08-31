package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createcallbackresponse
type Createcallbackresponse struct { 
	// Conversation - The conversation associated with the callback
	Conversation *Domainentityref `json:"conversation,omitempty"`


	// CallbackIdentifiers - The list of communication identifiers for the callback participants
	CallbackIdentifiers *[]Callbackidentifier `json:"callbackIdentifiers,omitempty"`

}

func (o *Createcallbackresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createcallbackresponse
	
	return json.Marshal(&struct { 
		Conversation *Domainentityref `json:"conversation,omitempty"`
		
		CallbackIdentifiers *[]Callbackidentifier `json:"callbackIdentifiers,omitempty"`
		*Alias
	}{ 
		Conversation: o.Conversation,
		
		CallbackIdentifiers: o.CallbackIdentifiers,
		Alias:    (*Alias)(o),
	})
}

func (o *Createcallbackresponse) UnmarshalJSON(b []byte) error {
	var CreatecallbackresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CreatecallbackresponseMap)
	if err != nil {
		return err
	}
	
	if Conversation, ok := CreatecallbackresponseMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if CallbackIdentifiers, ok := CreatecallbackresponseMap["callbackIdentifiers"].([]interface{}); ok {
		CallbackIdentifiersString, _ := json.Marshal(CallbackIdentifiers)
		json.Unmarshal(CallbackIdentifiersString, &o.CallbackIdentifiers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createcallbackresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
