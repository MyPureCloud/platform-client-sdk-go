package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflownotificationerrormessageparams
type Architectflownotificationerrormessageparams struct { 
	// AdditionalProperties
	AdditionalProperties *map[string]string `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflownotificationerrormessageparams) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
