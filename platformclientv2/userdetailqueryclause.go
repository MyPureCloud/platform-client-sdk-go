package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userdetailqueryclause
type Userdetailqueryclause struct { 
	// VarType - Boolean operation to apply to the provided predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Userdetailquerypredicate `json:"predicates,omitempty"`

}

func (o *Userdetailqueryclause) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userdetailqueryclause
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Predicates *[]Userdetailquerypredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Userdetailqueryclause) UnmarshalJSON(b []byte) error {
	var UserdetailqueryclauseMap map[string]interface{}
	err := json.Unmarshal(b, &UserdetailqueryclauseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := UserdetailqueryclauseMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Predicates, ok := UserdetailqueryclauseMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userdetailqueryclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
