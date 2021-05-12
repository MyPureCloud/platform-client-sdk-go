package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queryrequestclause
type Queryrequestclause struct { 
	// VarType - The logic used to combine the predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - The list of predicates used to filter the data
	Predicates *[]Queryrequestpredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queryrequestclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
