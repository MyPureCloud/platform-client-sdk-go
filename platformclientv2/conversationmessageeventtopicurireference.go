package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicurireference
type Conversationmessageeventtopicurireference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

func (o *Conversationmessageeventtopicurireference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicurireference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessageeventtopicurireference) UnmarshalJSON(b []byte) error {
	var ConversationmessageeventtopicurireferenceMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessageeventtopicurireferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationmessageeventtopicurireferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ConversationmessageeventtopicurireferenceMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicurireference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
