package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentruleruntopiclearningassignmentscreated
type Learningassignmentruleruntopiclearningassignmentscreated struct { 
	// Module
	Module *Learningassignmentruleruntopiclearningmodulereference `json:"module,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentruleruntopiclearningassignmentscreated) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
