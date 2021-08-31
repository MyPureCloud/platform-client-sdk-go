package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicjourneycontext
type Conversationsocialexpressioneventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationsocialexpressioneventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationsocialexpressioneventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationsocialexpressioneventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Conversationsocialexpressioneventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationsocialexpressioneventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Conversationsocialexpressioneventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationsocialexpressioneventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationsocialexpressioneventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationsocialexpressioneventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var ConversationsocialexpressioneventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsocialexpressioneventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := ConversationsocialexpressioneventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := ConversationsocialexpressioneventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := ConversationsocialexpressioneventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
