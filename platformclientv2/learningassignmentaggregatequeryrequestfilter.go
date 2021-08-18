package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Learningassignmentaggregatequeryrequestfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentaggregatequeryrequestfilter

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Clauses *[]Learningassignmentaggregatequeryrequestclause `json:"clauses,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Clauses: u.Clauses,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryrequestfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
