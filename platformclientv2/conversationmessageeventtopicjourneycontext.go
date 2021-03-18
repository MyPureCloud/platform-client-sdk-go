package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
