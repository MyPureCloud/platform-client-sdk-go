package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcalleventtopicjourneycontext
type Queueconversationcalleventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationcalleventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationcalleventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationcalleventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (u *Queueconversationcalleventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcalleventtopicjourneycontext

	

	return json.Marshal(&struct { 
		Customer *Queueconversationcalleventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Queueconversationcalleventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Queueconversationcalleventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
