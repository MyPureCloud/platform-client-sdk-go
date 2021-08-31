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

func (o *Analyticsconversationwithoutattributesmultigetresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsconversationwithoutattributesmultigetresponse
	
	return json.Marshal(&struct { 
		Conversations *[]Analyticsconversationwithoutattributes `json:"conversations,omitempty"`
		*Alias
	}{ 
		Conversations: o.Conversations,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsconversationwithoutattributesmultigetresponse) UnmarshalJSON(b []byte) error {
	var AnalyticsconversationwithoutattributesmultigetresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsconversationwithoutattributesmultigetresponseMap)
	if err != nil {
		return err
	}
	
	if Conversations, ok := AnalyticsconversationwithoutattributesmultigetresponseMap["conversations"].([]interface{}); ok {
		ConversationsString, _ := json.Marshal(Conversations)
		json.Unmarshal(ConversationsString, &o.Conversations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsconversationwithoutattributesmultigetresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
