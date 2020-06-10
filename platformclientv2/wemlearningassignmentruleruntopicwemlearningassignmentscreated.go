package platformclientv2
import (
	"encoding/json"
)

// Wemlearningassignmentruleruntopicwemlearningassignmentscreated
type Wemlearningassignmentruleruntopicwemlearningassignmentscreated struct { 
	// Module
	Module *Wemlearningassignmentruleruntopiclearningmodulereference `json:"module,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemlearningassignmentruleruntopicwemlearningassignmentscreated) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
