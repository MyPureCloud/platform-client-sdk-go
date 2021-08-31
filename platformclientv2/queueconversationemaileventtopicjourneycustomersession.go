package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationemaileventtopicjourneycustomersession
type Queueconversationemaileventtopicjourneycustomersession struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *Queueconversationemaileventtopicjourneycustomersession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationemaileventtopicjourneycustomersession
	
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

func (o *Queueconversationemaileventtopicjourneycustomersession) UnmarshalJSON(b []byte) error {
	var QueueconversationemaileventtopicjourneycustomersessionMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationemaileventtopicjourneycustomersessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationemaileventtopicjourneycustomersessionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if VarType, ok := QueueconversationemaileventtopicjourneycustomersessionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicjourneycustomersession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
