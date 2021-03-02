package platformclientv2
import (
	"encoding/json"
)

// Conversationcallbackeventtopicjourneycontext
type Conversationcallbackeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationcallbackeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationcallbackeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationcallbackeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcallbackeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
