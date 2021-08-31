package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicjourneycustomersession
type Queueconversationvideoeventtopicjourneycustomersession struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *Queueconversationvideoeventtopicjourneycustomersession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicjourneycustomersession
	
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

func (o *Queueconversationvideoeventtopicjourneycustomersession) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicjourneycustomersessionMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicjourneycustomersessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationvideoeventtopicjourneycustomersessionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if VarType, ok := QueueconversationvideoeventtopicjourneycustomersessionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicjourneycustomersession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
