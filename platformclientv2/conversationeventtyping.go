package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtyping - A Typing event.
type Conversationeventtyping struct { 
	// VarType - Describes the type of Typing event.
	VarType *string `json:"type,omitempty"`

}

func (o *Conversationeventtyping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtyping
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationeventtyping) UnmarshalJSON(b []byte) error {
	var ConversationeventtypingMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtypingMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationeventtypingMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtyping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
