package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcobrowseeventtopicjourneycontext
type Conversationcobrowseeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationcobrowseeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationcobrowseeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationcobrowseeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (u *Conversationcobrowseeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcobrowseeventtopicjourneycontext

	

	return json.Marshal(&struct { 
		Customer *Conversationcobrowseeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationcobrowseeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationcobrowseeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
