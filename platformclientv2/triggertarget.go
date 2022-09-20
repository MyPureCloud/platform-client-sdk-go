package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Triggertarget - The target of a trigger invocation
type Triggertarget struct { 
	// VarType - The entity type to target
	VarType *string `json:"type,omitempty"`


	// Id - The ID of the entity to target
	Id *string `json:"id,omitempty"`

}

func (o *Triggertarget) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Triggertarget
	
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

func (o *Triggertarget) UnmarshalJSON(b []byte) error {
	var TriggertargetMap map[string]interface{}
	err := json.Unmarshal(b, &TriggertargetMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := TriggertargetMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Id, ok := TriggertargetMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Triggertarget) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
