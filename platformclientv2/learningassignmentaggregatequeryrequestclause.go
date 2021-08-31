package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentaggregatequeryrequestclause
type Learningassignmentaggregatequeryrequestclause struct { 
	// VarType - The logic used to combine the predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - The list of predicates used to filter the data
	Predicates *[]Learningassignmentaggregatequeryrequestpredicate `json:"predicates,omitempty"`

}

func (o *Learningassignmentaggregatequeryrequestclause) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentaggregatequeryrequestclause
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Predicates *[]Learningassignmentaggregatequeryrequestpredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentaggregatequeryrequestclause) UnmarshalJSON(b []byte) error {
	var LearningassignmentaggregatequeryrequestclauseMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentaggregatequeryrequestclauseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := LearningassignmentaggregatequeryrequestclauseMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Predicates, ok := LearningassignmentaggregatequeryrequestclauseMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryrequestclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
