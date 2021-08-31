package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationscreenshareeventtopicjourneycontext
type Queueconversationscreenshareeventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationscreenshareeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationscreenshareeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationscreenshareeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Queueconversationscreenshareeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationscreenshareeventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Queueconversationscreenshareeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Queueconversationscreenshareeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Queueconversationscreenshareeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationscreenshareeventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var QueueconversationscreenshareeventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationscreenshareeventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := QueueconversationscreenshareeventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := QueueconversationscreenshareeventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := QueueconversationscreenshareeventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
