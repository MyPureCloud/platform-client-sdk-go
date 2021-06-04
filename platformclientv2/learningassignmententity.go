package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmententity
type Learningassignmententity struct { 
	// AssignmentId
	AssignmentId *string `json:"assignmentId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningassignmententity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
