package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsconversationqueryresponse
type Analyticsconversationqueryresponse struct { 
	// Aggregations
	Aggregations *[]Aggregationresult `json:"aggregations,omitempty"`


	// Conversations
	Conversations *[]Analyticsconversationwithoutattributes `json:"conversations,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
