package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentruleruntopiclearningassignmentrulerunnotification
type Learningassignmentruleruntopiclearningassignmentrulerunnotification struct { 
	// Entities
	Entities *[]Learningassignmentruleruntopiclearningassignmentscreated `json:"entities,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentruleruntopiclearningassignmentrulerunnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
