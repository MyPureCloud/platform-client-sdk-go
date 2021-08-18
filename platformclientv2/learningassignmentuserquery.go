package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentuserquery - Learning module users query request model
type Learningassignmentuserquery struct { 
	// Rule - Learning module rule object
	Rule *Learningmodulerule `json:"rule,omitempty"`


	// SearchTerm - The user name to be searched for
	SearchTerm *string `json:"searchTerm,omitempty"`

}

func (u *Learningassignmentuserquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentuserquery

	

	return json.Marshal(&struct { 
		Rule *Learningmodulerule `json:"rule,omitempty"`
		
		SearchTerm *string `json:"searchTerm,omitempty"`
		*Alias
	}{ 
		Rule: u.Rule,
		
		SearchTerm: u.SearchTerm,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignmentuserquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
