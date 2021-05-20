package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentaggregatequeryrequestfilter
type Learningassignmentaggregatequeryrequestfilter struct { 
	// VarType - The logic used to combine the clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - The list of clauses used to filter the data. Note that clauses must filter by attendeeId and a maximum of 100 user IDs are allowed
	Clauses *[]Learningassignmentaggregatequeryrequestclause `json:"clauses,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryrequestfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
