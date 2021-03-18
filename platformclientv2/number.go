package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Number
type Number struct { 
	// Start
	Start *string `json:"start,omitempty"`


	// End
	End *string `json:"end,omitempty"`

}

// String returns a JSON representation of the model
func (o *Number) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
