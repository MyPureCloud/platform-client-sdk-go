package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Executerecordingjobsquery
type Executerecordingjobsquery struct { 
	// State - The desired state for the job to be set to.
	State *string `json:"state,omitempty"`

}

// String returns a JSON representation of the model
func (o *Executerecordingjobsquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
