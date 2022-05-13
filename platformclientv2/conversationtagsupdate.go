package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationtagsupdate
type Conversationtagsupdate struct { 
	// ExternalTag - The external tag associated with the conversation.
	ExternalTag *string `json:"externalTag,omitempty"`

}

func (o *Conversationtagsupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationtagsupdate
	
	return json.Marshal(&struct { 
		ExternalTag *string `json:"externalTag,omitempty"`
		*Alias
	}{ 
		ExternalTag: o.ExternalTag,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationtagsupdate) UnmarshalJSON(b []byte) error {
	var ConversationtagsupdateMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationtagsupdateMap)
	if err != nil {
		return err
	}
	
	if ExternalTag, ok := ConversationtagsupdateMap["externalTag"].(string); ok {
		o.ExternalTag = &ExternalTag
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationtagsupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
