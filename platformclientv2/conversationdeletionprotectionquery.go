package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationdeletionprotectionquery
type Conversationdeletionprotectionquery struct { 
	// ConversationIds - This is a list of ConversationIds. The list cannot exceed 100 conversationids.
	ConversationIds *[]string `json:"conversationIds,omitempty"`

}

func (u *Conversationdeletionprotectionquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationdeletionprotectionquery

	

	return json.Marshal(&struct { 
		ConversationIds *[]string `json:"conversationIds,omitempty"`
		*Alias
	}{ 
		ConversationIds: u.ConversationIds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationdeletionprotectionquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
