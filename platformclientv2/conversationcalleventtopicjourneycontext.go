package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcalleventtopicjourneycontext
type Conversationcalleventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationcalleventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationcalleventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationcalleventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (u *Conversationcalleventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcalleventtopicjourneycontext

	

	return json.Marshal(&struct { 
		Customer *Conversationcalleventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationcalleventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationcalleventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
