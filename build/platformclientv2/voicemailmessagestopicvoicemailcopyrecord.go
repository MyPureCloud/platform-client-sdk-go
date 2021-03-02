package platformclientv2
import (
	"encoding/json"
)

// Voicemailmessagestopicvoicemailcopyrecord
type Voicemailmessagestopicvoicemailcopyrecord struct { 
	// User
	User *Voicemailmessagestopicowner `json:"user,omitempty"`


	// Group
	Group *Voicemailmessagestopicowner `json:"group,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemailmessagestopicvoicemailcopyrecord) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
