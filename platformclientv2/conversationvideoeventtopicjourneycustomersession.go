package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationvideoeventtopicjourneycustomersession
type Conversationvideoeventtopicjourneycustomersession struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *Conversationvideoeventtopicjourneycustomersession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationvideoeventtopicjourneycustomersession
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationvideoeventtopicjourneycustomersession) UnmarshalJSON(b []byte) error {
	var ConversationvideoeventtopicjourneycustomersessionMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationvideoeventtopicjourneycustomersessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationvideoeventtopicjourneycustomersessionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if VarType, ok := ConversationvideoeventtopicjourneycustomersessionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationvideoeventtopicjourneycustomersession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
