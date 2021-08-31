package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationemaileventtopicjourneycontext
type Conversationemaileventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationemaileventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationemaileventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationemaileventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Conversationemaileventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationemaileventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Conversationemaileventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationemaileventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationemaileventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationemaileventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var ConversationemaileventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationemaileventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := ConversationemaileventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := ConversationemaileventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := ConversationemaileventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
