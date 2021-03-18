package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wemlearningassignmentruleruntopicwemlearningassignmentscreated
type Wemlearningassignmentruleruntopicwemlearningassignmentscreated struct { 
	// Module
	Module *Wemlearningassignmentruleruntopiclearningmodulereference `json:"module,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemlearningassignmentruleruntopicwemlearningassignmentscreated) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
