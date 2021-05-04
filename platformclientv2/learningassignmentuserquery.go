package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Learningassignmentuserquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
