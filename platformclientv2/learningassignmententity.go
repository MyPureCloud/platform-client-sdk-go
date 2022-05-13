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

func (o *Learningassignmententity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmententity
	
	return json.Marshal(&struct { 
		AssignmentId *string `json:"assignmentId,omitempty"`
		*Alias
	}{ 
		AssignmentId: o.AssignmentId,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmententity) UnmarshalJSON(b []byte) error {
	var LearningassignmententityMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmententityMap)
	if err != nil {
		return err
	}
	
	if AssignmentId, ok := LearningassignmententityMap["assignmentId"].(string); ok {
		o.AssignmentId = &AssignmentId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmententity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
