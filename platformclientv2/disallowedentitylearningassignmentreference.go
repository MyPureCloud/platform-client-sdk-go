package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Disallowedentitylearningassignmentreference
type Disallowedentitylearningassignmentreference struct { 
	// ErrorCode - The error code associated with this disallowed entity
	ErrorCode *string `json:"errorCode,omitempty"`


	// Entity - The entity that was disallowed
	Entity *Learningassignmentreference `json:"entity,omitempty"`

}

// String returns a JSON representation of the model
func (o *Disallowedentitylearningassignmentreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
