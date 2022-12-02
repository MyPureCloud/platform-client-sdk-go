package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusalertstopiccondition
type V2mobiusalertstopiccondition struct { 
	// Conditions
	Conditions *[]V2mobiusalertstopiccondition `json:"conditions,omitempty"`


	// Predicates
	Predicates *[]V2mobiusalertstopicconditionrulepredicate `json:"predicates,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *V2mobiusalertstopiccondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2mobiusalertstopiccondition
	
	return json.Marshal(&struct { 
		Conditions *[]V2mobiusalertstopiccondition `json:"conditions,omitempty"`
		
		Predicates *[]V2mobiusalertstopicconditionrulepredicate `json:"predicates,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Conditions: o.Conditions,
		
		Predicates: o.Predicates,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *V2mobiusalertstopiccondition) UnmarshalJSON(b []byte) error {
	var V2mobiusalertstopicconditionMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusalertstopicconditionMap)
	if err != nil {
		return err
	}
	
	if Conditions, ok := V2mobiusalertstopicconditionMap["conditions"].([]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if Predicates, ok := V2mobiusalertstopicconditionMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	
	if VarType, ok := V2mobiusalertstopicconditionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusalertstopiccondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
