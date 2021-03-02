package platformclientv2
import (
	"encoding/json"
)

// Analyticsconversationwithoutattributesmultigetresponse
type Analyticsconversationwithoutattributesmultigetresponse struct { 
	// Conversations
	Conversations *[]Analyticsconversationwithoutattributes `json:"conversations,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationwithoutattributesmultigetresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
