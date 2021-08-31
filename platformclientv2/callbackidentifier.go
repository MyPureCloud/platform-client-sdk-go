package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callbackidentifier
type Callbackidentifier struct { 
	// VarType - The type of the associated callback participant
	VarType *string `json:"type,omitempty"`


	// Id - The identifier of the callback
	Id *string `json:"id,omitempty"`

}

func (o *Callbackidentifier) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callbackidentifier
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Callbackidentifier) UnmarshalJSON(b []byte) error {
	var CallbackidentifierMap map[string]interface{}
	err := json.Unmarshal(b, &CallbackidentifierMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := CallbackidentifierMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Id, ok := CallbackidentifierMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callbackidentifier) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
