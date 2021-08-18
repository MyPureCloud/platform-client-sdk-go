package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicjourneycontext
type Conversationsocialexpressioneventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationsocialexpressioneventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationsocialexpressioneventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationsocialexpressioneventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (u *Conversationsocialexpressioneventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationsocialexpressioneventtopicjourneycontext

	

	return json.Marshal(&struct { 
		Customer *Conversationsocialexpressioneventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationsocialexpressioneventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationsocialexpressioneventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
