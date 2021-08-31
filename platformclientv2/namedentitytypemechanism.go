package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentitytypemechanism
type Namedentitytypemechanism struct { 
	// Items - The items that define the named entity type.
	Items *[]Namedentitytypeitem `json:"items,omitempty"`


	// Restricted - Whether the named entity type is restricted to the items provided. Default: false
	Restricted *bool `json:"restricted,omitempty"`


	// VarType - The type of the mechanism.
	VarType *string `json:"type,omitempty"`

}

func (o *Namedentitytypemechanism) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Namedentitytypemechanism
	
	return json.Marshal(&struct { 
		Items *[]Namedentitytypeitem `json:"items,omitempty"`
		
		Restricted *bool `json:"restricted,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Items: o.Items,
		
		Restricted: o.Restricted,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Namedentitytypemechanism) UnmarshalJSON(b []byte) error {
	var NamedentitytypemechanismMap map[string]interface{}
	err := json.Unmarshal(b, &NamedentitytypemechanismMap)
	if err != nil {
		return err
	}
	
	if Items, ok := NamedentitytypemechanismMap["items"].([]interface{}); ok {
		ItemsString, _ := json.Marshal(Items)
		json.Unmarshal(ItemsString, &o.Items)
	}
	
	if Restricted, ok := NamedentitytypemechanismMap["restricted"].(bool); ok {
		o.Restricted = &Restricted
	}
	
	if VarType, ok := NamedentitytypemechanismMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Namedentitytypemechanism) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
