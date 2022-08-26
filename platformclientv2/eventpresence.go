package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Eventpresence - A Presence event.
type Eventpresence struct { 
	// VarType - Describes the type of Presence event.
	VarType *string `json:"type,omitempty"`

}

func (o *Eventpresence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Eventpresence
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Eventpresence) UnmarshalJSON(b []byte) error {
	var EventpresenceMap map[string]interface{}
	err := json.Unmarshal(b, &EventpresenceMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := EventpresenceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Eventpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
