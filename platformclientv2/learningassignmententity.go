package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmententity
type Learningassignmententity struct { 
	// AssignmentId
	AssignmentId *string `json:"assignmentId,omitempty"`

}

func (u *Learningassignmententity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmententity

	

	return json.Marshal(&struct { 
		AssignmentId *string `json:"assignmentId,omitempty"`
		*Alias
	}{ 
		AssignmentId: u.AssignmentId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignmententity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
