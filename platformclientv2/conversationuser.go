package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationuser
type Conversationuser struct { 
	// Id - The globally unique identifier for this user.
	Id *string `json:"id,omitempty"`

}

func (o *Conversationuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationuser) UnmarshalJSON(b []byte) error {
	var ConversationuserMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationuserMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
