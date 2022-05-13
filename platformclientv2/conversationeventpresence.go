package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventpresence - A Presence event.
type Conversationeventpresence struct { 
	// VarType - Describes the type of Presence event.
	VarType *string `json:"type,omitempty"`

}

func (o *Conversationeventpresence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventpresence
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationeventpresence) UnmarshalJSON(b []byte) error {
	var ConversationeventpresenceMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventpresenceMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationeventpresenceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
