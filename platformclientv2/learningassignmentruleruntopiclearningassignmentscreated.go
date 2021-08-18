package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentruleruntopiclearningassignmentscreated
type Learningassignmentruleruntopiclearningassignmentscreated struct { 
	// Module
	Module *Learningassignmentruleruntopiclearningmodulereference `json:"module,omitempty"`

}

func (u *Learningassignmentruleruntopiclearningassignmentscreated) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentruleruntopiclearningassignmentscreated

	

	return json.Marshal(&struct { 
		Module *Learningassignmentruleruntopiclearningmodulereference `json:"module,omitempty"`
		*Alias
	}{ 
		Module: u.Module,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignmentruleruntopiclearningassignmentscreated) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
