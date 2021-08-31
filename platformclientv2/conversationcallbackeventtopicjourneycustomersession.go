package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcallbackeventtopicjourneycustomersession
type Conversationcallbackeventtopicjourneycustomersession struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *Conversationcallbackeventtopicjourneycustomersession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcallbackeventtopicjourneycustomersession
	
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

func (o *Conversationcallbackeventtopicjourneycustomersession) UnmarshalJSON(b []byte) error {
	var ConversationcallbackeventtopicjourneycustomersessionMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcallbackeventtopicjourneycustomersessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationcallbackeventtopicjourneycustomersessionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if VarType, ok := ConversationcallbackeventtopicjourneycustomersessionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcallbackeventtopicjourneycustomersession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
