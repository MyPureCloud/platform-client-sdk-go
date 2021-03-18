package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationmessageeventtopicjourneycontext
type Queueconversationmessageeventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationmessageeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationmessageeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationmessageeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
