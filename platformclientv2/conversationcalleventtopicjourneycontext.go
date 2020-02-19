package platformclientv2
import (
	"encoding/json"
)

// Conversationcalleventtopicjourneycontext
type Conversationcalleventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationcalleventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationcalleventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationcalleventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
