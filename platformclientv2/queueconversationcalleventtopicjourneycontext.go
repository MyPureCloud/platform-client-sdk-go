package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcalleventtopicjourneycontext
type Queueconversationcalleventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationcalleventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationcalleventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationcalleventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Queueconversationcalleventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcalleventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Queueconversationcalleventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Queueconversationcalleventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Queueconversationcalleventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationcalleventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var QueueconversationcalleventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcalleventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := QueueconversationcalleventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := QueueconversationcalleventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := QueueconversationcalleventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
