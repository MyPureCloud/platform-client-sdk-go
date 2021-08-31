package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcalleventtopicjourneycontext
type Conversationcalleventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationcalleventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationcalleventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationcalleventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Conversationcalleventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcalleventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Conversationcalleventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationcalleventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationcalleventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcalleventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var ConversationcalleventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcalleventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := ConversationcalleventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := ConversationcalleventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := ConversationcalleventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
