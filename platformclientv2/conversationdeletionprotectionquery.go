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

func (o *Conversationdeletionprotectionquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationdeletionprotectionquery
	
	return json.Marshal(&struct { 
		ConversationIds *[]string `json:"conversationIds,omitempty"`
		*Alias
	}{ 
		ConversationIds: o.ConversationIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationdeletionprotectionquery) UnmarshalJSON(b []byte) error {
	var ConversationdeletionprotectionqueryMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationdeletionprotectionqueryMap)
	if err != nil {
		return err
	}
	
	if ConversationIds, ok := ConversationdeletionprotectionqueryMap["conversationIds"].([]interface{}); ok {
		ConversationIdsString, _ := json.Marshal(ConversationIds)
		json.Unmarshal(ConversationIdsString, &o.ConversationIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationdeletionprotectionquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
