package platformclientv2
import (
	"encoding/json"
)

// Createcallbackresponse
type Createcallbackresponse struct { 
	// Conversation - The conversation associated with the callback
	Conversation *Domainentityref `json:"conversation,omitempty"`


	// CallbackIdentifiers - The list of communication identifiers for the callback participants
	CallbackIdentifiers *[]Callbackidentifier `json:"callbackIdentifiers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createcallbackresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
