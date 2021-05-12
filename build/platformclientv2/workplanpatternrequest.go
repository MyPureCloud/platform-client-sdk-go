package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanpatternrequest
type Workplanpatternrequest struct { 
	// WorkPlanIds - List of work plan IDs in order of rotation on a weekly basis. Values in the list cannot be null or empty
	WorkPlanIds *[]string `json:"workPlanIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanpatternrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
