package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsconversationwithoutattributesmultigetresponse
type Analyticsconversationwithoutattributesmultigetresponse struct { 
	// Conversations
	Conversations *[]Analyticsconversationwithoutattributes `json:"conversations,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationwithoutattributesmultigetresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
