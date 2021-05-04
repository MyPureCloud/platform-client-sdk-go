package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentitem
type Learningassignmentitem struct { 
	// ModuleId - The Learning Module ID associated with this assignment
	ModuleId *string `json:"moduleId,omitempty"`


	// UserId - The User ID associated with this assignment
	UserId *string `json:"userId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
