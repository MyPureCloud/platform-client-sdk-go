package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Addconversationrequest - Update coaching appointment request
type Addconversationrequest struct { 
	// ConversationId - The id of the conversation to add
	ConversationId *string `json:"conversationId,omitempty"`

}

func (u *Addconversationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Addconversationrequest

	

	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		*Alias
	}{ 
		ConversationId: u.ConversationId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Addconversationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
