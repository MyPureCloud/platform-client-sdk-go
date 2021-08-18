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

func (u *Journeycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeycontext

	

	return json.Marshal(&struct { 
		Customer *Journeycustomer `json:"customer,omitempty"`
		
		CustomerSession *Journeycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Journeyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Journeycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
