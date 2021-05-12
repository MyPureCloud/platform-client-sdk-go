package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bullseye
type Bullseye struct { 
	// Rings
	Rings *[]Ring `json:"rings,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bullseye) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
