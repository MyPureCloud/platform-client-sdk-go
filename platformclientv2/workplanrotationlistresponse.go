package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanrotationlistresponse
type Workplanrotationlistresponse struct { 
	// Entities
	Entities *[]Workplanrotationresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanrotationlistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
