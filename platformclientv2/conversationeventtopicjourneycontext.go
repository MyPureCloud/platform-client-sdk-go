package platformclientv2
import (
	"encoding/json"
)

// Conversationeventtopicjourneycontext
type Conversationeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
