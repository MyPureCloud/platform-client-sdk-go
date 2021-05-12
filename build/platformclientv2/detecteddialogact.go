package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Detecteddialogact
type Detecteddialogact struct { 
	// Name - The name of the detected dialog act.
	Name *string `json:"name,omitempty"`


	// Probability - The probability of the detected dialog act.
	Probability *float64 `json:"probability,omitempty"`

}

// String returns a JSON representation of the model
func (o *Detecteddialogact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
