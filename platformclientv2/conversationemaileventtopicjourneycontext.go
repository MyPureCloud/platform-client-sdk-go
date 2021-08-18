package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationemaileventtopicjourneycontext
type Conversationemaileventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationemaileventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationemaileventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationemaileventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (u *Conversationemaileventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationemaileventtopicjourneycontext

	

	return json.Marshal(&struct { 
		Customer *Conversationemaileventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationemaileventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationemaileventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
