package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Digitlength
type Digitlength struct { 
	// Start
	Start *string `json:"start,omitempty"`


	// End
	End *string `json:"end,omitempty"`

}

// String returns a JSON representation of the model
func (o *Digitlength) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
