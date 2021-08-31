package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcallbackeventtopicjourneycontext
type Queueconversationcallbackeventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationcallbackeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationcallbackeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationcallbackeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Queueconversationcallbackeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcallbackeventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Queueconversationcallbackeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Queueconversationcallbackeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Queueconversationcallbackeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationcallbackeventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var QueueconversationcallbackeventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcallbackeventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := QueueconversationcallbackeventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := QueueconversationcallbackeventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := QueueconversationcallbackeventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
