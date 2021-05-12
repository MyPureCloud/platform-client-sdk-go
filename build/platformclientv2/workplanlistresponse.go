package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanlistresponse
type Workplanlistresponse struct { 
	// Entities
	Entities *[]Workplanlistitemresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanlistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
