package platformclientv2
import (
	"encoding/json"
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
	return string(j)
}
