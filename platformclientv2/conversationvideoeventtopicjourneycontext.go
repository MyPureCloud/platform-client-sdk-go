package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationvideoeventtopicjourneycontext
type Conversationvideoeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationvideoeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationvideoeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationvideoeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Conversationvideoeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationvideoeventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Conversationvideoeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationvideoeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationvideoeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationvideoeventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var ConversationvideoeventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationvideoeventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := ConversationvideoeventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := ConversationvideoeventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := ConversationvideoeventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationvideoeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
