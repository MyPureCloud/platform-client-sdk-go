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

func (u *Conversationcallbackeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcallbackeventtopicjourneycontext

	

	return json.Marshal(&struct { 
		Customer *Conversationcallbackeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationcallbackeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationcallbackeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationcallbackeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
