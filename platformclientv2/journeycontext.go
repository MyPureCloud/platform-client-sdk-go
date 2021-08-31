package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeycontext
type Journeycontext struct { 
	// Customer - A subset of the Journey System's customer data at a point-in-time (for external linkage and internal usage/context)
	Customer *Journeycustomer `json:"customer,omitempty"`


	// CustomerSession - A subset of the Journey System's tracked customer session data at a point-in-time (for external linkage and internal usage/context)
	CustomerSession *Journeycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction - A subset of the Journey System's action data relevant to a part of a conversation (for external linkage and internal usage/context)
	TriggeringAction *Journeyaction `json:"triggeringAction,omitempty"`

}

func (o *Journeycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeycontext
	
	return json.Marshal(&struct { 
		Customer *Journeycustomer `json:"customer,omitempty"`
		
		CustomerSession *Journeycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Journeyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeycontext) UnmarshalJSON(b []byte) error {
	var JourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &JourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := JourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := JourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := JourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
