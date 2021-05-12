package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanpatternresponse
type Workplanpatternresponse struct { 
	// WorkPlans - List of work plans in order of rotation on a weekly basis
	WorkPlans *[]Workplanreference `json:"workPlans,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanpatternresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
