package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationchateventtopicjourneycontext
type Conversationchateventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationchateventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationchateventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationchateventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Conversationchateventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationchateventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Conversationchateventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationchateventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationchateventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationchateventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var ConversationchateventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationchateventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := ConversationchateventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := ConversationchateventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := ConversationchateventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationchateventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
