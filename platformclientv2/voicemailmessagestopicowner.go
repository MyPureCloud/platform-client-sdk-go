package platformclientv2
import (
	"encoding/json"
)

// Voicemailmessagestopicowner
type Voicemailmessagestopicowner struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemailmessagestopicowner) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
