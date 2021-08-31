package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcallbackeventtopicjourneycontext
type Conversationcallbackeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationcallbackeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationcallbackeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationcallbackeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Conversationcallbackeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcallbackeventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Conversationcallbackeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationcallbackeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationcallbackeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcallbackeventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var ConversationcallbackeventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcallbackeventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := ConversationcallbackeventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := ConversationcallbackeventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := ConversationcallbackeventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcallbackeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
