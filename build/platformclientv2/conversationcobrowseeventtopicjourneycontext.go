package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
