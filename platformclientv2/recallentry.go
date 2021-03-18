package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Recallentry
type Recallentry struct { 
	// NbrAttempts
	NbrAttempts *int `json:"nbrAttempts,omitempty"`


	// MinutesBetweenAttempts
	MinutesBetweenAttempts *int `json:"minutesBetweenAttempts,omitempty"`

}

// String returns a JSON representation of the model
func (o *Recallentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
