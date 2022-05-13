package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivityaggregatequeryrequestclause
type Developmentactivityaggregatequeryrequestclause struct { 
	// VarType - The logic used to combine the predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - The list of predicates used to filter the data
	Predicates *[]Developmentactivityaggregatequeryrequestpredicate `json:"predicates,omitempty"`

}

func (o *Developmentactivityaggregatequeryrequestclause) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Developmentactivityaggregatequeryrequestclause
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Predicates *[]Developmentactivityaggregatequeryrequestpredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Developmentactivityaggregatequeryrequestclause) UnmarshalJSON(b []byte) error {
	var DevelopmentactivityaggregatequeryrequestclauseMap map[string]interface{}
	err := json.Unmarshal(b, &DevelopmentactivityaggregatequeryrequestclauseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DevelopmentactivityaggregatequeryrequestclauseMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Predicates, ok := DevelopmentactivityaggregatequeryrequestclauseMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryrequestclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
