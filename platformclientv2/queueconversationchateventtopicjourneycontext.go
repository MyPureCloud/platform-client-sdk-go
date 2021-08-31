package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationchateventtopicjourneycontext
type Queueconversationchateventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationchateventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationchateventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationchateventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Queueconversationchateventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationchateventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Queueconversationchateventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Queueconversationchateventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Queueconversationchateventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationchateventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var QueueconversationchateventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationchateventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := QueueconversationchateventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := QueueconversationchateventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := QueueconversationchateventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationchateventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
