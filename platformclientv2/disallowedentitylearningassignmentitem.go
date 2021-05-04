package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Disallowedentitylearningassignmentitem
type Disallowedentitylearningassignmentitem struct { 
	// ErrorCode - The error code associated with this disallowed entity
	ErrorCode *string `json:"errorCode,omitempty"`


	// Entity - The entity that was disallowed
	Entity *Learningassignmentitem `json:"entity,omitempty"`

}

// String returns a JSON representation of the model
func (o *Disallowedentitylearningassignmentitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
