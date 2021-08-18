package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationscreenshareeventtopicjourneycontext
type Conversationscreenshareeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationscreenshareeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationscreenshareeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationscreenshareeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (u *Conversationscreenshareeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationscreenshareeventtopicjourneycontext

	

	return json.Marshal(&struct { 
		Customer *Conversationscreenshareeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationscreenshareeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationscreenshareeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
