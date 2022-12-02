package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusrulestopiccondition
type V2mobiusrulestopiccondition struct { 
	// Conditions
	Conditions *[]V2mobiusrulestopiccondition `json:"conditions,omitempty"`


	// Predicates
	Predicates *[]V2mobiusrulestopicconditionrulepredicate `json:"predicates,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *V2mobiusrulestopiccondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2mobiusrulestopiccondition
	
	return json.Marshal(&struct { 
		Conditions *[]V2mobiusrulestopiccondition `json:"conditions,omitempty"`
		
		Predicates *[]V2mobiusrulestopicconditionrulepredicate `json:"predicates,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Conditions: o.Conditions,
		
		Predicates: o.Predicates,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *V2mobiusrulestopiccondition) UnmarshalJSON(b []byte) error {
	var V2mobiusrulestopicconditionMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusrulestopicconditionMap)
	if err != nil {
		return err
	}
	
	if Conditions, ok := V2mobiusrulestopicconditionMap["conditions"].([]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if Predicates, ok := V2mobiusrulestopicconditionMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	
	if VarType, ok := V2mobiusrulestopicconditionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusrulestopiccondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
