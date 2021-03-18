package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Activealertcount
type Activealertcount struct { 
	// Count - The count of active alerts for a user.
	Count *int `json:"count,omitempty"`

}

// String returns a JSON representation of the model
func (o *Activealertcount) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
