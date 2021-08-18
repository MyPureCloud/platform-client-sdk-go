package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcobrowseeventtopicjourneycontext
type Queueconversationcobrowseeventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationcobrowseeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationcobrowseeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationcobrowseeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (u *Queueconversationcobrowseeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcobrowseeventtopicjourneycontext

	

	return json.Marshal(&struct { 
		Customer *Queueconversationcobrowseeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Queueconversationcobrowseeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Queueconversationcobrowseeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
