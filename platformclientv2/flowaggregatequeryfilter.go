package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowaggregatequeryfilter
type Flowaggregatequeryfilter struct { 
	// VarType - Boolean operation to apply to the provided predicates and clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - Boolean 'and/or' logic with up to two-levels of nesting
	Clauses *[]Flowaggregatequeryclause `json:"clauses,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Flowaggregatequerypredicate `json:"predicates,omitempty"`

}

func (o *Flowaggregatequeryfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowaggregatequeryfilter
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Clauses *[]Flowaggregatequeryclause `json:"clauses,omitempty"`
		
		Predicates *[]Flowaggregatequerypredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Clauses: o.Clauses,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowaggregatequeryfilter) UnmarshalJSON(b []byte) error {
	var FlowaggregatequeryfilterMap map[string]interface{}
	err := json.Unmarshal(b, &FlowaggregatequeryfilterMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := FlowaggregatequeryfilterMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Clauses, ok := FlowaggregatequeryfilterMap["clauses"].([]interface{}); ok {
		ClausesString, _ := json.Marshal(Clauses)
		json.Unmarshal(ClausesString, &o.Clauses)
	}
	
	if Predicates, ok := FlowaggregatequeryfilterMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowaggregatequeryfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
