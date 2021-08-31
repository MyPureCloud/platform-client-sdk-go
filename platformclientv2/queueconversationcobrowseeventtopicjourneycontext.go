package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcobrowseeventtopicjourneycontext
type Queueconversationcobrowseeventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationcobrowseeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationcobrowseeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationcobrowseeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Queueconversationcobrowseeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcobrowseeventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Queueconversationcobrowseeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Queueconversationcobrowseeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Queueconversationcobrowseeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationcobrowseeventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var QueueconversationcobrowseeventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcobrowseeventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := QueueconversationcobrowseeventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := QueueconversationcobrowseeventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := QueueconversationcobrowseeventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
