package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationvideoeventtopicurireference
type Conversationvideoeventtopicurireference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

func (o *Conversationvideoeventtopicurireference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationvideoeventtopicurireference
	
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

func (o *Conversationvideoeventtopicurireference) UnmarshalJSON(b []byte) error {
	var ConversationvideoeventtopicurireferenceMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationvideoeventtopicurireferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationvideoeventtopicurireferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ConversationvideoeventtopicurireferenceMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationvideoeventtopicurireference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
