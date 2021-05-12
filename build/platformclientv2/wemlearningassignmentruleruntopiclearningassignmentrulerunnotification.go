package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wemlearningassignmentruleruntopiclearningassignmentrulerunnotification
type Wemlearningassignmentruleruntopiclearningassignmentrulerunnotification struct { 
	// Entities
	Entities *[]Wemlearningassignmentruleruntopicwemlearningassignmentscreated `json:"entities,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemlearningassignmentruleruntopiclearningassignmentrulerunnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
