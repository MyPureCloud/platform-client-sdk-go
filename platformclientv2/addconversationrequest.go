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

func (o *Addconversationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Addconversationrequest
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		*Alias
	}{ 
		ConversationId: o.ConversationId,
		Alias:    (*Alias)(o),
	})
}

func (o *Addconversationrequest) UnmarshalJSON(b []byte) error {
	var AddconversationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AddconversationrequestMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := AddconversationrequestMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Addconversationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
