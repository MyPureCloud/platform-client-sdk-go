package platformclientv2
import (
	"encoding/json"
)

// Analyticsconversationqueryresponse
type Analyticsconversationqueryresponse struct { 
	// Conversations
	Conversations *[]Analyticsconversation `json:"conversations,omitempty"`


	// Aggregations
	Aggregations *[]Aggregationresult `json:"aggregations,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationqueryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
