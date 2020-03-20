package platformclientv2
import (
	"encoding/json"
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
	return string(j)
}
