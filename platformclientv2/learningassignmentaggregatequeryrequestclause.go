package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryrequestclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
