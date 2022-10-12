package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentqueryclause
type Documentqueryclause struct { 
	// Operator - Specifies how the predicates will be applied together.
	Operator *string `json:"operator,omitempty"`


	// Predicates - To apply multiple conditions. Limit of 10 predicates across all clauses.
	Predicates *[]Documentquerypredicate `json:"predicates,omitempty"`

}

func (o *Documentqueryclause) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentqueryclause
	
	return json.Marshal(&struct { 
		Operator *string `json:"operator,omitempty"`
		
		Predicates *[]Documentquerypredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		Operator: o.Operator,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentqueryclause) UnmarshalJSON(b []byte) error {
	var DocumentqueryclauseMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentqueryclauseMap)
	if err != nil {
		return err
	}
	
	if Operator, ok := DocumentqueryclauseMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Predicates, ok := DocumentqueryclauseMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentqueryclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
