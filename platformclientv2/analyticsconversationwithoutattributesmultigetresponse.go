package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsconversationwithoutattributesmultigetresponse
type Analyticsconversationwithoutattributesmultigetresponse struct { 
	// Conversations
	Conversations *[]Analyticsconversationwithoutattributes `json:"conversations,omitempty"`

}

func (u *Analyticsconversationwithoutattributesmultigetresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsconversationwithoutattributesmultigetresponse

	

	return json.Marshal(&struct { 
		Conversations *[]Analyticsconversationwithoutattributes `json:"conversations,omitempty"`
		*Alias
	}{ 
		Conversations: u.Conversations,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticsconversationwithoutattributesmultigetresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
