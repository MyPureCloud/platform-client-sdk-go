package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationchateventtopicjourneycontext
type Queueconversationchateventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationchateventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationchateventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationchateventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (u *Queueconversationchateventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationchateventtopicjourneycontext

	

	return json.Marshal(&struct { 
		Customer *Queueconversationchateventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Queueconversationchateventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Queueconversationchateventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationchateventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
