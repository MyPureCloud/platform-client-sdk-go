package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentupdate
type Learningassignmentupdate struct { 
	// State - The Learning Assignment state
	State *string `json:"state,omitempty"`


	// Assessment - An updated Assessment
	Assessment *Learningassessment `json:"assessment,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
