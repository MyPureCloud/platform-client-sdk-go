package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationvideoeventtopicjourneycontext
type Conversationvideoeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationvideoeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationvideoeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationvideoeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (u *Conversationvideoeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationvideoeventtopicjourneycontext

	

	return json.Marshal(&struct { 
		Customer *Conversationvideoeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationvideoeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationvideoeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationvideoeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
