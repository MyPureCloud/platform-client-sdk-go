package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Termattribute
type Termattribute struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *Termattribute) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Termattribute
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Termattribute) UnmarshalJSON(b []byte) error {
	var TermattributeMap map[string]interface{}
	err := json.Unmarshal(b, &TermattributeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TermattributeMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := TermattributeMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if VarType, ok := TermattributeMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Termattribute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
