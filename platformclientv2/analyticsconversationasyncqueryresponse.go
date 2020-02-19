package platformclientv2
import (
	"encoding/json"
)

// Analyticsconversationasyncqueryresponse
type Analyticsconversationasyncqueryresponse struct { 
	// Conversations
	Conversations *[]Analyticsconversation `json:"conversations,omitempty"`


	// Cursor - Optional cursor to indicate where to resume the results
	Cursor *string `json:"cursor,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationasyncqueryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
