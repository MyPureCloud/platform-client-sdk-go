package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowobservationqueryclause
type Flowobservationqueryclause struct { 
	// VarType - Boolean operation to apply to the provided predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Flowobservationquerypredicate `json:"predicates,omitempty"`

}

func (o *Flowobservationqueryclause) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowobservationqueryclause
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Predicates *[]Flowobservationquerypredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowobservationqueryclause) UnmarshalJSON(b []byte) error {
	var FlowobservationqueryclauseMap map[string]interface{}
	err := json.Unmarshal(b, &FlowobservationqueryclauseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := FlowobservationqueryclauseMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Predicates, ok := FlowobservationqueryclauseMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowobservationqueryclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
