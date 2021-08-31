package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingstatusdetailqueryfilter
type Routingstatusdetailqueryfilter struct { 
	// VarType - Boolean operation to apply to the provided predicates and clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - Boolean 'and/or' logic with up to two-levels of nesting
	Clauses *[]Routingstatusdetailqueryclause `json:"clauses,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Routingstatusdetailquerypredicate `json:"predicates,omitempty"`

}

func (o *Routingstatusdetailqueryfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingstatusdetailqueryfilter
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Clauses *[]Routingstatusdetailqueryclause `json:"clauses,omitempty"`
		
		Predicates *[]Routingstatusdetailquerypredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Clauses: o.Clauses,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Routingstatusdetailqueryfilter) UnmarshalJSON(b []byte) error {
	var RoutingstatusdetailqueryfilterMap map[string]interface{}
	err := json.Unmarshal(b, &RoutingstatusdetailqueryfilterMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := RoutingstatusdetailqueryfilterMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Clauses, ok := RoutingstatusdetailqueryfilterMap["clauses"].([]interface{}); ok {
		ClausesString, _ := json.Marshal(Clauses)
		json.Unmarshal(ClausesString, &o.Clauses)
	}
	
	if Predicates, ok := RoutingstatusdetailqueryfilterMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routingstatusdetailqueryfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
