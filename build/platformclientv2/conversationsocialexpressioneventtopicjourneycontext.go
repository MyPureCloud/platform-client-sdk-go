package platformclientv2
import (
	"encoding/json"
)

// Conversationsocialexpressioneventtopicjourneycontext
type Conversationsocialexpressioneventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationsocialexpressioneventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationsocialexpressioneventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationsocialexpressioneventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
