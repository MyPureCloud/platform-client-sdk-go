package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationscreenshareeventtopicjourneycontext
type Conversationscreenshareeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationscreenshareeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationscreenshareeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationscreenshareeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Conversationscreenshareeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationscreenshareeventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Conversationscreenshareeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationscreenshareeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationscreenshareeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationscreenshareeventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var ConversationscreenshareeventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationscreenshareeventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := ConversationscreenshareeventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := ConversationscreenshareeventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := ConversationscreenshareeventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
