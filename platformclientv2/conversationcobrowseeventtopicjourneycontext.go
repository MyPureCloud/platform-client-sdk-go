package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcobrowseeventtopicjourneycontext
type Conversationcobrowseeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationcobrowseeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationcobrowseeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationcobrowseeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Conversationcobrowseeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcobrowseeventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Conversationcobrowseeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationcobrowseeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationcobrowseeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcobrowseeventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var ConversationcobrowseeventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcobrowseeventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := ConversationcobrowseeventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := ConversationcobrowseeventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := ConversationcobrowseeventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
