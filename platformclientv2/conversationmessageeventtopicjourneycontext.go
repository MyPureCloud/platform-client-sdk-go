package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicjourneycontext
type Conversationmessageeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationmessageeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationmessageeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationmessageeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (u *Conversationmessageeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicjourneycontext

	

	return json.Marshal(&struct { 
		Customer *Conversationmessageeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationmessageeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationmessageeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: u.Customer,
		
		CustomerSession: u.CustomerSession,
		
		TriggeringAction: u.TriggeringAction,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
