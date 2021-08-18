package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationchateventtopicjourneycontext
type Conversationchateventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationchateventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationchateventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationchateventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (u *Conversationchateventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationchateventtopicjourneycontext

	

	return json.Marshal(&struct { 
		Customer *Conversationchateventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationchateventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationchateventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationchateventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
