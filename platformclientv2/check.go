package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Check
type Check struct { 
	// Result - The result of a check executed. This indicates if the check was successful or not.
	Result *string `json:"result,omitempty"`


	// VarType - The type of check executed.
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Check) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
